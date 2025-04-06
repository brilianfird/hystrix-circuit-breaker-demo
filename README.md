# Hystrix Circuit Breaker Demo

This repository contains a demo application showcasing the use of the Hystrix circuit breaker pattern in a Go application.

## Table of Contents

- [Introduction](#introduction)
- [Installation](#installation)
- [Usage](#usage)
- [Hystrix Dashboard](#hystrix-dashboard)


## Introduction

The Hystrix Circuit Breaker Demo project demonstrates how to implement the circuit breaker pattern using Hystrix in a Go application. The circuit breaker pattern is used to prevent cascading failures and to improve the resilience of a system by stopping the flow of requests to a failing service.

## Installation

To install and run this demo application, follow these steps:

1. Clone the repository:
   ```sh
   git clone https://github.com/brilianfird/hystrix-circuit-breaker-demo.git
   cd hystrix-circuit-breaker-demo
   ```

2. Install the dependencies:
   ```sh
   go mod tidy
   ```

## Usage

To run the demo application, use the following command:
```sh
go run .\cmd\server\main.go
```

You can then access the application at `http://localhost:8080`. The application provides endpoints to demonstrate the circuit breaker in action.

## Hystrix Dashboard

To visualize the circuit breaker metrics, you can use the Hystrix Dashboard. Follow these steps to set up the dashboard:

1. Run the Hystrix Dashboard using Docker:
   ```sh
   docker run --rm -p 7979:7979 --name hystrix-dashboard steeltoeoss/hystrix-dashboard
   ```

2. Open your browser and navigate to `http://localhost:7979`.

3. Point the dashboard to `http://host.docker.internal:8081` to monitor the circuit breaker metrics.

The dashboard will display real-time metrics and help you understand the state of the circuit breakers in the application.
