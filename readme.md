[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com) [![forthebadge](https://forthebadge.com/images/badges/fuck-it-ship-it.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/makes-people-smile.svg)](https://forthebadge.com)

##  :blue_heart: Golang
 
## Requirements
 
  - Golang 
  - Go Module
  - other needed modules
  - Docker

## Project layout structure
We're following this [project-layout](https://github.com/golang-standards/project-layout)

## Setup
Install `docker-compose` as described [here](https://docs.docker.com/compose/install/)

Make sure you duplicate `.env.sample` to `.env` and fill in the missing `app_name` value. The remaining can be used as the default value (which are already configured as the same in `docker-compose.yml` ).

## Development

To start the service, simply run the following command:
```
docker-compose up
```

Once the message `Serving from port 3001` is shown in the console, simply use the service from `localhost:3001` to use the service.

To create a new migration, run:
```
make create_migration name=%NAME OF YOUR MIGRATION%
```

To run all pending migrations, run:
```
make migrate_up_all
```

To reverse all migrations, run:
```
make migrate_down_all
```

To create a new seed file, run
```
make create_seed name=%NAME OF YOUR SEED%
```

To run all seed sql files again
```
make run_seeds
```
