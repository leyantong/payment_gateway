# Payment Gateway

## Project Overview

The Payment Gateway project simulates a payment processing system that allows merchants to process payments and retrieve payment details. This solution includes a payment gateway service, a mock bank simulator, and integrated Swagger documentation for API exploration.

## Project Structure

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
│   └── payment_service.go
├── utils/
│   └── logger.go
├── payments.db
├── go.mod
├── go.sum
├── docs/         # Swagger documentation
├── swagger/
│   └── swagger.go
└── tests/
    └── payment_test.go
```

### File Descriptions

- **`main.go`**: Entry point for the application. Initializes the database, service, and router, and starts the server.
- **`bank_simulator/`**: Directory containing files related to the bank simulator.
  - **`bank_simulator.go`**: Simulates the acquiring bank, mocking payment processing responses. 
    - **Main Function**: Sets up an HTTP server that listens for requests and simulates various bank responses.
    - **Endpoints**:
      - **`/simulate_bank`**: A POST endpoint that mimics responses from the acquiring bank based on request parameters. 
    - **Simulation Logic**: Handles different payment scenarios such as success, decline, and error conditions to test the payment gateway’s interactions.
- **`config/config.yaml`**: Configuration file for the application settings.
- **`controllers/payment_controller.go`**: Handles HTTP requests for processing and retrieving payments.
- **`middleware/validation.go`**: Middleware for validating payment requests.
- **`models/payment.go`**: Defines data models for payments and payment requests.
- **`repository/payment_repository.go`**: Interfaces and implementation for payment data storage and retrieval.
- **`router/router.go`**: Sets up the HTTP routes and handlers.
- **`services/payment_service.go`**: Contains business logic for processing payments and interacting with the bank simulator.
- **`utils/logger.go`**: Configures and provides a logger using Uber's Zap.
- **`docs/`**: Contains generated Swagger documentation.
- **`swagger/swagger.go`**: Configuration for serving Swagger UI.
- **`tests/payment_test.go`**: Contains test cases for the payment service.

## Project Workflow

1. **Processing Payments**:
   - A merchant sends a POST request to `/process_payment` with payment details.
   - The payment request is validated by middleware.
   - The service layer interacts with the bank simulator to process the payment.
   - The result is saved to the SQLite database and returned to the merchant.

2. **Retrieving Payments**:
   - A merchant sends a GET request to `/retrieve_payment/:id` with the payment ID.
   - The service layer retrieves the payment details from the database.
   - Payment details, with masked card information, are returned to the merchant.

3. **Swagger Documentation**:
   - Swagger UI is available at `/swagger/index.html` for interactive API documentation.

## Features

- **Payment Processing**: Simulates payment processing and stores results in a database.
- **Payment Retrieval**: Allows merchants to retrieve and view details of past payments.
- **Validation Middleware**: Ensures that payment requests meet required criteria.
- **Swagger Documentation**: Provides interactive API documentation for testing and exploring endpoints.

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

5. **Generate Swagger Documentation** (if changes are made):
   ```bash
   swag init
   ```

6. **Run the Payment Gateway**:
   ```bash
   go run main.go
   ```

7. **Access Swagger UI**:
   Open a browser and go to:
   ```
   http://localhost:8080/swagger/index.html
   ```

8. **Run Tests**:
   ```bash
   go test ./...
   ```

## Assumptions Made

- **Environment Variables**: The `.env` file is used to configure necessary environment variables. The application assumes this file is properly set up.
- **Database**: SQLite is used for simplicity. In a production environment, a more robust database system like PostgreSQL or MySQL would be preferred.
- **Bank Simulator**: The provided bank simulator is used to mock responses from an acquiring bank. Real-world scenarios may require more complex implementations.

## Areas for Improvement

- **Database**: Switch to a more scalable database solution like PostgreSQL or MySQL for production use.
- **Error Handling**: Improve error handling to provide more detailed and user-friendly error messages.
- **Configuration Management**: Use a more advanced configuration management tool to handle environment-specific settings.
- **Security**: Implement additional security measures such as HTTPS, input sanitization, and rate limiting.

## Cloud Technologies

- **Cloud Database**: Consider using a managed database service like Amazon RDS or Google Cloud SQL for better scalability and reliability.
- **Deployment**: Use containerization with Docker and orchestration with Kubernetes to ensure consistent deployments and scaling.
- **CI/CD**: Implement continuous integration and continuous deployment pipelines using tools like GitHub Actions or Jenkins to automate testing and deployment.