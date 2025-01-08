# Go Currency API

This Go-based currency/banking service provides a simple REST API to retrieve user account balances. It demonstrates an example structure using [chi](https://github.com/go-chi/chi) for routing, basic authentication middleware, and a mock database.

## Table of Contents

- [Overview](#overview)  
- [Features](#features)  
- [Prerequisites](#prerequisites)  
- [Installation](#installation)  

---

## Overview

- **Language**: Go (Golang)  
- **Framework**: [chi](https://github.com/go-chi/chi)  
- **Purpose**: Demonstrates a basic API structure with user authentication, mock database interactions, and error handling.

This service starts an HTTP server on `localhost:8000` and exposes endpoints that allow you to query an account’s balance using a username and token-based authorization.

---

## Features

1. **Authentication Middleware** – Simple token-based auth to verify user credentials.  
2. **Account Balance Retrieval** – Fetch a user’s current balance from a mock database.  
3. **Scalable Structure** – Demonstrates a modular approach with separation of concerns:
   - `api` package for shared request/response types and error handling
   - `cmd/api` for the main application entry point
   - `internal` package for all internal logic (handlers, middleware, tools)

---

## Prerequisites

- Go 1.18+ (or a version that supports Go modules)
- Internet access for fetching Go dependencies

---

## Installation

1. **Clone the repository**:
   git clone https://github.com/Drake-Programming/go-currency-api.git
   cd go-currency-api
2. **Install Dependencies**:
   go mod download
4. **Build**:
   go build ./cmd/api
