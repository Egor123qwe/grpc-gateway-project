# Quote Web Service

This web service has basic endpoints for working with user data, such as creating, deleting and getting a user, and also implements a system of subscriptions and unsubscribes. The project is an example of the use of grps-gateway in Go.

## Technologies Used

- **Language**: Go
- **API's**: REST, gRPC
- **Technologies**: grpc-gateway, Docker, Docker Compose, mockery, testify, protobuff, Swagger
- **Storage**: MongoDB

## Additional Notes

- In the project directory there is a client directory, that allows you to run gRPC calls to the server API.
- REST endpoints are available with the help of grpc-gateway, that is factically reverse-proxy, that allows you to send http request to the gRPC server