# Microservice Series
Creating a set of general microservices (auth, user etc...) for fun and potential use in other projects.
All services will finally be deployed to a kubernetes cluster. This is the first service created in this series.

Links to other microservices:
- https://github.com/JamieBShaw/user-service

## auth-service
Implementation of an auth microservice currently using http, generates access tokens (with refresh, expiry dates etc.).
This connects to the already implemented user-service, link above.
