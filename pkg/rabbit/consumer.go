package rabbit

import (
	"fmt"
	"github.com/nfnt/resize"
	"github.com/streadway/amqp"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"math/rand"
)

const url = "amqp://guest:guest@localhost:5672/"

func (rabbit *MQ) Consumer(){
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

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for d := range msgs {
			// create a random name for the new image
			s1 := rand.NewSource(time.Now().Unix())
			r1 := rand.New(s1)
			name := strconv.Itoa(r1.Intn(100000))
			path := "D:\\image-softcery\\templates\\img\\" + name + ".png"

			err := ioutil.WriteFile(path, d.Body, 0644)
			if err != nil {
				log.Fatal(err)
			}

			file, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}

			// Deserialize file
			img, err := jpeg.Decode(file)
			if err != nil {
				log.Fatal(err)
			}
			file.Close()

			// resizing
			m := resize.Resize(500, 500, img, resize.Lanczos3)

			// create the file
			out, err := os.Create(path)
			if err != nil {
				log.Fatal(err)
			}

			// write new image to file
			err = jpeg.Encode(out, m, nil)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s : OK\n",path)
			time.Sleep(time.Second*5)
		}
	}()

}

//func (rabbit *MQ) Consumer() {
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
//	msgs, err := ch.Consume(
//		q.Name, // queue
//		"",     // consumer
//		true,   // auto-ack
//		false,  // exclusive
//		false,  // no-local
//		false,  // no-wait
//		nil,    // args
//	)
//	failOnError(err, "Failed to register a consumer")
//
//	go func() {
//		for d := range msgs {
//			log.Printf("Received a message: %s", d.Body)
//		}
//	}()
//
//	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
//
//}