# Go Starter Template

This is a REST API built with Golang and the Fiber framework. This repository also includes a payment gateway if your service requires one.

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

`add on: In this project I use Qodana, Static code analysis by Qodana helps development teams to follow agreed quality standards and deliver readable, maintainable and secure code by JetBrains. If you want to test it on your code, just install Qodana and run it with:` 
```shell
qodana scan
```

## Setup

1. Clone this repository:
    ```shell
    git clone https://github.com/LinguaKidAcademy/lingua-be.git
    cd lingua-be
    ```
2. Set the environment variables in a `.env` file:
    ```shell
    cp .env.example .env
    ```
   And set the environment inside it. If you didn't install postgres yet, you can use docker to build the container.
    ```bash
    docker-compose up -d
   ```
3. Install the dependencies:
    ```shell
    go mod download
    ```
4. Run the application:
- At the first, run the app to migrate and seed the database:
    ```bash
    go run cmd/main.go -migrate -seed
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

### Reference:
- [Go-Fiber](https://github.com/gofiber/recipes/tree/master/clean-architecture)
- [Clean Architecture](https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Qodana](https://www.jetbrains.com/qodana/)
- [Midtrans](https://github.com/Midtrans/midtrans-go)
- [JWT](https://github.com/golang-jwt/jwt)
