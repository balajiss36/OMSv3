# Order Management System

This is an Order Management System built using Golang. It leverages gRPC for internal communication, RabbitMQ for asynchronous operations, and MongoDB as the database. The system is built on a microservice architecture with four services: Gateway, Orders, Kitchen, and Stock.

## Features

- **Golang**: The core of the application is built using Golang.
- **gRPC**: Used for efficient internal communication between services.
- **RabbitMQ**: Handles asynchronous operations and messaging.
- **MongoDB**: Stores order data and other related information.
- **Microservices**: Four services - Gateway, Orders, Kitchen, and Stock.

## Prerequisites

- Go 1.22 or higher
- MongoDB
- RabbitMQ
- Docker

## Installation

1. **Clone the repository:**

    ```sh
    git clone https://github.com/yourusername/order-management-system.git
    cd order-management-system
    ```

2. **Install dependencies:**

    ```sh
    go mod tidy
    ```

3. **Set up environment variables:**

    Create a `.env` file in the root directory and add the following variables:

    ```env
    MONGODB_URI=mongodb://localhost:27017
    RABBITMQ_URI=amqp://guest:guest@localhost:5672/
    ```

4. **Build Docker images for each service:**

    ```sh
    docker build -t gateway-service ./services/gateway
    docker build -t orders-service ./services/orders
    docker build -t kitchen-service ./services/kitchen
    docker build -t stock-service ./services/stock
    ```

5. **Run Docker containers for each service:**

    ```sh
    docker run -d --name gateway-service -p 8080:8080 --env-file .env gateway-service
    docker run -d --name orders-service --env-file .env orders-service
    docker run -d --name kitchen-service --env-file .env kitchen-service
    docker run -d --name stock-service --env-file .env stock-service
    ```


## Microservices
1. **Gateway Service**
Handles external communication and routes requests to the appropriate internal services.

2. **Orders Service**
Handles order creation and management. It includes routes for creating, updating, and retrieving orders.

3. **Kitchen Service**
Manages the preparation of items. It processes orders and updates their status as they are prepared.

4. **Stock Service**
Checks for items in stock from the database and updates stock levels as orders are processed.    