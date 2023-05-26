# 服务注册中心
### 功能
类别|说明
---|---
保存注册表	|内存中的map/Etcd
增删查改	|服务方（注册、下线、续约）、消费方（获取列表）、注册中心（剔除过期、最低剔除保护）
对外提供http服务 |

### 技术选型
类别|对比|选择
---|---|---
节点关系（点对点/主从）	|点对点：读写性能佳、数据存在不一致。实现简单，节点状态只考虑上线/下线	|☑️
&nbsp; |主从（读写分离）：数据一致、写性能较差。需要考虑分布式一致算法、维护状态机	
CP/AP  |&nbsp; |注册中心允许短暂不一致、不能容忍不可用，所以选择AP

### 高性能、高可用（副本）
类别|说明
---|---
数据同步	|（节点启动时）从其他节点同步数据，（服务变更本节点数据后）同步给其他节点
一致性	|若同步其他节点失败，会通过下次操作/超时自动清除来达到一致，也可以进行定时弥补来达到一致。
写方式	|同步/异步/半同步，权衡的是性能和数据一致性。
### 海量存储（数据分区）
增加zone层来隔离数据。

## 运行
- 服务构建 
```shell
cd cmd
go build -o discovery main.go
./discovery -c configs.yaml
```

- 服务注册
```shell
curl -XPOST http://127.0.0.1:6666/api/register -H 'Content-Type:application/json' -d'{"env":"dev", "appid":"testapp","hostname":"testhost1","addrs":["rpc:aaa","rpc:bbb"],"status":1,"replication":true}'
```

- 服务发现
```shell
curl -XPOST http://127.0.0.1:6666/api/fetch -H 'Content-Type:application/json' -d'{"env":"dev", "appid":"testapp","status":1}'

curl -XPOST http://127.0.0.1:6666/api/fetchall
```

- 服务续约
```shell
curl -XPOST http://127.0.0.1:8866/api/renew -H 'Content-Type:application/json' -d'{"env":"dev","appid":"testapp","hostname":"testhost","replication":true}'
```

- 服务取消
```shell
curl -XPOST http://127.0.0.1:8866/api/cancel -H 'Content-Type:application/json' -d'{"env":"dev","appid":"testapp","hostname":"testhost","replication":true}'
```