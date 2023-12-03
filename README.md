# Tasks API
API to save your tasks.

## Localhost
Create .env file:
```
cp .env.sample .env
cp .env.db.sample .env.db
```

Run in terminal:
```
docker run -d -p 5432:5432 --name my-postgres --env-file .env.db postgres:15.3-alpine
```

Then run:
```
env $(grep -v '^#' .env | xargs) go run .
```

## Build
Change .env vars and run:
```
go build
```

## Using docker
### Building
```
docker build -t tasks-api .
```

### Running
```
docker run --env-file .env -p 8080:8080 tasks-api
```

## Using docker compose
### Building and running locally
Change the DB_HOST .env file to `DB_HOST=db` and run:

```
docker compose up --build
```