# My Go Starter

This is a starter project for Go.

# Project Structure

```
cmd/ - Main applications for this project.
pkg/ - Library code that's ok to use by external applications.
deploy/ - Configuration files and templates for deployment.
model/ - Data models.
scripts/ - Scripts to perform various build, install, analysis, etc operations.
test/ - unit and integration tests.
repository/ - Data access layer.
```

# Getting Started

## Prerequisites

- Go 1.13 or higher
- Docker
- Helm
- MiniKube

## Running the application

```shell
make run
```

## API Documentation

- GET http://localhost:8080/users
- POST http://localhost:8080/users/mask