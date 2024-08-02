## Payment Gateway

### Project Overview

The Payment Gateway project simulates a payment processing system that allows merchants to process payments and retrieve payment details. This solution includes a payment gateway service, a mock bank simulator, and integrated Swagger documentation for API exploration.

### Project Structure

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
├── docs/         # Swagger documentation
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
- **`docs/`**: Contains generated Swagger documentation.
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

3. **Swagger Documentation**:
   - Swagger UI is available at `/swagger/index.html` for interactive API documentation.

### Features

- **Payment Processing**: Simulates payment processing and stores results in a database.
- **Payment Retrieval**: Allows merchants to retrieve and view details of past payments.
- **Validation Middleware**: Ensures that payment requests meet required criteria.
- **UUID Generation**: Generates UUIDs for each payment based on card number, amount, and timestamp.
- **Swagger Documentation**: Provides interactive API documentation for testing and exploring endpoints.

### How to Run Your Solution

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

### Assumptions Made

- **Environment Variables**: The `.env` file is used to configure necessary environment variables. The application assumes this file is properly set up.
- **Database**: SQLite is used for simplicity. In a production environment, a more robust database system like PostgreSQL or MySQL would be preferred.
- **Bank Simulator**: The provided bank simulator is used to mock responses from an acquiring bank. Real-world scenarios may require more complex implementations.

### Areas for Improvement

- **Database**: Switch to a more scalable database solution like PostgreSQL or MySQL for production use.
- **Error Handling**: Improve error handling to provide more detailed and user-friendly error messages.
- **Configuration Management**: Use a more advanced configuration management tool to handle environment-specific settings.
- **Security**: Implement additional security measures such as HTTPS, input sanitization, and rate limiting.
- **Logging**: Enhance logging to capture more detailed information and use centralized logging solutions for better monitoring and analysis.
- **CI/CD Pipeline**: Implement a continuous integration and continuous deployment pipeline using tools like GitHub Actions or Jenkins to automate testing and deployment.

### Technical Debt and Considerations

1. **Code Duplication**: Ensure that there's no duplicated code across services and controllers to improve maintainability.
2. **Hardcoded Values**: Move all hardcoded values to configuration files to enhance flexibility and adaptability to different environments.
3. **Testing**: Expand the test coverage to include more edge cases and integrate with a CI tool for automated testing.
4. **Documentation**: Enhance the documentation to include more detailed explanations of the codebase, especially for new developers.
5. **Scalability**: Evaluate the system's scalability and refactor the code where necessary to handle a larger volume of transactions.
6. **Monitoring and Metrics**: Integrate monitoring and metrics collection to track system performance and identify bottlenecks.

### Cloud Technologies

- **Cloud Database**: Consider using a managed database service like Amazon RDS or Google Cloud SQL for better scalability and reliability.
- **Deployment**: Use containerization with Docker and orchestration with Kubernetes to ensure consistent deployments and scaling.
- **CI/CD**: Implement continuous integration and continuous deployment pipelines using tools like GitHub Actions or Jenkins to automate testing and deployment.
- **Monitoring**: Use monitoring tools like Prometheus and Grafana for real-time system monitoring and alerting.

### Highlights of UUID Implementation

- **UUID Generation**: UUIDs are generated based on card number, amount, and timestamp to ensure uniqueness.
- **Consistency**: The use of UUIDs ensures consistent and unique identification of each payment, making retrieval and tracking easier.
- **Security**: Masking card numbers in stored data enhances security and PCI compliance.

### Database Implementation

- **SQLite**: The project uses SQLite for simplicity and ease of setup. This choice is suitable for development and testing environments. In a production environment, a more robust and scalable database system like PostgreSQL or MySQL should be considered.
- **ORM**: GORM is used as the ORM for database interactions. GORM provides an easy-to-use interface for CRUD operations and integrates well with Go's ecosystem.

### Why Choose Go

- **Performance**: Go provides high performance with its statically compiled binaries, making it an excellent choice for a high-throughput payment gateway.
- **Concurrency**: Go's built-in support for concurrency with goroutines and channels makes it ideal for handling multiple payment transactions simultaneously.
- **Simplicity**: Go's syntax is simple and clean, reducing the complexity of the codebase and making it easier to maintain.
- **Strong Standard Library**: Go's standard library provides robust support for networking, HTTP servers, and cryptography, all of which are essential for a payment gateway.
- **Community and Ecosystem**: Go has a strong and active community, with a wealth of libraries and frameworks that can accelerate development.

### Workflow of Bank Simulator

- **Endpoint**: The bank simulator exposes a POST endpoint `/simulate_bank` that mimics the behavior of an acquiring bank.
- **Request Handling**: When a request is received, the simulator parses the payment details and randomly determines whether the payment is approved or declined.
- **Response Simulation**: The simulator returns a response with a status of either "APPROVED" or "DECLINED", along with a masked card number for security.
- **Logging**: All requests and responses are logged for debugging and analysis purposes.
