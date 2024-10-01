# Go Payment Processor

This project is a simple payment processing application written in Go. It provides basic functionalities to handle payment transactions accross different payment gateways.

## Features

- Ability to manage multiple payment gateways seamlessly
- Supports multiple data formats (JSON, XML) integration over HTTP
- Flexible data model layer to incorporate all popular databases (SQL Server, Postgres, MySQL)

## Architecture Overview

- The service exposes 2 transaction endpoints to be consumed on the client side (Deposit, Withdrawal) which can be used to make account deposits and withdrawals
- The service also exposes 1 endpoint for gateways to consume for asynchronous transaction updates (reversal)
- The service implements the driving and driven architecture as the PaymentHandler initiates the request to the service layer and the service layer invokes the gateway (driven)
- The gatewayIdentifier service identifies which gateway should be the request initiated based on which the customer is registered.
- The database of choice is Postgres but the DB layer function can be easily extended to support other DBs as the implementation is dictated by interfaces.

![Architecture Overview](architecture.jpg)

## Setup Guide

- Clone the repository `git@github.com:wadeed-deriv/go-payment-processor.git`
- Go to the cloned directory `go-payment-processor`
- Inside the repo do `cd test-server` and run `npm i` (make sure you have `node` and `npm` installed on the system)
- Afterward run `node index.js` this will start the mock gateway server 
- Open a new terminal and go to the project directory 
- If you are on linux you need to open the `docker-compose.yml` file located on the root of `go-payment-processor` directory
- replace the `http://host.docker.internal` to the ip of the system for `GATEWAY_A_URL` and `GATEWAY_A_URL` env variable
- Now in the terminal enter `docker-compose up app -d`
- This will create and run 2 container, Postgres db container which will have the database already initialize in it through script `dbscript\init.sql`
- The other container will be the payment process App container exposing server on `localhost:8080`
- You can start interacting with the Api now. 

# API Documentation


