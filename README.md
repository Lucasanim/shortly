# Link Shortening Service

A high-performance URL shortening service built with Go, Fiber, DynamoDB, and Redis for caching.

## Features
- Generate short links from long URLs
- Redirect users to the original URL
- Cache short links using Redis for fast lookups
- Store links persistently in DynamoDB
- API-based service with Fiber framework

## Technologies Used
- **Go**: Backend service
- **Fiber**: Fast web framework for Go
- **DynamoDB**: NoSQL database for persistent storage
- **Redis**: Cache for faster link resolution
- **Docker**: Containerized development setup

## Getting Started

### Prerequisites
Ensure you have the following installed:
- [Go](https://go.dev/dl/)
- [Docker & Docker Compose](https://docs.docker.com/get-docker/)
- [AWS session file](https://docs.aws.amazon.com/sdk-for-go/api/aws/session/)

### Setup & Run
Clone the repository:
```sh
git clone git@github.com:Lucasanim/shortly.git
cd shortly
```

Start the dependencies (DynamoDB & Redis):
```sh
./run.sh
```

This script will:
1. Start DynamoDB and Redis using Docker Compose.
2. Ensure DynamoDB and Redis are ready before launching the Go app.
3. Run the Go application.

### API Endpoints

#### 1. Shorten a URL
**POST** `/app/register`
```json
{
  "url": "https://example.com"
}
```
Response:
```json
{
  "message": "Link created",
  "shortened_link": "http://localhost:3000/r/abc123"
}
```

#### 2. Redirect to the Original URL
**GET** `/r/:hash`
- Redirects the user to the original URL.

### Testing
Run unit tests:
```sh
go test ./...
```

### License
This project is licensed under the MIT License. For more information read LICENSE file.

