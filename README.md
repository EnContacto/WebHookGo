# WebHookGo
 Application built in Go using the Webhook architecture. 
 This **Go** project implements an HTTP server that exposes a webhook on the `/webhook` endpoint. It also includes an internal client that makes POST requests to the   webhook to demonstrate its functionality.
## Features

- Endpoint Webhook: 
  - Listens on `/webhook` and processes POST requests with a JSON payload.
  - Returns a response message in JSON format as `{“response”: “Hello, Bryan!”}`.
- **Internal Client**:
  - Sends POST requests to the webhook with a custom message.
 
## Prerequisites

Before starting, make sure you have the following installed:

- **Go** (version 1.17 or higher).  
  Download it from [golang.org](https://golang.org/).
  
1. **Clone the Repository**.  
   Clone this repository on your local machine using the command:
  `https://github.com/EnContacto/WebHookGo.git`
   `cd WebHookGo`
2. **Run the server**.
   Run the server
   Starts the server by running the compiled file:
   `go run main.go`
   The server will be listening at http://localhost:8080.

## Project Structure.
   The project has the following basic structure:
   ```bash
📁 project-webhook
 ┣ 📄 main.go # Main server and client code.
 ┗ 📄 README.md # Documentation of the project
Code Functionality
Webhook Server
 
