# Microservice Series
Creating a set of general microservices (auth, user etc...) for fun and potential use in other projects.
All services will finally be deployed to a kubernetes cluster. This is the second service created in this series.

Links to other microservices:
- https://github.com/JamieBShaw/user-service

## auth-service
![Alt Text](https://github.com/JamieBShaw/auth-service/blob/master/gif/output.gif)


Implementation of an auth microservice currently using http, generates access tokens (with refresh, expiry dates etc.).
This connects to the already implemented user-service, link above.

Can use both HTTP and GRPC, however, main implementation is to be used with grpc as this is meant for intra-microservice
communication. It just uses a simple redis database to delete and insert users auth tokens.

