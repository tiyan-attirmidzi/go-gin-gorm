# GO REST API

This is a Rest API with [GO](https://go.dev/) using [GIN](https://github.com/gin-gonic/gin) framework and ORM with [GORM](https://gorm.io/docs/index.html), database using `MySQL`.

## Getting Started

Get All Dependency


```bash
go mod tidy
```

Create `.env` based on `.env.example` then complete the variables

```bash
cp .env.example .env
```

Run the development server:

```bash
go run main.go
```