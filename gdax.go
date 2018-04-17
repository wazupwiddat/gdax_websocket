package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

type gdax struct {
	conn *websocket.Conn
}

func (g *gdax) connect(addr string) error {
	u := url.URL{Scheme: "wss", Host: addr, Path: "/"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	g.conn = c
	return err
}

func (g gdax) subscribe(s *subscribe) error {
	if g.conn == nil {
		return errors.New("Need to make a connections before subscribing")
	}

	b, err := json.Marshal(s)
	log.Println("Suscbribing to:", s)
	if err != nil {
		log.Println("json:", err)
		return err
	}

	err = g.conn.WriteMessage(websocket.TextMessage, b)
	if err != nil {
		log.Println("write:", err)
		return err
	}
	return nil
}

func (g gdax) startListening(gk *GdaxKinesis) {
	defer g.conn.Close()
	done := make(chan struct{})

	// Read message from GDAX
	go func() {
		defer close(done)
		for {
			_, message, err := g.conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			//log.Println(message)
			gk.writeMessage(message)
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// Wait on the interrupt channel
	for {
		select {
		case <-done:
			return
		case <-interrupt:
			log.Println("Interrupted")
			err := g.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}

}
