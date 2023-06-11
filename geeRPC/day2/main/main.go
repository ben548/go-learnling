package main

import (
	"fmt"
	geerpc "geeRpc"
	"log"
	"net"
	"sync"
	"time"
)

func startServer(addr chan string) {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("network error:", err)
	}
	log.Println("start rpc server on", l.Addr())
	fmt.Println(l.Addr().String())
	addr <- l.Addr().String()
	geerpc.Accept(l)
}

func main() {
	addr := make(chan string)
	go startServer(addr)
	addrString := <-addr
	fmt.Println("addrString:", addrString)
	client, err := geerpc.Dial("tcp", addrString)
	//conn, err := net.Dial("tcp", addrString)
	fmt.Println("err1", err)
	defer func() {
		err := client.Close()
		fmt.Println("err2", err)
	}()

	//time.Sleep(time.Second)
	//err = json.NewDecoder(conn).Decode(geerpc.DefaultServer)
	//fmt.Println("err3", err)
	//
	//cc := codec.NewGobCodec(conn)
	time.Sleep(time.Second)
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		fmt.Println(i)
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			args := fmt.Sprintf("geerpc req %d", i)
			var reply string
			if err := client.Call("Foo.Sum", args, &reply); err != nil {
				log.Fatal("call Foo.Sum err:", err)
			}
			log.Println("reply:", reply)
		}(i)
	}
	wg.Wait()
}
