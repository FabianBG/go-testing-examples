# Do testing with Golang

This repo contains a set of examples of how to test in GO.

## Project structure
```
1_testing           # Basic testing
2_testingTools      # How improve testing with go
3_benchmark         # Benchmarking go functions
4_mocking           # Create mocks to make easier the unitari testing
```

## Start testing

To run all the test:

`go test ./...`

To run a test in a single file:

`go test avalith/testing/2_testingTools -run ^TestMain$`

To run bechmark tests:

`go test -benchmem avalith/testing/3_benchmark -bench ^(BenchmarkMain)$`


To get a coverage form the tests:

`go test ./... -coverprofile cover.out; go tool cover -func cover.out`