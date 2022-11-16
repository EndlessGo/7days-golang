#Day4 超时处理
原因：服务的可用性

GeeRPC 在 3 个地方添加了超时处理机制。分别是：
1）客户端创建连接时
    net.DialTimeout替换net.Dial    
2）客户端 Client.Call() 整个过程导致的超时（包含发送报文，等待处理，接收报文所有阶段）
    用户可以使用 context.WithTimeout 创建具备超时检测能力的 context 对象来控制
3）服务端处理报文，即 Server.handleRequest 超时。
    实现与客户端很接近，使用 time.After() 结合 select+chan 完成。

 