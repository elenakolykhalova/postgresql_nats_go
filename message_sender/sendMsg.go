package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	sc, err := stan.Connect(
		"test-cluster",
		"client_send",
		stan.Pings(1, 3),
		stan.NatsURL(strings.Join(os.Args[1:], ",")),
	)
	if err != nil {
		log.Fatalln(err)
	}
	defer sc.Close()

	sub, err := sc.Subscribe("superTrouper", func(m *stan.Msg) {
	})
	if err != nil {
		log.Fatalln(err)
	}
	defer sub.Unsubscribe()
	jsonFile, err := os.Open("message_sender/validJsonTemplates/model2.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	byteValJson, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	if err := sc.Publish("superTrouper", byteValJson); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Message is sent!")
}
