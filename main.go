package main

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/url"
	"time"
)
type Messages struct {
Message string `json:message`
Name string `json:name`
}
func main() {
	uri, err := url.Parse(" test/topic")
	if err != nil {
		log.Fatal(err)
	}
	topic := uri.Path[1:len(uri.Path)]
	if topic == "" {
		topic = "test"
	}

	go listen(uri, topic)

	client := connect("pub", uri)

	timer := time.NewTicker(1 * time.Second)
	for _ = range timer.C {
		client.Publish(topic, 0, false, "")
	}

}

func connect(clientId string, uri *url.URL) mqtt.Client {
	opts := createClientOptions(clientId, uri)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}

func createClientOptions(clientId string, uri *url.URL) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", ":1883"))
	opts.SetUsername(uri.User.Username())
	password, _ := uri.User.Password()
	opts.SetPassword(password)
	opts.SetClientID(clientId)
	return opts
}
func listen(uri *url.URL, topic string) {
	client := connect("sub", uri)
	client.Subscribe(topic, 1, func(client mqtt.Client, msg mqtt.Message) {
		if string(msg.Payload()) != "" {
			m:=new(Messages)
			json.Unmarshal(msg.Payload(), &m)

			SaveMessage(m.Message, m.Name)
		}
	})
}
