````markdown
# Message Microservices Project

A simple microservices project for sending and logging messages using **Go**, **RabbitMQ**, and **Python**, structured with Clean Architecture principles.

The project consists of:

1. **Go API Service** – exposes an HTTP endpoint to send messages.
2. **RabbitMQ** – acts as the message broker between services.
3. **Python Consumer** – listens to RabbitMQ messages and logs them.

## Requirements

- Go 1.25.1
- Docker (for RabbitMQ and optional Docker Compose)

## Running the Project

### Using Docker Compose

Run all services together with:

```bash
docker-compose up
```
````

This will start:

- Go API on `http://localhost:8080`
- RabbitMQ on `http://localhost:15672` (management UI)
- Python consumer logs in the console

### Running Without Docker Compose

1. Start RabbitMQ in Docker:

```bash
docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:4-management
```

2. Run the Go API:

```bash
go run ./app
```

3. Run the Python consumer:

```bash
python consumer.py
```

The Go API will be available on `http://localhost:8080`, and the Python consumer will log messages to the console.

## Go API Endpoints

### POST /messages

Send a message.

**Request Body**:

```json
{
  "author": "Your Name",
  "message": "Your message"
}
```

**Response**:

- `200 OK` — message sent successfully
- `400 Bad Request` — validation failed
- `500 Internal Server Error` — publishing error
