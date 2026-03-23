# GoChat

## Project Architecture

```bash
chat-system/
├── cmd/
│   └── server/
│       └── main.go          # entry point
│
├── internal/
│   ├── app/                 # application wiring (DI)
│   │   └── app.go
│   ├── gateway/             # WebSocket layer (real-time core)
│   │   ├── hub.go
│   │   ├── client.go
│   │   └── handler.go
│   ├── service/             # business logic
│   │   └── message_service.go
│   ├── repository/          # DB layer
│   │   └── message_repo.go
│   ├── pubsub/              # Redis abstraction
│   │   └── redis_pubsub.go
│   ├── model/               # domain models
│   │   └── message.go
│   └── transport/           # HTTP routing (if needed)
│       └── http.go
│
├── pkg/                     # reusable utilities (optional)
│   └── logger/
│
├── configs/
│   └── config.go
│
├── .env
├── go.mod
└── README.md
```

### Architecture Overview
- `gateway/`: Handles WebSocket connections and real-time messaging
- `service/`: Contains business logic
- `repository/`: Handles data persistence
- `pubsub/`: Enables horizontal scaling via Redis