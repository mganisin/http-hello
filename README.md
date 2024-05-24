# Hello World HTTP Server

This is a simple HTTP server written in Go that responds with a configurable
message. The message and port can be set using environment variables.


## Building the Docker Image
1. Build the Docker image:
    ```sh
    docker build -t http-hello .
    ```

## Running the Docker Container
1. Run the Docker container:
    ```sh
    docker run --rm -p 8080:8080 -e HTTP_HELLO_MSG="I will not write another 'Hello, World!' program." http-hello
    ```
