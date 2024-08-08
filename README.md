# Payment Gateway

## Project Overview

The Payment Gateway project simulates a payment processing system that allows merchants to process payments and retrieve payment details. This solution includes a payment gateway service and a mock bank simulator.

## Project Structure

![Untitled drawing (2)](https://github.com/user-attachments/assets/61b89bb6-0dd9-481f-a072-b6d62b0fb39d)


```plaintext
payment_gateway/
├── README.md
├── main.go
├── bank_simulator/
│   ├── bank_simulator.go
├── config/
│   ├── config.yaml
├── controllers/
│   └── payment_controller.go
├── middleware/
│   └── validation.go
├── models/
│   └── payment.go
├── repository/
│   └── payment_repository.go
├── router/
│   └── router.go
├── services/
│   ├── payment_service.go
│   └── payment_service_test.go
├── utils/
│   ├── logger.go
│   └── uuid_generator.go
├── .env
├── config.yaml
├── go.mod
├── go.sum
└── payments.db
```

### File Descriptions

- **`main.go`**: Entry point for the application. Initializes the database, service, and router, and starts the server.
- **`bank_simulator/`**: Directory containing files related to the bank simulator.
  - **`bank_simulator.go`**: Simulates the acquiring bank, mocking payment processing responses.
- **`config/config.yaml`**: Configuration file for the application settings.
- **`controllers/payment_controller.go`**: Handles HTTP requests for processing and retrieving payments.
- **`middleware/validation.go`**: Middleware for validating payment requests.
- **`models/payment.go`**: Defines data models for payments and payment requests.
- **`repository/payment_repository.go`**: Interfaces and implementation for payment data storage and retrieval.
- **`router/router.go`**: Sets up the HTTP routes and handlers.
- **`services/payment_service.go`**: Contains business logic for processing payments and interacting with the bank simulator.
- **`services/payment_service_test.go`**: Contains test cases for the payment service.
- **`utils/logger.go`**: Configures and provides a logger using Uber's Zap.
- **`utils/uuid_generator.go`**: Utility for generating UUIDs based on card number and amount.
- **`.env`**: Environment variable configuration file.
- **`config.yaml`**: Configuration settings for the application.

### Project Workflow

1. **Processing Payments**:
   - A merchant sends a POST request to `/process_payment` with payment details.
   - The payment request is validated by middleware.
   - The service layer interacts with the bank simulator to process the payment.
   - The result is saved to the SQLite database and returned to the merchant.

2. **Retrieving Payments**:
   - A merchant sends a GET request to `/retrieve_payment/:id` with the payment ID.
   - The service layer retrieves the payment details from the database.
   - Payment details, with masked card information, are returned to the merchant.

## Workflow of Bank Simulator

- **Endpoint**: The bank simulator exposes a POST endpoint `/simulate_bank` that mimics the behavior of an acquiring bank.
- **Request Handling**: When a request is received, the simulator parses the payment details and randomly determines whether the payment is approved or declined.
- **Response Simulation**: The simulator returns a response with a status of either "APPROVED" or "DECLINED", along with a masked card number for security.
- **Logging**: All requests and responses are logged for debugging and analysis purposes.

## Features

- **Payment Processing**: Simulates payment processing and stores results in a database.
- **Payment Retrieval**: Allows merchants to retrieve and view details of past payments.
- **Validation Middleware**: Ensures that payment requests meet required criteria.
- **UUID Generation**: Generates UUIDs for each payment based on card number, amount, and timestamp.

## How to Run Your Solution

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/leyantong/payment_gateway.git
   cd payment_gateway
   ```

2. **Install Dependencies**:
   ```bash
   go mod tidy
   ```

3. **Create `.env` File**:
   Create a `.env` file in the root directory with the following content:
   ```plaintext
   BANK_SIMULATOR_URL=http://localhost:8081/simulate_bank
   PORT=:8080
   ```

4. **Run the Bank Simulator**:
   Navigate to the `bank_simulator/` directory and start the simulator:
   ```bash
   go run bank_simulator.go
   ```

5. **Run the Payment Gateway**:
   ```bash
   go run main.go
   ```

6. **Access Swagger UI**:
   Open a browser and go to:
   ```
   http://localhost:8080/swagger/index.html
   ```

7. **Run Tests**:
   ```bash
   go test ./...
   ```
   
### Using cURL to Send Requests and Retrieve Data

cURL is a command-line tool for transferring data using various network protocols. It is commonly used for making HTTP requests, which can be useful for testing and interacting with APIs. Below is a guide on how to use cURL to send requests to the payment gateway and retrieve data.

#### Sending a Payment Request

To process a payment, you need to send a POST request with the payment details to the `/process_payment` endpoint.

**Example cURL Command:**

```sh
 curl -X POST http://localhost:8080/process_payment \             
  -H "Content-Type: application/json" \
  -d '{
    "card_number": "4242424242424242",
    "expiry_month": "12",
    "expiry_year": "2024",
    "cvv": "123",
    "amount": 11,
    "currency": "SD"
  }'

