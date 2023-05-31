# 服务注册中心
> 手动实现服务注册中心

目录 |二级目录
---|---
[理解](#理解) |[读写模型](#读写模型)
[实现](#实现) |[功能分解](#功能分解)，[技术选型](#技术选型)，[高性能高可用海量存储](#高性能高可用海量存储)
[运行展示](#运行展示) |[开启服务](#开启服务)；[注册](#注册)、[下线](#下线)、[续约](#续约)，[获取列表](#获取列表)

***

## 理解
### 读写模型
> 服务注册中心的基本功能是实现 **读写模型**，即 *增删查改*。

所有的数据库都是读写模型。

***

## 实现
### 功能分解
类别|说明
---|---
保存注册表	|内存中的map/Etcd
增删查改	|服务方（注册、下线、续约）、消费方（获取列表）、注册中心（剔除过期、最低剔除保护）
对外提供http服务 |使用Web框架，或语言自带网络库
### 技术选型
类别|对比|选择
---|---|---
节点关系（点对点/主从）	|点对点：读写性能佳、数据存在不一致。实现简单，节点状态只考虑上线/下线	|☑️
&nbsp; |主从（读写分离）：数据一致、写性能较差。需要考虑分布式一致算法、维护状态机	
CP/AP  |&nbsp; |注册中心允许短暂不一致、不能容忍不可用，所以选择AP
### 高性能高可用海量存储
> *副本 + 分区*

类别|说明
---|---
数据同步	|（节点启动时）从其他节点同步数据，（服务变更本节点数据后）同步给其他节点
一致性	|若同步其他节点失败，会通过下次操作/超时自动清除来达到一致，也可以进行定时弥补来达到一致。
写方式	|同步/异步/半同步，权衡的是性能和数据一致性。
数据分区 |通过zone层来隔离数据

***

## 运行展示
### 开启服务 
```shell
cd cmd
go build -o discovery main.go
./discovery -c configs.yaml
```

### 注册
```shell
curl -XPOST http://127.0.0.1:6666/api/register -H 'Content-Type:application/json' -d'{"env":"dev", "appid":"testapp","hostname":"testhost1","addrs":["rpc:aaa","rpc:bbb"],"status":1,"replication":true}'
```
### 下线
```shell
curl -XPOST http://127.0.0.1:8866/api/cancel -H 'Content-Type:application/json' -d'{"env":"dev","appid":"testapp","hostname":"testhost","replication":true}'
```
### 续约
```shell
curl -XPOST http://127.0.0.1:8866/api/renew -H 'Content-Type:application/json' -d'{"env":"dev","appid":"testapp","hostname":"testhost","replication":true}'
```

### 获取列表
```shell
curl -XPOST http://127.0.0.1:6666/api/fetch -H 'Content-Type:application/json' -d'{"env":"dev", "appid":"testapp","status":1}'

curl -XPOST http://127.0.0.1:6666/api/fetchall
```