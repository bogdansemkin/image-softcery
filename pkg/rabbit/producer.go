package rabbit

import (
	"fmt"
	"github.com/streadway/amqp"
	"image-softcery/pkg/services"
	"io/ioutil"
	"log"
	"os"
)
type MQ struct{
	service *services.Service
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func (rabbit *MQ) Producer(filepath string){
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("Failed to connect to server. Error: %s", err.Error())
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel. Error: %s", err.Error())
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("resizing", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare a queue. Error: %s", err.Error())
	}

	files, err := ioutil.ReadDir("D:\\image-softcery\\templates\\img")
	if err != nil {
		log.Fatalf("Failed to open dir. Error: %s", err.Error())
	}

	for _, f := range files {
		// opening the image
		file, err := os.Open(filepath)
		if err != nil {
			log.Fatalf("Failed to open file. Error: %s", err.Error())
		}

		img, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()

		go func(img []byte, filename string) {
			// Create a message
			msg := amqp.Publishing{
				Body: img,
			}

			// Publishing
			ch.Publish("", q.Name, false, false, msg)
			fmt.Printf("%s in QUEUE\n", filename)
		}(img, f.Name())
	}
}

//func (rabbit *MQ) Producer(filepath string) {
//	conn, err := amqp.Dial(url)
//	failOnError(err, "Failed to connect to RabbitMQ")
//	defer conn.Close()
//
//	ch, err := conn.Channel()
//	failOnError(err, "Failed to open a channel")
//	defer ch.Close()
//
//	q, err := ch.QueueDeclare(
//		"hello", // name
//		false,   // durable
//		false,   // delete when unused
//		false,   // exclusive
//		false,   // no-wait
//		nil,     // arguments
//	)
//	failOnError(err, "Failed to declare a queue")
//
//	body := "Hello World!"
//	err = ch.Publish(
//		"",     // exchange
//		q.Name, // routing key
//		false,  // mandatory
//		false,  // immediate
//		amqp.Publishing{
//			ContentType: "text/plain",
//			Body:        []byte(body),
//		})
//	failOnError(err, "Failed to publish a message")
//	log.Printf(" [x] Sent %s\n", body)
//}