```

**Explanation:**
- `-X POST`: Specifies the request method as POST.
- `http://localhost:8080/process_payment`: The URL of the payment processing endpoint.
- `-H "Content-Type: application/json"`: Sets the `Content-Type` header to `application/json`, indicating that the request body contains JSON data.
- `-d '{...}'`: The `-d` flag sends the specified data in the request body. The data is in JSON format and includes the payment details.

#### Retrieving Payment Details

To retrieve the details of a specific payment, you need to send a GET request to the `/retrieve_payment/:id` endpoint, where `:id` is the unique identifier of the payment.

**Example cURL Command:**

```sh
curl -X GET http://localhost:8080/retrieve_payment/your-payment-id
```

**Explanation:**
- `-X GET`: Specifies the request method as GET.
- `http://localhost:8080/retrieve_payment/your-payment-id`: The URL of the payment retrieval endpoint, with `your-payment-id` replaced by the actual payment ID.

#### Example Usage

1. **Processing a Payment:**

   ```sh
 curl -X POST http://localhost:8080/process_payment \             
  -H "Content-Type: application/json" \
  -d '{
    "card_number": "4242424242424242",
    "expiry_month": "12",
    "expiry_year": "2024",
    "cvv": "123",
    "amount": 11,
    "currency": "SD"
  }'

   ```

   **Expected Response:**
   ```json
   {
     "Status": "APPROVED",
     "PaymentID": "80f8ec93-6ff8-5dcc-9e83-1c438a658d86"
   }
   ```

2. **Retrieving Payment Details:**

   ```sh
   curl -X GET http://localhost:8080/retrieve_payment/80f8ec93-6ff8-5dcc-9e83-1c438a658d86
   ```

   **Expected Response:**
   ```json
   {
     "ID": "80f8ec93-6ff8-5dcc-9e83-1c438a658d86",
     "CardNumber": "4242",
     "ExpiryMonth": "12",
     "ExpiryYear": "2024",
     "Amount": 100.00,
     "Currency": "USD",
     "Status": "APPROVED"
   }
   ```

## MVC Structure

### Model

The model represents the data layer of the application. It defines the structure of the data and the relationships between different data elements. In this application, the `models` package contains the `Payment` struct, which represents a payment entity.

### View

In the context of a web API, the view is represented by the responses sent back to the client. The Gin framework handles the serialization of responses to JSON format. The `controllers` package contains functions that handle the HTTP requests and send appropriate responses.

### Controller

The controller acts as an intermediary between the model and the view. It handles incoming HTTP requests, processes them (e.g., by interacting with the database or calling external services), and sends back the HTTP responses. In this application, the `controllers` package contains the logic for processing payments and retrieving payment information.

## Why Gin Framework?

Gin is a high-performance HTTP web framework written in Go. It was chosen for the following reasons:

1. **Performance**: Gin is known for its speed and efficiency, making it suitable for building high-performance web applications.
2. **Simplicity**: Gin's API is straightforward and easy to use, which accelerates development and reduces the learning curve.
3. **Middleware Support**: Gin has built-in support for middleware, which is useful for adding functionalities like logging, authentication, and error handling.
4. **Routing**: Gin provides a powerful and flexible routing mechanism, allowing for the easy definition of routes and route groups.

## Cache Mechanism

To prevent duplicate payments, a cache mechanism is used. The cache stores recent payment requests and checks for duplicates before processing a new payment. The cache is implemented as a map where the key is a combination of payment details (card number, expiry date, amount, etc.) and the value is the timestamp of the payment request.

### Use Cases

- **Preventing Duplicate Payments**: The cache helps prevent duplicate payments by checking if a similar payment has been made within a specified duration (e.g., 1 hour).
- **Performance Improvement**: By using the cache to quickly check for duplicates, the application can reduce the number of redundant database queries and improve performance.

## Locks

Locks are used to ensure thread safety when accessing shared resources like the cache. In this application, a `sync.Mutex` is used to protect the cache from concurrent access issues. The mutex is locked before accessing the cache and unlocked after the access is complete. This ensures that only one goroutine can access the cache at a time, preventing race conditions and ensuring data integrity.

