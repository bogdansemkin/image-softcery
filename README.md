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
