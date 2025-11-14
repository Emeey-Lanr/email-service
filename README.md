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
  "correlation_id": "uuid",
  "data": {
    "template_code": "PASSWORD_RESET",
    "name": "Habeeb",
    "link": "https://hng.tech/reset/abc123"
  },
  "email": "habeeb@example.com",
  "subject": "Reset Your Password",
  "html_body": "<h1>Password Reset</h1><p>Hi Habeeb,</p><p>We received a request to reset your password. Click the link below to proceed:</p><p><a href=\"https://hng.tech/reset/abc123\">Reset Password</a></p>",
  "text_body": "Hi Habeeb, we received a request to reset your password. Use this link: https://hng.tech/reset/abc123"
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
