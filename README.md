# Go Starter Template

[![Mentioned in Awesome Fiber](https://awesome.re/mentioned-badge-flat.svg)](https://github.com/gofiber/awesome-fiber)

This is a REST API built with Golang and the Fiber framework. This repository also includes payment gateway and storage management if your service requires one.

## API Spec

All api-spec is in the `api/` directory

## Clean Architecture Principles

Clean Architecture is a software design philosophy that emphasizes the separation of concerns, making the codebase more maintainable, testable, and scalable. In this example, the Go Fiber application follows Clean Architecture principles by organizing the code into distinct layers, each with its own responsibility.

### Layers in Clean Architecture

1. **Entities (Core Business Logic)**

- Located in the pkg/entities directory.
- Contains the core business logic and domain models, which are independent of any external frameworks or technologies.

2. **Use Cases (Application Logic)**

- Located in the pkg/ directory (example: pkg/user).
- Contains the application-specific business rules and use cases. This layer orchestrates the flow of data to and from the entities.

3. **Interface Adapters (Adapters and Presenters)**

- Located in the api directory.
- Contains the HTTP handlers, routes, and presenters. This layer is responsible for converting data from the use cases into a format suitable for the web framework (Fiber in this case).

4. **Frameworks and Drivers (External Interfaces)**

- Located in the cmd directory.
- Contains the main application entry point and any external dependencies like the web server setup.

## Requirements

- [Docker](https://www.docker.com/)
- [Go 1.20 or higher](https://go.dev/dl/)

## Setup

1. Clone this repository:
   ```shell
   git clone https://github.com/paundraP/Go-Starter-Template.git
   cd Go-Starter-Template
   ```
2. Set the environment variables in a `.env` file:
   ```shell
   cp .env.example .env
   ```
3. Install the dependencies:
   ```shell
   go mod download
   ```
4. Run the application:

- At the first, migrate and seed the database:
  ```bash
  go run cmd/database/main.go -migrate -seed
  ```
- Then you can run with **air** to automatically reload your application during development whenever you make changes to the source code (dont forget to install air first)

  ```shell
  go install github.com/air-verse/air@latest
  ```

  - If you use mac:
    ```shell
    air -c .air.toml
    ```
  - If you use windows:
    ```shell
    air -c .air.windows.conf
    ```
  - If you use linux:
    ```shell
    air -c .air.windows.conf
    ```

The API should now be running on http://127.0.0.1:3000.

## Contributing

Im excited to have you contribute to this project! If you’d like to help out, feel free to fork the repository, make changes, and submit a pull request. Here's how:

- Fork the repository to create your own copy.
- Clone your fork to your local machine.
- Create a new branch for your changes.
- Make your edits, fix issues, or add new features.
- Test your changes to make sure everything works.
- Commit your changes and push them to your fork.
- Open a pull request with a detailed description of your changes.

Im always happy to collaborate, so don’t hesitate to reach out if you have any questions!


### Reference:

- [Go-Fiber](https://github.com/gofiber/recipes/tree/master/clean-architecture)
- [Clean Architecture](https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Midtrans](https://github.com/Midtrans/midtrans-go)
- [JWT](https://github.com/golang-jwt/jwt)
- [S3 Bucket](https://github.com/aws/aws-sdk-go-v2)
