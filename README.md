# Golang Template Project

A template for Golang projects with Docker containerization.

## Features

- HTTP server with GET endpoint
- Hot reloading for development
- Docker containerization

## Development

### Prerequisites

- Docker
- Docker Compose
- VS Code or Cursor (with Remote - Containers extension) for container-based development

### Getting Started

Clone this repository:
   ```bash
   git clone https://github.com/AGiantSquid/golang-template.git
   cd golang-template
   ```

#### Option 1: Command Line Development

1. Start the development environment:
   ```bash
   make up
   ```

   This will start the application with hot reloading enabled through Air.

1. Access the application:
   ```
   http://localhost:8080
   ```

#### Option 2: VSCode/Cursor Development Container

1. Open the project in VSCode or Cursor:
   ```bash
   code .  # for VSCode
   cursor .  # for Cursor
   ```

1. When prompted, click "Reopen in Container" or use the command palette (F1) and select "Remote-Containers: Reopen in Container"

1. The container will build and VSCode/Cursor will connect to it. This provides:
   - Full Go language support with gopls
   - Debugging capabilities with delve
   - Code linting with golangci-lint
   - Hot reloading with Air
   - All extensions pre-configured

1. The development container provides an interactive shell by default. You can then run any of the `make` commands (run `make help` for a full list). For example, run `make dev` to start the application with hot reloading. The application will be available at:
   ```
   http://localhost:8080
   ```

   > Note: The container will remain running as long as your IDE is connected or you have an active terminal session.


## Make Commands

Normal development flow commands are exposed via Make.

Run `make help` for a full list.


1. Run tests

   ```make test```

1. Run linter

   ```make lint```

1. Format code

   ```make fmt```

## Production Build

1. Build the production Docker image:
   ```bash
   docker build -t golang-app .
   ```

1. Run the container:
   ```bash
   docker run -p 8080:8080 golang-app
   ```

## Customizing the Template

- Update the module name in `go.mod`
- Modify the `main.go` file to implement your application logic
- Add dependencies using `go get`