# full-cycle-2.0-hexagonal-architecture

Files I produced during the Hexagonal Architecture classes of my [Microservices Full Cycle 2.0 course](https://drive.google.com/file/d/1MdN-qK_8Pfg6YI3TSfSa5_2-FHmqGxEP/view?usp=sharing).

## Installing dependencies

```sh
go get github.com/asaskevich/govalidator
go get github.com/golang/mock
go get github.com/satori/go.uuid
go get github.com/stretchr/testify
go get github.com/mattn/go-sqlite3
go install github.com/golang/mock/mockgen@v1.5.0

sudo apt-get update
sudo apt-get install -y sqlite3
```

## Generate mocks

```sh
mockgen -destination=application/mocks/application.go -source=application/product.go application
```

## Testing the application

```sh
go test ./...
```
