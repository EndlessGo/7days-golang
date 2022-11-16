#Day2 高性能客户端
1、封装了结构体 Call 来承载一次 RPC 调用所需要的信息

2、实现 GeeRPC 客户端最核心的部分 Client

Client 的字段比较复杂：
cc 是消息的编解码器，和服务端类似，用来序列化将要发送出去的请求，以及反序列化接收到的响应。
sending 是一个互斥锁，和服务端类似，为了保证请求的有序发送，即防止出现多个请求报文混淆。
header 是每个请求的消息头，header 只有在请求发送时才需要，而请求发送是互斥的，因此每个客户端只需要一个，声明在 Client 结构体中可以复用。
seq 用于给发送的请求编号，每个请求拥有唯一编号。
pending 存储未处理完的请求，键是编号，值是 Call 实例。
closing 和 shutdown 任意一个值置为 true，则表示 Client 处于不可用的状态，但有些许的差别，closing 是用户主动关闭的，即调用 Close 方法，而 shutdown 置为 true 一般是有错误发生。

Client 方法
    registerCall：将参数 call 添加到 client.pending 中，并更新 client.seq。
    removeCall：根据 seq，从 client.pending 中移除对应的 call，并返回。
    terminateCalls：服务端或客户端发生错误时调用，将 shutdown 设置为 true，且将错误信息通知所有 pending 状态的 call。

    receive()
        对一个客户端端来说，接收响应、发送请求是最重要的 2 个功能。那么首先实现接收功能，接收到的响应有三种情况： 
        call 不存在，可能是请求没有发送完整，或者因为其他原因被取消，但是服务端仍旧处理了。
        call 存在，但服务端处理出错，即 h.Error 不为空。
        call 存在，服务端处理正常，那么需要从 body 中读取 Reply 的值。

    Go 和 Call 是客户端暴露给用户的两个 RPC 服务调用接口，Go 是一个异步接口，返回 call 实例。
    Call 是对 Go 的封装，阻塞 call.Done，等待响应返回，是一个同步接口。