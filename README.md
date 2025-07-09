# Soundex Implementation

This repository provides a Soundex implementation, primarily written in Go. The project is structured as an API server to manage a collection of books, making use of the Gin web framework.

## Features

- RESTful API for managing books
- Endpoints to create, read, update, and delete book records
- Modular code structure with separate controllers and database connection logic

## Getting Started

### Prerequisites

- Go installed on your system
- Dependencies listed in the project (Gin, custom bookstore packages)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/adasarpan404/soundex-implementation.git
   cd soundex-implementation
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

### Usage

1. Start the API server:
   ```bash
   go run main.go
   ```

2. The server runs on `localhost:8080` and provides the following endpoints:
   - `GET /books` — List all books
   - `GET /books/:id` — Retrieve a specific book by ID
   - `POST /books` — Add a new book
   - `PUT /books/:id` — Update an existing book
   - `DELETE /books/:id` — Delete a book by ID

### API Example

```bash
curl http://localhost:8080/books
```

## Contributing

Feel free to fork the repository and submit pull requests.

## License

This project is currently unlicensed. Please open an issue if you wish to discuss adding a license.

## Author

[Arpan Das](https://github.com/adasarpan404)
