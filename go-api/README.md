# Message API

A simple Go API for sending messages via RabbitMQ, structured using Clean Architecture principles.

## Requirements

- Go 1.25.1
- RabbitMQ running locally at `amqp://guest:guest@localhost:5672/`

## Running the Application

From the project root:

```bash
go run ./app
```

The server will start on `http://localhost:8080`.

## Endpoints

### POST /messages

Send a message:

**Request Body**

```json
{
  "author": "Your Name",
  "messsage": "Your message"
}
```

**Response**

- `200 OK` if the message was sent successfully
- `400 Bad Request` if validation fails
- `500 Internal Server Error` on publishing errors
