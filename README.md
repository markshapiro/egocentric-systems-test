
run project with `go run main.go`

sample requests:

```
POST http://localhost:8080/add
POST http://localhost:8080/subtract
POST http://localhost:8080/multiply
POST http://localhost:8080/divide

{
  "operandA":"3434.21334",
  "operandB":"2133.23423"
}
```

get last 5:

```
GET http://localhost:8080/getRecentN?n=7
```