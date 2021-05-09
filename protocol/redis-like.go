package protocol

import (
	"context"
	"flag"
	"github.com/Allenxuxu/gev"
	"github.com/Allenxuxu/gev/connection"
	"github.com/Allenxuxu/ringbuffer"
	"log"
	"redis-like/protocol/analysis"
	"strconv"
	"time"
)

type RedisExample struct {
}

func (s *RedisExample) OnConnect(c *connection.Connection) {
}

func (s *RedisExample) OnMessage(c *connection.Connection, ctx interface{}, data []byte) (out []byte) {
	d := time.Now().Add(1000 * time.Millisecond)
	cctx, closeFunc := context.WithDeadline(context.Background(), d)
	defer closeFunc()
	dealCmd, bs := analysis.RespProtocolAnalysis(cctx, data)
	if len(bs) != 0 {
		out = bs
		return
	}
	out = dealCmd.Deal()
	return
}

func (s *RedisExample) OnClose(c *connection.Connection) {
	log.Println("OnClose ：", c.PeerAddr())
	log.Println("============")
}

func (s *RedisExample) UnPacket(c *connection.Connection, buffer *ringbuffer.RingBuffer) (interface{}, []byte) {
	ret := buffer.Bytes()
	buffer.RetrieveAll()
	return nil, ret
}

func (s *RedisExample) Packet(c *connection.Connection, data []byte) []byte {
	return append(data, []byte("\r\n")...)
}

func Start() {

	handler := &RedisExample{}

	var port int
	var loops int

	flag.IntVar(&port, "port", 6379, "server port")
	flag.IntVar(&loops, "loops", 0, "num loops")
	flag.Parse()

	s, err := gev.NewServer(handler,
		gev.Network("tcp"),
		gev.Address(":"+strconv.Itoa(port)),
		gev.NumLoops(loops),
		gev.ReusePort(true),
		gev.Protocol(handler),
	)
	if err != nil {
		panic(err)
	}

	s.Start()
}