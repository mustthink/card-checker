# Card checker

## Description
Simple application to check if a card number is valid or not.

## Installation
1. Clone the repository
2. Build via `task build-rest` or `task build-grpc`
> This commands build docker image for the application

## Running
1. Run the application via `task run`
> This command runs the application in a docker container

## Usage 
The application can be used via rest or grpc. 
- Rest: 
    - Send a POST request to `http://host:port/validate` with a json body containing the card number
    - Example: 
    ```json
    {
      "card_number":"5375414109067073", 
      "expiration_month":2, 
      "expiration_year":2024
  }
    ```
- Grpc:
  - Send a request to the server using the proto files (internal/grpc/proto)


## Configuration
The application can be configured via flags: (for local running)
- `--host`- host on which the application will listen
- `--port` - port on which the application will listen
- `--s` - type of server to run (rest or grpc)
