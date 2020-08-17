# Caesar Cipher Solver
This application has 3 main functions to interact with caesar ciphers:
- Solve: checks the validity of words after each shift attempt (0-25).
- Encrypt: shift every letter forward, in the given string alphabetically, by a certain shift.
- Decrypt: shift every letter backward, in the given string alphabetically, by a certain shift.

## Features
The solving feature consists of breaking the input string by words and calculating the number of words found in a dictionary after each shift, keeping track of shift that has the most matches.

Rate limiting capabilities to protect against spam from clients and even denial of service (DoS) attacks.

## Usage
Since this application is written in [Go](https://golang.org/), make sure that you have it installed before proceeding.

To get started, execute the following command from the project root directory
```
go run main.go
```
and navigate to `localhost:8090` on your browser.