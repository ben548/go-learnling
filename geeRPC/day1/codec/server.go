package codec

import (
	"fmt"
	"io"
	"log"
	"net"
	"reflect"
	"sync"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

var DefaultServer = NewServer()

func (server *Server) Accept(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("rpc server accept error", err)
			return
		}
		//go server.()
	}
}

func Accept(listener net.Listener) {
	DefaultServer.Accept(listener)
}

func (server *Server) ServerConn(conn io.ReadWriteCloser) {
	defer func() { _ = conn.Close() }()

	var opt Option
}

func (server *Server) ServerCodec(cc Codec) {
	sending := new(sync.Mutex)
	wg := new(sync.WaitGroup)

}

type Request struct {
	h            *Header
	argv, replyv reflect.Value
}

func (server *Server) ReadRequestHeader(cc Codec) (*Header, error) {
	var h Header
	if err := cc.ReadHeader(&h); err != nil {
		if err != io.EOF && err != io.ErrUnexpectedEOF {
			log.Println("rpc server: read header error:", err)
		}
		return nil, err
	}
	return &h, nil
}

func (server *Server) readRequest(cc Codec) (*Request, error) {
	h, err := server.ReadRequestHeader(cc)
	if err != nil {
		return nil, err
	}

	req := &Request{
		h: h,
	}

	req.argv = reflect.New(reflect.TypeOf(""))
	if err = cc.ReadBody(req.argv.Interface()); err != nil {
		log.Println("rpc server: read argv err:", err)
	}
	return req, nil
}

func (server *Server) SendResponse(cc Codec, h *Header, body interface{}, sending *sync.Mutex) {
	sending.Lock()
	defer sending.Unlock()
	if err := cc.Write(h, body); err != nil {
		log.Println("rpc server: write response error:", err)
	}
}

func (server *Server) handleRequest(cc Codec, req *Request, sending *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	req.replyv = reflect.ValueOf(fmt.Sprintf("geerpc resp %d", req.h.Seq))
	server.SendResponse(cc, req.h, req.replyv.Interface(), sending)
}
