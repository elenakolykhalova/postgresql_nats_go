package main

import (
	"encoding/json"
	"fmt"
	"L0/repository"
	"L0/valid"
	"github.com/nats-io/stan.go"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var Cache map[string]string

func main() {
	Cache = make(map[string]string)
	repository.FillCache(Cache)
	natsConn := StanConnect("test-cluster", "client-new", "")
	defer natsConn.Close()
	subcr, err := natsConn.Subscribe("superTrouper", StMsgHandler)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("Subscription is succeeded...")
	}
	defer subcr.Unsubscribe()
	WebServer()
}


func StMsgHandler(m *stan.Msg) {
	var MyModel repository.Order
	isMessageValidJSON := valid.ValidateMyJSON(m.Data)
	fmt.Println("New message is valid ", isMessageValidJSON == 1)
	if isMessageValidJSON == 1 {
		json.Unmarshal(m.Data, &MyModel)
		MyModel.Data = string(m.Data)
		if _, ok := Cache[MyModel.Order_uid]; !ok {
			db, err := repository.ConnectToDB()
			if err != nil {
				log.Fatal(err)
			}
			defer db.Close()
			repository.AddOrderTx(db, MyModel)
			Cache[MyModel.Order_uid] = string(m.Data)
			fmt.Println("Added new order with id: ", MyModel.Order_uid)
		} else {
			fmt.Println("\"" + MyModel.Order_uid + "\"" + "  order_uid is already in DB")
		}
	}
}

func StanConnect(cluster, client, url string) stan.Conn {
	sc, err := stan.Connect(
		cluster,
		client,
		stan.Pings(1, 3),
		stan.NatsURL(""),
	)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Connected to \"%s\" as \"%s\"...\n", cluster, client)
	return sc
}


func WebServer() {
	ws := http.NewServeMux()
	ws.HandleFunc("/", showMain)
	ws.HandleFunc("/id", showOrder)
	http.ListenAndServe("127.0.0.1:8080", ws)
}

func showOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	answer, ok := Cache[id]
	if ok {
		w.Write([]byte(answer))
	} else {
		bad_ans := "{\"order_uid\": \"" + id + " is nil\"}"
		w.Write([]byte(bad_ans))
	}
}

func showMain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	htmlMain, err := os.Open("html/index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer htmlMain.Close()
	htmlMainBytes, err := ioutil.ReadAll(htmlMain)
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(htmlMainBytes))
}
