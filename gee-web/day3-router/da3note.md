#Day3 前缀树路由
1、使用 Trie 树实现动态路由(dynamic route)解析：支持两种模式:name和*filepath。

动态路由，即一条路由规则可以匹配某一类型而非某一条固定的路由。例如/hello/:name，可以匹配/hello/geektutu、hello/jack等。
开源的路由实现gorouter支持在路由规则中嵌入正则表达式，另一个开源实现httprouter就不支持正则表达式。而gin早起使用httprouter，后来自己实现。

前缀树（Trie），实现动态路由最常用的数据结构：每一个节点的所有的子节点都拥有相同的前缀。
![img.png](img.png)