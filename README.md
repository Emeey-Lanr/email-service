# Email Service

Processes email notifications from the message queue and sends emails via SMTP or email service providers.

## Responsibilities

- Consume messages from `email.queue`
- Send emails via SMTP (SendGrid)
- Handle delivery confirmations and bounces
- Retry failed sends with exponential backoff
- Move permanently failed messages to dead-letter queue

## Tech Stack

- **Language**: Go
- **Message Queue**: RabbitMQ (amqp091-go)
- **Email Provider**:SendGrid


## Message Format

Messages consumed from `email.queue`:

```json
{
  "notification_id": "uuid",
  "user_id": "uuid",
  "email": "user@example.com",
  "template_code": "welcome_email",
  "variables": {
    "name": "John Doe",
    "link": "https://example.com",
    "meta": {}
  },
  "priority": 1,
  "request_id": "unique-request-id",
  "retry_count": 0
}
```

## Environment Variables

```env
RABBITMQ_URL=your-rabbitmq-url
SENDGRID_API_KEY=your-api-key
```

## Running Locally

```bash
git clone https://github.com/Emeey-Lanr/email-service.git
cd email-service
git mod tidy
go run main.go
```



## Retry Logic

1. Attempt fails → waits 10 seconds in 3 tries 
