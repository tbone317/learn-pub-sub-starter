# Use the official RabbitMQ image with management plugin
# This Dockerfile extends the base RabbitMQ image to enable the STOMP plugin, which allows clients to connect using the STOMP protocol.
# This is an example of how to customize a Docker image for specific use cases, in this case, enabling a plugin for RabbitMQ.
FROM rabbitmq:3.13-management
RUN rabbitmq-plugins enable rabbitmq_stomp