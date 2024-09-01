# Companion API Proxy Service

Companion API Proxy Service is a GoLang-based API layer designed to securely route requests from the Companion Chrome extension to external services such as the Google Gemini API and OpenAI API. The service is optimized for performance, security, and scalability, with a strong focus on handling authentication, rate limiting, and secure API key management.

## Directory Structure

```plaintext
.
├── cmd
│   └── server
│       └── main.go        # Entry point for the Go application
├── internal               # Application-specific code
│   ├── auth
│   │   ├── auth.go        # Authentication-related logic (JWT generation, Clerk integration)
│   │   └── middleware.go  # Authentication middleware
│   ├── cache
│   │   └── cache.go       # Caching layer implementation
│   ├── config
│   │   └── config.go      # Configuration management (env variables, etc.)
│   ├── handlers
│   │   ├── gemini.go      # Handlers for Gemini API requests
│   │   ├── openai.go      # Handlers for OpenAI API requests
│   │   └── stripe.go      # Handlers for Stripe payment requests
│   ├── logging
│   │   └── logging.go     # Centralized logging logic
│   ├── middleware
│   │   └── middleware.go  # General middleware (CORS, rate limiting, etc.)
│   ├── models
│   │   └── models.go      # Data models and types
│   ├── monitoring
│   │   └── monitoring.go  # API monitoring and observability logic
│   ├── rate_limiter
│   │   └── rate_limiter.go # Rate limiting implementation
│   ├── router
│   │   └── router.go      # API routes setup
│   └── services
│       ├── gemini_service.go # Service for handling Gemini API requests
│       ├── openai_service.go # Service for handling OpenAI API requests
│       └── stripe_service.go # Service for handling Stripe payments
├── pkg                    # Utility packages (generic, reusable code)
│   ├── errors
│   │   └── errors.go      # Error handling utilities
│   └── utils
│       └── utils.go       # Utility functions (e.g., for request parsing, etc.)
├── api
│   └── swagger
│       └── swagger.yaml   # Swagger API specification
├── build
│   ├── Dockerfile         # Dockerfile for containerization
│   └── Makefile           # Build and deployment automation
├── deployments
│   ├── kubernetes         # Kubernetes deployment files (if using Kubernetes)
│   └── docker-compose.yaml # Docker Compose setup (if applicable)
├── docs
│   └── README.md          # Documentation for the project
└── tests
    ├── integration        # Integration tests
    └── unit               # Unit tests
```

## Project Overview

The Companion API Proxy Service is designed to securely manage API requests to external services, serving as an intermediary between the Companion Chrome extension and Google Gemini and OpenAI. The service ensures that requests are authenticated, rate-limited, and securely processed while providing a caching layer for optimal performance.

## Key Features

-   Authentication and Authorization: Integrates with Firebase and Firestore to authenticate users. Initially the architecture is single tenant, however, it will later be shifted to multi-tenant with support for organizations.
-   Rate Limiting: Implements rate limiting on a per-user basis to prevent abuse and manage resource usage efficiently.
-   Caching Layer: Enhances performance by caching frequently requested data.
    Error Handling and Retry: Provides solid error handling with retry mechanisms for robust operations.
-   Stripe Integration: Manages payments using Stripe, allowing users to sign up for a paid plan and access premium features or go with an initial free trial period.
-   Monitoring and Observability: Integrates with monitoring tools for real-time API observability (APM/metrics).
-   Swagger Documentation: Uses Swagger for API documentation, allowing easy exploration of endpoints.

## Technologies Used

-   GoLang: The core language used for building the API service, chosen for its performance and scalability.
-   Clerk: Handles user authentication and authorization.
-   Stripe: Integrated for payment processing.
-   Docker: Containerizes the application for easy deployment.
-   Kubernetes: Supports deployment and scaling (if applicable).
-   Swagger: Provides API documentation.
-   Redis or Memcached: Caching layer to optimize performance.
-   Prometheus/Grafana: Optional monitoring and observability stack.

## Running the Project

### Prerequisites

-   Go 1.19+
-   Docker (for containerization)
-   Redis/Memcached (for caching)
-   Clerk and Stripe accounts

### Setup Instructions

1. Clone the repository:

```bash
Copy code
git clone https://github.com/yourusername/companion-api-proxy.git
cd companion-api-proxy
```

2. Install dependencies:

```bash
Copy code
go mod tidy
```

3. Set up environment variables:

-   Create a .env file in the root directory.
-   Add your Clerk, Stripe, and external API credentials.

4. Run the server:

```bash
Copy code
go run cmd/server/main.go
```

4. Load the Chrome extension (development):

-   Navigate to chrome://extensions/.
-   Enable Developer mode.
-   Click Load unpacked and select the Chrome extension's build directory.

## Containerization with Docker

1. Build the Docker image:

```bash
Copy code
docker build -t companion-api-proxy .
```

2. Run the Docker container:

```bash
Copy code
docker run -d -p 8080:8080 companion-api-proxy
```
