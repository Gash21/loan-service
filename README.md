# Welcome to My Amartha Test

## Introduction

Hello, my name is Galih Arghubi, I'm full stack software engineer.
I have 11 years of experience in software development, and I have worked in various industries, including finance, telecommunications, logistics and etc.

Here's my CV: [Download Here](https://drive.google.com/file/d/1qzRJaBjWEieonfmAEyoj4VavMcDkh9ol/view?usp=sharing)

For your information, this code is created from scratch using Fiber as the framework and SQLite as the database.
I learn many things when creating this code. I'm trying to implement DDD (Domain Driven Design) pattern and SOLID principles as far as I can.

Please check and try the code. Have fun :D

## How To Run

### Preparation

1. Please copy the .env.example to .env
2. Please fill the .env file with the required values

### Manual Run

```bash
    go run ./cmd
```

### Run using VSCode Debugger

1. Please install the Go extension for VSCode
2. Ctrl+Shift+P / Cmd+Shift+P -> Go: Install/Update Tools
3. Open the folder in VSCode
4. Press F5

## Folder Structure Explanation

```bash
.
├── cmd                                     # Here's the entrypoint for the application
│   ├── main.go                             # Entrypoint of the application
│   └── bootstrap.go                        # Bootstrap the application and registering the route
├── internal                                # The main application code
│   ├── application                         # Application Layer: to orchestrate the business logic
│   │   ├── borrower                        # Borrower Related Business Logic
│   │   │   ├── dto                         # Borrower's DTO to be used on the usecase and presentation layer
│   │   │   ├── usecase.go                  # Borrower's Use Case Interface and Initiation
│   │   │   ├── list.go                     # Borrower's Use Case Implementation
│   │   │   └── ...go                       # Borrower's Use Case Implementation
│   │   ├── employee                        # Employee Related Business Logic
│   │   │   ├── dto                         # Employee's DTO to be used on the usecase and presentation layer
│   │   │   ├── usecase.go                  # Employee's Use Case Interface and Initiation
│   │   │   ├── list.go                     # Employee's Use Case Implementation
│   │   │   └── ...go                       # Employee's Use Case Implementation
│   │   ├── investor                        # Investor Related Business Logic
│   │   │   ├── dto                         # Investor's DTO to be used on the usecase and presentation layer
│   │   │   ├── usecase.go                  # Investor's Use Case Interface and Initiation
│   │   │   ├── list.go                     # Investor's Use Case Implementation
│   │   │   └── ...go                       # Investor's Use Case Implementation
│   │   └── loan                            # Loan Related Business Logic
│   │       ├── dto                         # Loan's DTO to be used on the usecase and presentation layer
│   │       ├── usecase.go                  # Loan's Use Case Interface and Initiation
│   │       ├── list.go                     # Loan's Use Case Implementation
│   │       └── ...go                       # Loan's Use Case Implementation
│   ├── domain                              # Domain Layer
│   │   ├── borrower                        # Borrower's Domain
│   │   ├── employee                        # Employee's Domain
│   │   ├── investor                        # Investor's Domain
│   │   ├── loan                            # Loan's Domain
│   │   └── loan_investor                   # Loan Investor's Domain
│   ├── infrastructure                      # Infrastructure Layer
│   │   ├── borrower                        # Borrower's Database Repository
│   │   ├── employee                        # Employee's Database Repository
│   │   ├── investor                        # Investor's Database Repository
│   │   └── loan                            # Loan's Database Repository
│   ├── presentation                        # Presentation Layer
│   │   └── http                            # HTTP Presentation
│   │       └── handler                     # HTTP Handler
│   │           ├── borrower                # Borrower's HTTP Handler
│   │           │   └── handler.go          # Borrower routing management
│   │           ├── employee                # Employee's HTTP Handler
│   │           │   └── handler.go          # Employee routing management
│   │           ├── investor                # Investor's HTTP Handler
│   │           │   └── handler.go          # Investor routing management
│   │           └── loan                    # Loan's HTTP Handler
│   │               └── handler.go          # Loan routing management
│   └── shared                              # Shared Package and Library
│       ├── config                          # Configuration
│       ├── database                        # Database
│       ├── helper                          # Helper
│       ├── logger                          # Logger
│       ├── pdf                             # PDF
│       └── validator                       # Validator
├── pkg                                     # Package
│   └── deps                                # Dependencies
└── uploads                                 # Uploaded file will be saved here
```
