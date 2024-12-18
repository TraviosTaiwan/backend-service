# YueTai Backend

## Description

Backend service for YueTai

## Set up

### Prerequisites

- Go 1.23.3
- Docker

### Create local database

```
make docker-up
```

## Build the application

```
make build
```

## Run migrations

```
make migrate-run
```

## Start the application

```
make run
```

### API Documentation

Access `/swagger/index.html` on your local machine, you should see the swagger documentation

### Health check

Access `/yuetai/health` on your local machine, you should see the json response

## Troubleshooting

#### Cannot find docker-compose command

This is due to the https://collabnix.com/docker-compose-now-shipped-with-docker-by-default/, hence in some cases the docker-compose command won't be available in the system. To fix this, you can do the following:

1. Create an alias in `/usr/local/bin/docker-compose` or any places in the $PATH

```bash
sudo touch /usr/local/bin/docker-compose
sudo echo '#!/bin/bash' > /usr/local/bin/docker-compose
sudo echo 'docker compose $@' >> /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
```

2. Restart your shell
