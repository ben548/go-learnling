package main

import (
	"encoding/json"
	"fmt"
	geerpc "geeRpc"
	"geeRpc/codec"
	"log"
	"net"
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
	conn, err := net.Dial("tcp", addrString)
	fmt.Println("err1", err)
	defer func() {
		err := conn.Close()
		fmt.Println("err2", err)
	}()

	//time.Sleep(time.Second)
	err = json.NewDecoder(conn).Decode(geerpc.DefaultServer)
	fmt.Println("err3", err)

	cc := codec.NewGobCodec(conn)

	for i := 0; i < 5; i++ {
		fmt.Println("i:", i)
		h := &codec.Header{
			ServiceMethod: "Foo.Sum",
			Seq:           uint64(i),
		}

		err = cc.Write(h, fmt.Sprintf("geerpc req d%", h.Seq))
		err = cc.Close()
		fmt.Println("err4", err)

		var reply string
		err = cc.ReadBody(&reply)
		fmt.Println("err5", err)

		fmt.Println("reply:", reply)
	}
}
