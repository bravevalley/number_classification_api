# Number Classification API

This project is a simple HTTP API written in Go using the [Gin](https://github.com/gin-gonic/gin) framework. The API classifies a given number by determining its properties, such as whether it is prime, perfect, Armstrong, and its parity (even/odd). It also fetches a fun fact about the number from [Numbers API](http://numbersapi.com/).

## Features

- **Classify Number**: Determine if a number is even or odd, prime, perfect, or Armstrong.
- **Digit Properties**: Computes the digit sum, counts the number of digits, and collects various properties.
- **Fun Fact**: Retrieves a math-related fact about the number from an external API.
- **RESTful API**: Exposes a GET endpoint `/api/classify-number?number=YOUR_NUMBER` for classification.

## Files Overview

- **main.go**  
  Contains the main entry point of the application. It sets up the Gin router and defines the endpoints:
  - `/api/classify-number`: Endpoint for classifying a number.
  - `/`: A default endpoint returning a simple JSON message.

- **API.go**  
  Implements the HTTP handler (`DigiProp`) for the `/api/classify-number` endpoint.  
  It:
  - Retrieves the number from the query parameters.
  - Validates and parses the number.
  - Calls the `SetupNum` function to classify the number.
  - Returns the classification results as JSON.

- **numbers.go**  
  Contains the custom logic for classifying a number, including:
  - The `Digit` struct that holds all the computed properties.
  - Methods to check for prime (`Primer`), perfect (`Perfecter`), Armstrong (`Armstonger`), and parity (`Parityer`).
  - A method (`Facter`) that fetches a fun fact about the number.
  - The `Setup` method that orchestrates the classification.
  - The `SetupNum` function that prepares a `Digit` instance based on an input number.