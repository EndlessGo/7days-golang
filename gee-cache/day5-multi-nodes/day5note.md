#Day5 分布式节点
TODO: 待细看

流程回顾
    检查是否被缓存    
    是否应当从远程节点获取（*）
    调用`回调函数`，获取值并添加到缓存

（*）流程
    使用一致性哈希选择节点
    HTTP 客户端访问远程节点
    回退到本地节点处理

抽象 PeerPicker
    在这里，抽象出 2 个接口，PeerPicker 的 PickPeer() 方法用于根据传入的 key 选择相应节点 PeerGetter。
    接口 PeerGetter 的 Get() 方法用于从对应 group 查找缓存值。PeerGetter 就对应于上述流程中的 HTTP 客户端。

节点选择与 HTTP 客户端
    创建具体的 HTTP 客户端类 httpGetter，实现 PeerGetter 接口。
    为 HTTPPool 添加节点选择的功能。

实现主流程
    新增 RegisterPeers() 方法，将 实现了 PeerPicker 接口的 HTTPPool 注入到 Group 中。
    新增 getFromPeer() 方法，使用实现了 PeerGetter 接口的 httpGetter 从访问远程节点，获取缓存值。
    修改 load 方法，使用 PickPeer() 方法选择节点，若非本机节点，则调用 getFromPeer() 从远程获取。若是本机节点或失败，则回退到 getLocally()。

main 函数测试
    startCacheServer() 用来启动缓存服务器：创建 HTTPPool，添加节点信息，注册到 gee 中，启动 HTTP 服务（共3个端口，8001/8002/8003），用户不感知。
    startAPIServer() 用来启动一个 API 服务（端口 9999），与用户进行交互，用户感知。
    main() 函数需要命令行传入 port 和 api 2 个参数，用来在指定端口启动 HTTP 服务。
    run.sh