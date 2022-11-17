#Day6 负载均衡
1、负载均衡策略
    随机选择
    轮训算法Round Robin
    加权轮询Weight Round Robin
    哈希/一致性哈希策略

2、服务发现
    MultiServersDiscovery，服务列表由手工维护的服务发现的结构体

3、支持负载均衡的客户端
    XClient 向用户暴露一个支持负载均衡的客户端
        dial()
        Broadcast()