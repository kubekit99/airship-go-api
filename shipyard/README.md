# Use Swagger to generate Shipyard GO API code and simulator

```bash
swagger generate server --exclude-main -A shipyard -t . -f ./swagger/swagger.yml
swagger generate client -A shipyard -t . -f ./swagger/swagger.yml
```

Running the Config API Server

```bash
$ go run ./main.go
```
