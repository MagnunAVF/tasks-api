# Tasks API
API to save your tasks.

## Localhost
Create .env file:
```
cp .env.sample .env
```

Run in terminal:
```
docker run -d -p 5432:5432 --name my-postgres \
-e POSTGRES_PASSWORD=mysecretpassword \
-e POSTGRES_USER=postgres \
-e POSTGRES_DB=mydatabase \
postgres:15.3-alpine
```

Then run:
```
go run .
```

## Build
Change .env vars and run:
```
go build
```