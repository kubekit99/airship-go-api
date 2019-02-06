# Use Swagger to generate Armada GO API code and simulator

```bash
swagger generate server --exclude-main -A armada -t . -f ./swagger/swagger.yml
swagger generate client -A armada -t . -f ./swagger/swagger.yml
```

Running the Config API Server

```bash
$ go run ./main.go
```
