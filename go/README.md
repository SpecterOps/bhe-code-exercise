# Solution

The multiples for each entry O(n) marked as `false`. The remaining entries marked as `true` are prime numbers. I created an API endpoint `/prime` to allow the caller to retrieve the prime number based on the position within the array. 


## Usage:

`go run main.go` command stand's up an HTTP Server and listens for client requests on port `8080`. 

Client request sample:
```
curl -H 'Content-Type: application/json' \
      -d '{ "nth_prime":101}' \
      -X POST \
      localhost:8080/prime
```

## Intergration Tests

During the main thread, the integration test simulated a client request using the httptest golang package. Golang spins up a smaller HTTP server and serves the request to the handler. command to run the tests `go test -v -tags=integration`

## TODOs

The unit tests consumes a lot of memory with large numbers. Update the code to reduce time complexity.