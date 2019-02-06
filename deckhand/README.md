# Use Swagger to generate Deckhand GO API code and simulator

```bash
swagger generate server --exclude-main -A deckhand -t . -f ./swagger/swagger.yml
swagger generate client -A deckhand -t . -f ./swagger/swagger.yml
```

Running the Config API Server

```bash
$ go run ./main.go
```
