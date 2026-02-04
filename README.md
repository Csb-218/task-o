# Go Fiber To-Do API

A simple RESTful API for managing To-Do items, built using [Go](https://golang.org/) and the [Fiber](https://gofiber.io/) web framework.

## Features

- Create, Read, Update, and Delete (CRUD) operations for To-Do items.
- Fast and low memory footprint using Fiber.
- JSON-based API.

## Prerequisites

- Go installed on your machine.

## Installation

1. Navigate to the project directory.
2. Install the necessary dependencies:

```bash
go mod tidy
```

## Running the Application

Start the server using the following command:

```bash
go run main.go
```

The server will start on port `3000`. You can access it at `http://localhost:3000`.

## API Endpoints

### General

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `GET` | `/` | Health check, returns "Hello, World!" |

### To-Do Items

All To-Do routes are prefixed with `/todo`.

| Method | Endpoint | Description | Request Body Example |
| :--- | :--- | :--- | :--- |
| `GET` | `/todo` | Retrieve all To-Do items | N/A |
| `GET` | `/todo/:id` | Retrieve a specific item by ID | N/A |
| `POST` | `/todo` | Create a new To-Do item | `{"title": "Buy groceries"}` |
| `PUT` | `/todo` | Update an existing item | `{"id": 1, "title": "Buy milk"}` |
| `DELETE` | `/todo/:id` | Delete an item by ID | N/A |


```