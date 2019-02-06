# Use Swagger to generate Drydock GO API code and simulator

```bash
swagger generate server --exclude-main -A drydock -t . -f ./swagger/swagger.yml
swagger generate client -A drydock -t . -f ./swagger/swagger.yml
```

Running the Config API Server

```bash
$ go run ./main.go
```
