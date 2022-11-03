#Day7  使用Protobuf通信

protobuf 即 Protocol Buffers，Google 开发的一种数据描述语言，是一种轻便高效的结构化数据存储格式，与语言、平台无关，可扩展可序列化。protobuf 以二进制方式存储，占用空间小。
https://geektutu.com/post/quick-go-protobuf.html Go Protobuf 简明教程

protobuf 广泛地应用于远程过程调用(RPC) 的二进制传输，使用 protobuf 的目的非常简单，为了获得更高的性能。传输前使用 protobuf 编码，接收方再进行解码，可以显著地降低二进制传输的大小。另外一方面，protobuf 可非常适合传输结构化数据，便于通信字段的扩展。

使用 protobuf 一般分为以下 2 步：

按照 protobuf 的语法，在 .proto 文件中定义数据结构，并使用 protoc 生成 Go 代码（.proto 文件是跨平台的，还可以生成 C、Java 等其他源码文件）。
在项目代码中引用生成的 Go 代码。