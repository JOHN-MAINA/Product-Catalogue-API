# Product-Catalogue-API
The product management catalogue is a simple application to keep track of products.

##Dependencies

1. github.com/gorilla/mux

1. github.com/go-ozzo/ozzo-validation

1. github.com/go-ozzo/ozzo-validation/is

1. github.com/gorilla/handlers

1. github.com/rs/cors

1. github.com/jinzhu/gorm/dialects/mysql

1. github.com/jinzhu/gorm

##Installation Instructions

1. clone the repository using

    `git clone https://github.com/JOHN-MAINA/Product-Catalogue-API.git`

1. install dependencies

    `go get -u <dependecies name listed above>`

1. Configure database credential in /config/config.go

1. Set port number that the application will be served through default is set to `3001`

1. Serve the application via executing the executable or go run
    1. Create an executable 
    
        `go build`
        
    1. Server Via go run
    
        `go run main.go`
