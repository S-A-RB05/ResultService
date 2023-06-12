# ResultService

## Api Endpoints

### Currently no exposed Api Endpoints in this service

## Code documentation

`recieve()`

This Go code defines a function receive that establishes a connection to a RabbitMQ server, declares a queue, and consumes messages from that queue. The function failOnError is a helper function used to log errors and terminate the program if an error occurs. Here's a breakdown of the code:

- The failOnError function is defined as a helper function to log errors and panic if an error occurs.
- The receive function is defined, which sets up a connection to RabbitMQ and consumes messages from a queue.
- Inside the receive function, the code establishes a connection to the RabbitMQ server using the provided URL.
- A channel and a queue are opened using the connection.
- The QueueDeclare function is used to declare a queue named "mt5_results".
- The Consume function is called to start consuming messages from the queue. The consumed messages are acknowledged automatically (auto-ack set to true).
- A goroutine is started to handle incoming messages. Each consumed message is assumed to be a string and passed to the insertResult function.
- The main goroutine waits indefinitely (<-forever) to keep the program running. 
