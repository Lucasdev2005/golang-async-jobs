# golang-async-jobs
a simple example project, of a Go api implementation with RabbitMq. I divided the project into two contexts. Publishers, responsible of publish messages on message broker (in this case i used rabbitMQ) and Consumers, responsible of consume and remove messages from the queue.

# Buildind
 1) Make sure your machine have docker and docker compose.
 2) Run ``` docker compose up -d ```