# full-cycle-2.0-hexagonal-architecture

Files I produced during the Hexagonal Architecture classes of my [Microservices Full Cycle 2.0 course](https://drive.google.com/file/d/1MdN-qK_8Pfg6YI3TSfSa5_2-FHmqGxEP/view?usp=sharing).

## Installing dependencies

```sh
go get github.com/asaskevich/govalidator
go get github.com/golang/mock
go get github.com/satori/go.uuid
go get github.com/stretchr/testify
go get github.com/mattn/go-sqlite3
go get github.com/urfave/negroni
go get github.com/gorilla/mux
go install github.com/golang/mock/mockgen@v1.5.0
go install github.com/spf13/cobra/cobra@latest
# make sure to add $HOME/go/bin to your path

sudo apt-get update
sudo apt-get install -y sqlite3
```

## Generate mocks

```sh
mockgen -destination=application/mocks/application.go -source=application/product.go application
```

## Adding commands with cobra

```sh
cobra add cli  # Creates the `go run main.go cli` command
cobra add http  # Creates the `go run main.go http` command
```

## Running the application

```sh
go run main.go cli --help
productid=$(go run main.go cli --action create --product Mouse --price 27.8 | cut -d ' ' -f 3)
go run main.go cli --action get --id $productid

go run main.go http --help
go run main.go http  # Start the web server in another terminal

# Finds the created product
curl http://localhost:8080/product/$productid

# Creates a product
curl -X POST -d '{ "name": "Mouse", "price": 27.8 }' http://localhost:8080/product

# Forces a deserialization error
curl -X POST -d '{ "name": "Mouse", "price": "WEIRD PRICE" }' http://localhost:8080/product

# Forces a business rule error
curl -X POST -d '{ "name": "Mouse", "price": -100 }' http://localhost:8080/product

# Lists all products. Must print 2 products.
sqlite3 cli_and_web.sqlite3 "select * from products"

# Enables the product created with CLI command
curl -X PATCH http://localhost:8080/product/$productid/enable

go run main.go cli --action get --id $productid

# Disables the product created with CLI command
curl -X PATCH http://localhost:8080/product/$productid/disable

go run main.go cli --action get --id $productid
```

## Testing the application

```sh
go test ./...
```
