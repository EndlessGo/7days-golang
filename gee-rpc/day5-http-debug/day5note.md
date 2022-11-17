#Day5 支持HTTP协议
1、支持 HTTP 协议要做什么
对 RPC 服务端来，需要做的是将 HTTP 协议转换为 RPC 协议，对客户端来说，需要新增通过 HTTP CONNECT 请求创建连接的逻辑。

2、服务端支持 HTTP 协议
    Server(struct)  
        ServerHTTP() //http.handle
      
3、客户端支持 HTTP 协议
    NewHTTPClient()    

4、 DEBUG 页面