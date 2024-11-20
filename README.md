# Order Management Services

This project is an Order Management Service built using various technologies to ensure high performance, scalability, and reliability. The service uses HAProxy for load balancing, Go language for the core service implementation, RabbitMQ for message queuing, and gRPC for internal communication between services.

## Features
- **Load Balancing**: HAProxy is used to distribute incoming requests across multiple instances of the service to ensure high availability and reliability.
- **High Performance**: The core services are implemented in Go, known for its efficiency and performance.
- **Message Queuing**: RabbitMQ is used to handle asynchronous communication and task queuing.
- **Internal Communication**: gRPC is used for efficient and robust internal communication between microservices.

## Architecture
- **HAProxy**: Acts as a gateway to balance the load across multiple instances of the order management service.
- **Go Services**: Core business logic and API endpoints are implemented in Go.
- **RabbitMQ**: Used for queuing tasks and handling asynchronous processing.
- **gRPC**: Facilitates communication between different microservices within the system.

## Getting Started

### Prerequisites
- Go 1.22 or higher
- Docker
- RabbitMQ
- HAProxy

### Installation

1. **Clone the repository**:
    ```sh
    git clone https://github.com/balajiss36/OMSv3.git
    cd OMSv3
    ```

2. **Build the Go services**:
    ```sh
    go build -o order-service ./cmd/order-service
    go build -o payment-service ./cmd/payment-service
    ```

3. **Run RabbitMQ**:
    ```sh
    docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management
    ```

4. **Run HAProxy**:
    ```sh
    docker run -d --name haproxy -p 80:80 -v $(pwd)/haproxy.cfg:/usr/local/etc/haproxy/haproxy.cfg:ro haproxy:latest
    ```

5. **Run the services**:
    ```sh
    ./order-service
    ./payment-service
    ```

### Configuration

- **HAProxy Configuration**: The `haproxy.cfg` file should be configured to balance the load across the instances of your services.
- **RabbitMQ Configuration**: Ensure RabbitMQ is configured properly to handle the message queues required by your services.

### Usage

1. **Order Service**: Handles order creation, updates, and retrieval.
2. **Payment Service**: Manages payment processing and status updates.

### Communication

- **gRPC**: Used for internal communication between the order service and payment service. Ensure the gRPC endpoints are correctly configured and accessible.
