# PRODUCT STOCK MICROSERVICE

## Overview

GO application for keeping track of the stock quantity of a product.
The microservice has been build following the DDD design pattern.

Endpoints:

- /health (GET)
- /api/v1/product/:productID (GET)
- /api/v1/product/ (POST)
- /api/v1/product/ (PUT)
- /api/v1/product/:productID (DELETE)

### Installation

```sh
# go 1.18+
go run main.go
```

### Test

```sh
go test -v .\test\
```

### Docker

```sh
docker-compose up --build
```
