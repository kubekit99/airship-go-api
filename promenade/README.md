# Use Swagger to generate Promenade GO API code and simulator

```bash
swagger generate server --exclude-main -A promenade -t . -f ./swagger/swagger.yml
swagger generate client -A promenade -t . -f ./swagger/swagger.yml
```

Running the Config API Server

```bash
$ go run ./main.go
```
