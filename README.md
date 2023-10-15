# Suranect Api

## Required version
    - Golang (1.18)

## Getting Started
    - Copy file .env.example and rename with .env
    - And go to command prompt "go run main.go"

## URL Docs
- ### / (GET) <br>
    METHOD(GET) <br/>
    RETURN (message)
    <hr/>
- ### /auth/register<br/>
    METHOD(POST) <br/>
    JSON (username, email, password) <br/>
    RETURN (Message, Status)
    <hr/>

- ### /auth/login<br/>
    METHOD(POST) <br/>
    JSON (username, email, password) <br/>
    RETURN (Message, Status, Token)

- ### /auth/send_verify_email <br/>
    METHOD(POST) <br/>
    MIDDLEWARE (AUTH) <br/>
    JSON (username, email, password) <br/>
    RETURN (Message, Status)
- ### /auth/verify_email<br/>
    METHOD(POST) <br/>
    MIDDLEWARE (AUTH) <br/>
    JSON (username, email, password) <br/>
    RETURN (Message, Status)


### Development URL
- ### /migration-up (GET)
