# Ginkgo API Tests for Rancher

![Golang](https://img.shields.io/badge/Golang-1.17-blue.svg)
![Ginkgo](https://img.shields.io/badge/Ginkgo-1.16.4-9cf.svg)

This repository contains Ginkgo-based API tests for Rancher UI.

## Table of Contents

- [Introduction](#introduction)
- [Requirements](#requirements)
- [Getting Started](#getting-started)
- [Contributing](#contributing)
- [License](#license)

## Introduction

The purpose of this project is to perform end-to-end API testing for the Rancher UI using Ginkgo. It includes test cases to cover API authentication and other essential functionalities.

## Requirements

- Go (1.17 or later)
- Ginkgo (1.16.4 or later)
- Gomega (1.16.0 or later)

## Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/ginkgo-api-tests.git```
2. Navigate to the project directory:
   ```
   cd ginkgo-api-tests
   ```
3. Install Ginkgo and Gomega:

   ```
   go install github.com/onsi/ginkgo/ginkgo@latest
   go install github.com/onsi/gomega@latest
   ```

4. Run the tests:
   ```
   go test -v
   ```

## Contributing
Contributions are welcome! If you find any issues or want to add new features, feel free to create a pull request.

## License 
This project is licensed under the MIT License.

