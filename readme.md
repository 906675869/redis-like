# 寄语
希望能写一个分布式KV存储

## issue 2 
编写执行框架，完成类springMVC架构功能，将控制逻辑与业务逻辑分离

## 能力具备
1. reactor 网络模型
2. 分布式共识算法
3. 存储引擎

> 希望能尽量自己完成，而不借助于开源方案。思想可以用，代码尽量不用

## 前期准备
- [x] 熟悉go中的IO相关的api
- [ ] 理解reactor网络模型
- [ ] 理解分布式共识算法
- [ ] 理解并写出类levelDB的存储引擎（基于levelDB的思想，自己实现）
- [x] resp 协议规范与解析 [redis的resp解析](https://redis.io/topics/protocol)

## 目标
完成一个单体kv存储。
1. 协议层 使用redis的resp
2. 存储层 使用levelDB存储引擎
3. 将协议层与存储层分离

### 实现存储与业务分离
1. 协议与存储分离
2. 存储引擎换为两套，一套为levelDB，一套为内存缓存框架freecache（lru && ttl ，本处目的为将协议与存储分离）
