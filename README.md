# golang-async-jobs
a simple example project, of a Go api implementation with RabbitMQ. I divided the project into two contexts. Publishers, responsible of publish messages on message broker (in this case i used rabbitMQ) and Consumers, responsible of consume and remove messages from the queue.

# building
 1) Make sure your machine have docker and docker compose.
 2) Run ``` docker compose up -d ```

# Prepare ambient from devs
This will prepare your environment with all the necessary dependencies for development. (RabbitMQ, PostgresSQL)
 1) Run ``` docker compose -f docker-compose-develop.yaml up -d ``` to setup Database and rabbitMQ.
 2) Run ``` go run internal/consumer/consumer.go ``` to run Consumers.
 2) Run ``` go run internal/publisher/publisher.go ``` to run Publishers.

# Load Test
 1) Make sure your machine have docker and docker compose.
 2) Run ``` sh run-load-test.sh ```