package rabbit

import (
	"fmt"
	"github.com/nfnt/resize"
	"github.com/streadway/amqp"
	"image/jpeg"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const url = "amqp://guest:guest@localhost:5672/"

func (rabbit *MQ) Consumer() (string, string, string, string) {
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
	c := make(chan string, 4)
	go func() {
		for d := range msgs {

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

			img, err := jpeg.Decode(file)
			if err != nil {
				log.Fatal(err)
			}
			file.Close()

			m := resize.Resize(1280, 960, img, resize.Lanczos3)

			out, err := os.Create(path)
			if err != nil {
				log.Fatal(err)
			}

			err = jpeg.Encode(out, m, nil)
			if err != nil {
				log.Fatal(err)
			}
			if out.Name() != "" {
				c <- out.Name()
				c <- imageResize(out.Name())
				c <- imageHalfResize(out.Name())
				c <- imageFullResize(out.Name())
			}
			fmt.Printf("%s : OK\n", path)
			time.Sleep(time.Second * 5)
		}
	}()
	image := <-c
	seventy_five_image := <-c
	half_image := <-c
	twenty_five_image := <-c

	return image, seventy_five_image, half_image, twenty_five_image
}
