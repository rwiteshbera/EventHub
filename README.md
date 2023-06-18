# EventHub
![eventHub.png](assets%2FeventHub.png)
EventHub is an event management platform designed to simplify the organization and management of upcoming virtual or in-person meetups, conferences, and competitions. This project follows a microservices architecture, leveraging various technologies and services to provide a scalable and efficient solution. The project is still under development, with future updates planned to include the addition of an API gateway.

## Tech Stack
EventHub is built using the following technologies:
1. Go: The core programming language used for developing the microservices.
2. Redis: A fast in-memory data store used for caching and session management.
3. RabbitMQ: A message broker responsible for handling communication between services.
4. gRPC: A high-performance, open-source framework for remote procedure calls used for inter-service communication.
5. MongoDB: A NoSQL database used for storing user information and event data.
6. Postgres: A relational database used for storing additional data related to events.
7. Docker: A containerization platform used for packaging the microservices and their dependencies.

## Services
The EventHub platform consists of three distinct services:
### 1. User Management Service 
The User Management service is responsible for authentication and authorization. When a user signs up, they provide their email address, which triggers the generation and temporary storage of an OTP (One-Time Password). This OTP is then sent to RabbitMQ, which acts as a message broker. The Email Service consumes the OTP message from RabbitMQ and sends it to the user's email address. The user verifies their account by providing the received OTP. If the OTP is valid, a new user is created in MongoDB.

### 2. Email Service
The Email Service consumes OTP messages from RabbitMQ and sends them to the user's email address. It plays a vital role in the user verification process, ensuring secure and reliable communication.

### 3. Event Catalog Service
The Event Catalog Service is responsible for managing and displaying all events on the platform. Organizers can create events by sending gRPC requests to this service. The service validates the request, checking whether the user is a valid user, and then either creates a new event or rejects the request.

## Future Updates
The project is still actively being developed, and future updates will include the addition of an API gateway. This gateway will serve as a single entry point for all incoming requests and will provide features such as request routing, rate limiting, and authentication.

## Setup
1. EventHub relies on several services, including RabbitMQ, Redis, MongoDB, and Postgres. To start these services, use Docker Compose with the provided configuration files:

```bash
docker-compose -f broker.yaml -f database.yaml up
```
2. If you want to run a specific Gin server for a particular service, navigate to the respective service's directory and execute the following command. Make sure you provided the environment variable to run the service properly.
```bash
cd user-management-service
go run ./cmd/main.go
``` 
Replace user-management-service with the directory name of the service you want to run. Repeat this step for any other services you wish to start.

RabbitMQ Administrator Default Credential
```text
username: guest
password: guest
```


