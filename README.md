# File Reader Application

File Reader Application is a service that provides functionality to read files, save data to a database, calculate file content hashes, and return responses based on client requests.

## Table of Contents

1. [Prerequisites](#prerequisites)
2. [Installation](#installation)
3. [Configuration](#configuration)
4. [Usage](#usage)
5. [Testing](#testing)
6. [Deployment](#deployment)
7. [Idea Description](#idea)
8. [Next Steps](#next)

## 1. Prerequisites

Before you begin, ensure you have met the following requirements:

- Docker installed on your machine.
- Basic understanding of Docker, Go, and Nakama.

## 2. Installation

To install and run the File Reader Application, follow these steps:

- Clone this repository to your local machine:

  ```bash
  git clone https://github.com/your-username/file-reader.git
  ```

- Navigate to the project directory:

  ```bash
  cd file-reader
  ```

- Build the Docker image:

  ```bash
  docker-compose build
  ```

## 3. Configuration

- Ensure you have an `.env` file in the root of the project with the following variables:

```dotenv
# Path to JSON files
JSON_FILES_PATH=/jsonFiles
```

Replace `/jsonFiles` with your actual folder path, where the folders with our types are stored. 

## 4. Usage

To start the File Reader Application, run the following command:

```bash
docker-compose up -d
```
### RPC `file-reader`

The `file-reader` RPC feature will be available at [http://127.0.0.1:7351/#/apiexplorer?endpoint=file-reader](http://127.0.0.1:7351/#/apiexplorer?endpoint=file-reader).

Example Request Body:

```json
{
  "type": "type2",
  "version": "1.1.1",
  "hash": "example_hash1"
}
```

The application will be accessible at [http://localhost:7351](http://localhost:7351). You can change the port in the `docker-compose.yml` file if needed.

## 5. Testing

To run unit tests for the application, use the following command:

```bash
go test zepto-lab.com/file-reader/test
```

## 6. Deployment

For deployment, ensure you have the necessary configurations set up, such as Docker and environment variables for production settings.
You can then deploy the application to your preferred hosting platform using Docker.

## 7. Idea Description

I implemented a new RPC function that accepts a request. It attempts to locate a file and retrieve its contents. 
If the data with the same hash is not already in the database, it saves the data to the database. 
For the ID, I use the combination of type and version, as these values are similar to unique identifiers. 
The application then returns a response to the user. 
PostgreSQL was chosen for the database because its JSONB type is well-suited for storing JSON data.

## 8. Next Steps

1. Integration tests: We need to write integration tests to ensure the stability of the application and to accommodate future changes.
2. Load tests: These tests are necessary to evaluate how the application behaves under a large number of requests. We also need to assess the database and file handling performance, especially with very large files.
3. Environment customization: Currently, default values for Nakama have not been changed, and the application is designed for local deployment only. We need to customize it to run on an environment with the appropriate settings.
4. Refactoring: We need advice from an experienced GO programmer on best practices for writing code and tests.