## Highlights of UUID Implementation

- **UUID Generation**: UUIDs are generated based on card number, amount, and timestamp to ensure uniqueness.
- **Consistency**: The use of UUIDs ensures consistent and unique identification of each payment, making retrieval and tracking easier.
- **Security**: Masking card numbers in stored data enhances security and PCI compliance.

## Database Implementation

- **SQLite**: The project uses SQLite for simplicity and ease of setup. This choice is suitable for development and testing environments. In a production environment, a more robust and scalable database system like PostgreSQL or MySQL should be considered.
- **ORM**: GORM is used as the ORM for database interactions. GORM provides an easy-to-use interface for CRUD operations and integrates well with Go's ecosystem.


## Assumptions Made

- **Environment Variables**: The `.env` file is used to configure necessary environment variables. The application assumes this file is properly set up.
- **Database**: SQLite is used for simplicity. In a production environment, a more robust database system like PostgreSQL or MySQL would be preferred.
- **Bank Simulator**: The provided bank simulator is used to mock responses from an acquiring bank. Real-world scenarios may require more complex implementations.

### Areas for Improvement

1. **Database**: Switch to a more scalable database solution like PostgreSQL or MySQL for production use to handle larger datasets and provide better performance and reliability.

2. **Error Handling**: Improve error handling to provide more detailed and user-friendly error messages. This includes catching and logging errors effectively, and returning clear, actionable messages to the end users.

3. **Configuration Management**: Use a more advanced configuration management tool to handle environment-specific settings. Tools like Viper or Consul can provide more flexibility and manage configurations for different deployment environments.

4. **Security**: Implement additional security measures such as:
    - **HTTPS**: Ensure all communications between clients and the server are encrypted using HTTPS.
    - **TLS**: Use TLS for encrypting data in transit to prevent interception and tampering.
    - **Input Sanitization**: Validate and sanitize all user inputs to prevent SQL injection and other common attacks.
    - **Rate Limiting**: Implement rate limiting to prevent abuse and mitigate DDoS attacks.

5. **Logging**: Enhance logging to capture more detailed information and use centralized logging solutions for better monitoring and analysis. Tools like ELK stack (Elasticsearch, Logstash, Kibana) or Splunk can be used for this purpose.

6. **CI/CD Pipeline**: Implement a continuous integration and continuous deployment pipeline using tools like GitHub Actions or Jenkins to automate testing and deployment. This ensures that the code is always in a deployable state and reduces the risk of introducing errors in production.

7. **End-to-End Testing for Load Testing**: Implement end-to-end tests that simulate real-world usage and perform load testing to ensure the system can handle high traffic and large volumes of transactions.

8. **Swagger API Documentation**: Add Swagger for interactive API documentation. This will make it easier for developers to explore and test the API endpoints directly from the documentation.

### Technical Debt and Considerations

1. **Code Duplication**: Ensure that there's no duplicated code across services and controllers to improve maintainability. Refactor common logic into reusable functions or services.

2. **Hardcoded Values**: Move all hardcoded values to configuration files to enhance flexibility and adaptability to different environments. This makes it easier to manage environment-specific settings without changing the codebase.

3. **Testing**: Expand the test coverage to include more edge cases and integrate with a CI tool for automated testing. Consider adding end-to-end tests and load testing to ensure the system can handle high traffic and perform reliably under stress.

4. **Documentation**: Enhance the documentation to include more detailed explanations of the codebase, especially for new developers. Ensure that API endpoints are well-documented, and consider adding Swagger for interactive API documentation in the future.

5. **Scalability**: Evaluate the system's scalability and refactor the code where necessary to handle a larger volume of transactions. This may include optimizing database queries, improving caching mechanisms, and using load balancers to distribute traffic evenly across servers.

6. **Monitoring and Metrics**: Integrate monitoring and metrics collection to track system performance and identify bottlenecks. Tools like Prometheus and Grafana can be used to monitor application health, track key performance indicators, and set up alerts for potential issues.


## Cloud Technologies

- **Cloud Database**: Consider using a managed database service like Amazon RDS or Google Cloud SQL for better scalability and reliability.
- **Deployment**: Use containerization with Docker and orchestration with Kubernetes to ensure consistent deployments and scaling.
- **CI/CD**: Implement continuous integration and continuous deployment pipelines using tools like GitHub Actions or Jenkins to automate testing and deployment.
- **Monitoring**: Use monitoring tools like Prometheus and Grafana for real-time system monitoring and alerting.

---
