# poc_circuit_breaker

[Blog post PT-BR](https://mfbmina.dev/posts/golang-circuit-breaker/)

It defines two apps: A and B.

A is a microservice with two endppoints: 
- /success will always return 200
- /failure will always return 500

B is a microservice that depends on A. It will call A, and depending on the response, will activate the circuit breaker.

## Running

1. `$ go run app_a/main.go`
1. `$ go run app_b/main.go`
