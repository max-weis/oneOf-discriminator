# Job Scheduler API Example with `oneOf` using oapi-codegen

## Overview

This project demonstrates how to use `oapi-codegen` to implement a one-to-many relationship in a Go HTTP API. Specifically, it showcases how to use the `oneOf` OpenAPI schema feature to handle polymorphic request bodies in an HTTP endpoint.

The example implements a Job Scheduler API that accepts jobs of multiple types. Each job has a unique `id` and a `parameters` field that varies based on the job type. The `jobType` field acts as the discriminator for determining the specific job type.

## Setup

1. Install Dependencies

```bash
go mod tidy && go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
```

2. Generate Code from OpenAPI Spec

```bash
go generate .
```

The generated code will include the models (`Job`, `DataProcessingJob`, `NotificationJob`) and server interface.

## Running the Server

1. Start the server:

```bash
go run .
```

The server will start on `localhost:8080`.

2. Send a POST request to create a job:

**Submit a DataProcessingJob:**

```bash
curl -X POST http://localhost:8080/jobs \
-H "Content-Type: application/json" \
-d '{
  "id": "job-123",
  "parameters": {
    "jobType": "dataProcessing",
    "dataset": "dataset-001",
    "algorithm": "algorithm-x"
  }
}'
```

**Submit a NotificationJob:**

```bash
curl -X POST http://localhost:8080/jobs \
-H "Content-Type: application/json" \
-d '{
  "id": "job-456",
  "parameters": {
    "jobType": "notification",
    "recipient": "user@example.com",
    "message": "Your job is complete."
  }
}'
```

**Submit an Invalid Job Type:**

```bash
curl -X POST http://localhost:8080/jobs \
-H "Content-Type: application/json" \
-d '{
  "id": "job-789",
  "parameters": {
    "jobType": "unknownType",
    "data": "invalid"
  }
}'
```

## Running Tests

Run the unit tests for the PostJobs endpoint:

```bash
go test -v
```
