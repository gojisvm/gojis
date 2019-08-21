# How to contribute

1. Create an issue where you tell me that you want to contribute (and what, Runtime, Parser or API, or combinations of them)
2. I will add you to the repository and the respective team(s)

Alternatively, you can

1. Fork the repository
2. Make changes
3. Open a Pull Request against `develop`

# Code requirements

* No test cases can be removed and/or changed (if a test fails, but ran green in the past, this is a regression)
  * If you think a test is incorrect, open an issue, stating why the test should be changed/removed
* Readability > Performance (there **are** exceptions, but don't optimize prematurely)
* The project layout is conformant to [this standard](https://github.com/golang-standards/project-layout)

# Getting started

* Checkout the repository
* Run `go test -short ./...` to ensure everything works (provide the short flags, otherwise all conformance tests will be run, which will take a while)
* To test everything, run `go test -v ./...` (includes parser and conformance tests, takes a while)
* To build the application, run `go build -o gojisvm ./cmd` and it will build the executable file `./gojisvm`