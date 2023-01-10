# Jarvis

A Slack Bot built in Go

## Needed Slack App Permissions

- channels:read
- chat:write
- incoming-webhook
- users:read
- users:read.email

## Docker

```bash
docker build -t jarvis .     

docker run --name jarvis --env-file .env jarvis
```
