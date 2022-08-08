# image-softcery

This is a HTTP API of a upload-download image application.

This API based on Uncle's Bob architecture: 
![6894a85f-6047-4b4a-bdf3-6ea0b80c1e31](https://user-images.githubusercontent.com/40574816/183291225-67575c43-7fdc-475f-b5e7-a2e0fb3bd2de.jpg)

The entire application is contained within the `cmd/main.go` file.

`config.yml` is a minimal configuration for postgres.

## Clone

    git clone github.com/bogdansemkin/image-softcery

## Run the tests

    go run ./pkg/handlers/image_test.go
    
## Run Redis:

    docker run -d --name redis-stack-server -p 6379:6379 redis/redis-stack-server:latest
    
## Run DB:
   
    For start db use docker-compose -> image-softcery.db1
    
## Create table on DB: 

    migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up

# HTTP API

The API to the example app is described below.

## Upload image

 ### Request

`POST /images/`

    curl -v -F  upload=@localfilename http://8000/images/upload


### Response

    HTTP/1.1 200 OK
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 200 OK
    Connection: close
    Content-Type: multipart/form-data
    Content-Length: 2

    []
    
 ## Download image
 
  ### Request

`GET /images/{id}`

    curl -v -F  download=@localfilename http://8000/images/download/{id}


### Response

    HTTP/1.1 200 OK
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 200 OK
    Connection: close
    Content-Type: multipart/form-data
    Content-Length: 2

    {"imageFile": *multipart.FileHeader}


## P.S:
![image](https://user-images.githubusercontent.com/40574816/183312542-aded2b91-0ace-49f1-8a3d-4512ef8d3155.png)

At first, I thought of leaving business-entity at the root of the application so that we can use it at any level of the application without any problems. But for the sake of the beauty of the code, I decided to throw them into pkg

### For reviewer

Задача состояла в том, что нужно оптимизировать фотографии и делать ресайз под входящий аргумент на эндпоинте /download/:id. 
Да, изначально в голове была конструкция сделать динамический ресайз, который будет отталкиваться от входящего размера, в дальнейшем оптимизироваться(если image > 1920x1080, оно оптимизируется в full-hd, если же ниже, то нет).

Но, при написании кода я решил уйти немного в сторону.

![image](https://user-images.githubusercontent.com/40574816/183419718-da226b46-8f0d-4bc5-9856-fa913734090b.png)

Яркий пример: Для тех или иных обстоятельств будет подтягиваться в дальнейшем с сервиса то разрешение, которое нужно. В свою очередь, любое фото, загружаемое на сервер, оптимизируется до *минимального* 16х9(Минимальный в том плане, что фото сжато и занимает мало места, но оно всё так же четко видно. Очевидно, что можно сделать ещё меньше).

Да, это в идейном плане отличается от ТЗ, но сам функционал(оптимизация и ресайз) реализован, а я просто попытался объяснить почему и зачем статик, кхэ.

***About code

Сыграв на опережение, есть несколько участков кода, которые для меня не идеальны и я бы так не делал, будь возможность довести до идеала. 

Первый, и самый явный для меня, который портит восприятие, это:

`./pkg/handlers/image.go
	file, err := c.FormFile("imageFile")`
    
В коде есть модель, которую можно использовать, а тут явная тянучка по имени файла, что, как по мне, не есть хорошо.

Второй:

`	c := make(chan string, 4)
...
    image := <-c
	seventy_five_image := <-c
	half_image := <-c
	twenty_five_image := <-c
`

Считывание с канала. 

Третий:

Писал доку изначально, использовав swagger@v1.6.7, но, позже решил его апдейтнуть до @latest и начал ловить ошибку, дескать swag не видит .go файлов. Прочитал в одном из issue, что граммотным решением будет удалить старую доку и сделать `swag init -g` на новой версии. В итоге, все мои `paths` c `./pkg/handlers/` были утеряны. 
