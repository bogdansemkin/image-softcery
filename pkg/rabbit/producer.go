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

		 func(img []byte, filename string) {
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
