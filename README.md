# Simple Go Web Server

## Overview

This Go program sets up a basic web server that serves static files and handles HTTP requests for specific paths. The server runs on port 8080 and includes two main handlers: `/hello` and `/form`. Additionally, it serves static files from the "./static" directory.

## Usage

To run the server, execute the following command:

```bash
go run main.go
```
By default, the server listens on port 8080.

## Code Structure

`main` Function

The main function is the entry point of the program. It initializes the server using the NewServer function and starts the HTTP server using http.ListenAndServe.

`NewServer` Function

The NewServer function returns an instance of http.Handler configured with a ServeMux to handle different routes. It includes:
- A file server to handle static files in the "./static" directory.
- Handlers for the root ("/"), "/form", and "/hello" paths.

`helloHandler` Function

The helloHandler function handles requests to the "/hello" path with the HTTP GET method. If the request doesn't match these criteria, it returns a 404 Not Found error. Otherwise, it responds with a simple "Hello!" message.

`formHandler` Function

The formHandler function handles requests to the "/form" path. It parses the form data from the request and responds with a success message along with the values of "name" and "address" parameters. If there is an error parsing the form, it returns a Bad Request (400) error.

## Dependencies

The program uses the standard net/http package for handling HTTP requests and serving static files.

## Configuration

- The static files are served from the "./static" directory.
- The server runs on port 8080.

## Examples
Hello Endpoint

To access the "Hello!" endpoint:
```bash
curl http://localhost:8080/hello
```

Form Endpoint

To access the "Form" endpoint:
```bash
curl -X POST -d "name=John&address=Doe" http://localhost:8080/form
```

## Error Handling
Errors, if any, are logged using the log.Fatal function, which terminates the program in case of a critical error.

*Feel free to adjust and expand the documentation based on your specific needs or any additional features you might implement in the future.*
