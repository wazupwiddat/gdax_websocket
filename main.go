package main

import (
	"flag"
	"log"
)

var (
	addr       = flag.String("addr", "ws-feed.gdax.com", "")
	awsProfile = flag.String("profile", "jdub", "AWS Access Key Profile")
	stream     = flag.String("stream", "gdax-websocket", "your stream name")
	region     = flag.String("region", "us-east-1", "your AWS region")
)

func main() {
	flag.Parse()
	log.SetFlags(0)

	gk := NewKinesisStream(*stream, *region)

	g := connect(*addr)
	sendSubscribe("BTC-USD", g)
	g.startListening(gk)

	log.Println("finished")
}

func connect(u string) gdax {
	g := gdax{}
	err := g.connect(u)
	if err != nil {
		log.Fatal("connect:", err)
	}
	return g
}

func sendSubscribe(p string, g gdax) {
	s := &subscribe{
		MsgType:    "subscribe",
		ProductIds: []string{},
	}
	s.ProductIds = append(s.ProductIds, p)
	err := g.subscribe(s)
	if err != nil {
		log.Fatal("subscribe:", err)
	}
}
