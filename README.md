# Go-Morning
Send good morning and good evening messages to your Telegram's Family group! ❤️

## Requirements
- Golang
- MongoDB

### Run

Copy `.env.example` with the name `.env`, add your secrets, then run:

```bash
make run
make build
make clean
```

### Endpoints
Endpoints to send a message:

Example:
```bash
curl -X POST localhost:8080/morning
curl -X POST localhost:8080/evening
```

______
Copyright Daniele Boscolo.
