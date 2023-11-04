# Thanks for considering me in your job applicant interviews. 🤍🙏
![logo](./logo.png)
# Setup
in every microservice you see a `env.example` copy a new file from them with name `.env`
now you can run microservices 

# Benchmark and sender 
go inside sender directory and run `go run main.go -n 10` 
`-n` is number of requests. <br>
the sender will send requests 1 per 420ns <br>
my system results for 100k requests: <br>
```go
Successful requests: 100000  //100k
Failed requests: 0
Benchmark: 1m10.53930694s 
```


# Scalability 
To improve performance we use a queue/worker pattern for sending messages to next service. this pattern has helped me for holding up to 10000 messages in case of disconnecting other services that you requested.<br><br>
To improve scalability we used singleton dependency injection for sources that are limited and I used transient dependency injection for anywhere else. <br>  
To improve scalability/reliability we transfer data in multiple socket connections.  <br><br>
To get the best performance REST API implemented without any framework and implemented using golang http standard library. <br><br>



### Clean Architecture
Architecture of all microservices are clean architecture  


# Extra features
#### implemented unit test for all services
#### benchmark on sender microservice

