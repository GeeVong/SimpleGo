package main

import (
	"fmt"
	amqp "github.com/streadway/amqp"
	"log"
	"strconv"

	"time"
)

func main() {
	generateMessage()
}

func generateMessage() {
	// 连接RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if nil != err {
		log.Fatal(err)
	}
	defer conn.Close()

	// 创建信道
	ch, err := conn.Channel()
	defer ch.Close()

	// 声明要操作的队列
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// 发送消息
	//singleMessage(ch, q, 2)

	for i := 10; i > 0; i-- {
		time.Sleep(time.Second)
		singleMessage(ch, q, i)
	}

}

func singleMessage(ch *amqp.Channel, q amqp.Queue, i int) {
	body := "Hello rabbitMQ!" + " " + strconv.FormatInt(time.Now().Unix(), 10) + " == " + strconv.Itoa(i)
	err := ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
	fmt.Println("=======i:", i)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
