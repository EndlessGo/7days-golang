# 自学补充知识
Day0
	HTTP基础知识
	各种Web框架源码
	curl用法 https://www.ruanyifeng.com/blog/2019/09/curl-reference.html
Day2
	HTTP细节：
		Request/Response
		GET/POST
		HTML/JSON
Day3 
	前缀树

#Day0 序言
net/http提供了基础的Web功能：监听端口，映射静态路由，解析HTTP报文。
但一些Web开发中简单的需求并不支持，需要手工实现并频繁地手工处理：动态路由，鉴权，模板。
这就是框架的价值所在，同样也并不是每一个频繁处理的地方都适合在框架中完成。

Web框架
	python：django大而全，flask小而美，bottle著名的微框架（路由，模板，工具集cookies,headers处理机制，插件）
	go：Beego, Gin, Iris

练手项目Gee，参考Gin

#Day1 HTTP基础
1、简单介绍net/http库与http.Handler接口
2、搭建Gee框架雏形

#Day2 上下文Context
1、将路由(router)独立出来，方便之后增强。
2、设计上下文(Context)，封装 Request 和 Response ，提供对 JSON、HTML 等返回类型的支持。

#Day3 前缀树路由
1、使用 Trie 树实现动态路由(dynamic route)解析：支持两种模式:name和*filepath。

#Day4 分组控制
1、实现路由分组控制(Route Group Control)

#Day5 中间件

#Day6 模板（HTML Template）
1、实现静态资源服务(Static Resource)。
2、支持HTML模板渲染。