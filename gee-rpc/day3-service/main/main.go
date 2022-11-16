package main

import (
	"geerpc"
	"log"
	"net"
	"sync"
	"time"
)

type Foo int

type Args struct{ Num1, Num2 int }

func (f Foo) Sum(args Args, reply *int) error {
	*reply = args.Num1 + args.Num2
	return nil
}

func startServer(addr chan string) {
	var foo Foo
	if err := geerpc.Register(&foo); err != nil {
		log.Fatal("register error:", err)
	}
	// pick a free port
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("network error:", err)
	}
	log.Println("start rpc server on", l.Addr())
	addr <- l.Addr().String()
	geerpc.Accept(l)
}

func main() {
	log.SetFlags(0)
	addr := make(chan string)
	go startServer(addr)
	client, _ := geerpc.Dial("tcp", <-addr)
	defer func() { _ = client.Close() }()

	time.Sleep(time.Second)
	// send request & receive response
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			args := &Args{Num1: i, Num2: i * i}
			var reply int
			if err := client.Call("Foo.Sum", args, &reply); err != nil {
				log.Fatal("call Foo.Sum error:", err)
			}
			log.Printf("%d + %d = %d", args.Num1, args.Num2, reply)
		}(i)
	}
	wg.Wait()
}

//// 通过反射，我们能够非常容易地获取某个结构体的所有方法，并且能够通过方法，获取到该方法所有的参数类型与返回值
//func main() {
//	var wg sync.WaitGroup
//	typ := reflect.TypeOf(&wg)
//	for i := 0; i < typ.NumMethod(); i++ {
//		method := typ.Method(i)
//		argv := make([]string, 0, method.Type.NumIn())
//		returns := make([]string, 0, method.Type.NumOut())
//		// j 从 1 开始，第 0 个入参是 wg 自己。
//		for j := 1; j < method.Type.NumIn(); j++ {
//			argv = append(argv, method.Type.In(j).Name())
//		}
//		for j := 0; j < method.Type.NumOut(); j++ {
//			returns = append(returns, method.Type.Out(j).Name())
//		}
//		log.Printf("func (w *%s) %s(%s) %s",
//			typ.Elem().Name(),
//			method.Name,
//			strings.Join(argv, ","),
//			strings.Join(returns, ","))
//	}
//}
