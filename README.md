# Flotta Webapp Backend

## Before you begin
you need to have:
- Docker installed
- Docker-compose installed

## Getting Started
to see the app up and running, run:
```
make app-up
```
browse the app at: http://localhost:8080

## For Developers
run the following commands to get started:
```
cp .env.example .env
docker-compose up
```

Run the following command to run the tests:
```bash
make test
```

### Some Specifications for the App
this app is built with the following libraries/packages:
- [Gin Web framework](https://github.com/gin-gonic/gin/) for api routing
- [AWS SDK](https://aws.amazon.com/sdk-for-go/) for AWS S3

### Future Works
- Add more tests for the app (s3 package)
- Add more features to the app:
  - store the data in a database after reading from prometheus, related to [Example app: upload to prometheus](https://github.com/project-flotta/flotta-edge-example#future-works)
- Create analytics for the data collected:
  - Identify the most well known hops in the devices networks.
  - Estimate the connection quality for the devices based on the hoping times.