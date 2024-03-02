GeeCache

分布式缓存系统。

[TOC]



# 前言

设计一个分布式缓存系统，需要考虑资源控制、淘汰策略、并发、分布式节点通信等各个方面的问题。

[groupcache](https://github.com/golang/groupcache) 是 Go 语言版的 memcached，groupcache 的作者也是 memcached 的作者。无论是了解单机缓存还是分布式缓存，深入学习这个库的实现都是非常有意义的。

GeeCache是模仿groupcache，简化设计的分布式缓存系统。

# 目录

- [x] LRU 缓存淘汰策略
  - [x] FIFO/LFU/LRU 算法简介
  - [x] LRU 算法实现

- [ ] 单机并发缓存
  - [ ] sync.Mutex 互斥锁
  - [ ] 互斥锁封装LRU，支持并发读写
  - [ ] GeeCache核心数据结构Group，负责与用户交互，控制缓存值和获取的流程

- [ ] HTTP服务端
  - [ ] go语言http标准库
  - [ ] 单机节点搭建HTTP服务端

- [ ] 一致性哈希（Hash）
  - [ ] 原理
  - [ ] 实现

- [ ] 分布式节点
  - [ ] 节点注册与选择
  - [ ] 实现HTTP客户端，与远程服务端节点通信

- [ ] 防止缓存击穿
  - [ ] 缓存雪崩、缓存击穿与缓存穿透的概念简介
  - [ ] 使用 singleflight 防止缓存击穿

- [ ] 使用Protobuf通信
  - [ ] protobuf简介与优点
  - [ ] 使用protobuf替换字节流进行节点间通信






# 参考

7天用Go从零实现分布式缓存GeeCache

https://geektutu.com/post/geecache.html