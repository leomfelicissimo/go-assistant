# Chat Application with LangChainGO and Gin

This is a simple chat application that uses LangChainGO and Gin to create an endpoint for interacting with the Azure OpenAI GPT model. Users can send prompts and receive responses from the model.

## Features

- HTTP POST endpoint `/chats` for submitting prompts.
- Utilizes Azure OpenAI for generating responses.

## Prerequisites

- Go (1.17 or later)
- Azure OpenAI credentials

## Installation

1. Clone the repository:

```bash
git clone <repository-url>
cd <repository-directory>
```

2. Install dependencies:

```bash
go mod tidy
```

3. Create a .env file from the example:

```bash
cp .env.example .env
```

4. Fill in the .env file with your Azure OpenAI credentials:

```bash
OPENAI_BASE_URL=https://myresource.openai.azure.com
OPENAI_VERSION=2023-03-15-preview
OPENAI_API_KEY=<your-api-key>
OPENAI_MODEL=<deployment-name>
PORT=8080
GIN_MODE=debug
```

## Running the Application
To run the application, use the following command:

```bash
go run .
```
The server will start on localhost:8080.

## Usage

### Sending a Prompt

To send a prompt to the /chats endpoint, you can use curl or any HTTP client. Here's an example using curl:

```bash
curl -X POST http://localhost:8080/chats \
-H "Content-Type: application/json" \
-d '{"prompt": "What is the capital of France?"}'
```

### Example Response

A successful response will look like this:
```json
{
  "answer": "The capital of France is Paris."
}
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.