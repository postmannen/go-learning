package main

import (
	"fmt"
	"log"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// NewMQTTClient if successful will return a new client
// connection to the broker with options set
func NewMQTTClient(protocol string, address string, port string, clientID string) (mqtt.Client, error) {
	// -- Set default connect options
	// mqtt.DEBUG = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)
	opts := mqtt.NewClientOptions().AddBroker(protocol + "://" + address + ":" + port).SetClientID(clientID)
	opts.SetKeepAlive(2 * time.Second)
	opts.SetPingTimeout(1 * time.Second)

	// Create the connection to the broker
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return c, nil
}

func subscribe(c mqtt.Client, topic string) error {
	// Since the Subscribe method uses a callback function
	// for what to do with the message, we declare such a
	// method to print out the messages we receive.
	messageHandler := func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("*** Got message *** [%s] %s\n", msg.Topic(), string(msg.Payload()))
	}

	// Start the consuming of the topic
	if token := c.Subscribe(topic, 0, messageHandler); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}

func publish(c mqtt.Client, topic string, msgCh chan interface{}) {
	for {
		// Create a string with the data to publish to broker
		text := fmt.Sprintf("this is msg #%v!", <-msgCh)
		token := c.Publish(topic, 0, false, text)
		token.Wait()
	}
}

// unSubscribe will unsubscribe the client from the topic
func unSubscribe(c mqtt.Client, topic string) error {
	if token := c.Unsubscribe(topic); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func main() {
	var err error
	topic := "CloudBoundContainer"

	// Create new mqtt client
	client, err := NewMQTTClient("tcp", "10.0.0.26", "1883", "btclient")
	if err != nil {
		log.Printf("error: newMQTTClient failed: %v\n", err)
		return
	}
	defer client.Disconnect(250)

	// Subscribe to topic,
	// subscribe will also print the result to console.
	err = subscribe(client, topic)
	if err != nil {
		log.Printf("error: mqtt client subscribe failed: %v\n", err)
		return
	}

	msgCh := make(chan interface{})

	//start publish'er
	go publish(client, topic, msgCh)

	// send some data to the publisher channel
	for i := 0; i < 10; i++ {
		msgCh <- i
		time.Sleep(time.Millisecond * 300)
	}

	err = unSubscribe(client, topic)
	if err != nil {
		log.Printf("error: mqtt client unSubscribe failed: %v\n", err)
		return
	}

}
