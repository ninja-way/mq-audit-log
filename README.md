# mq-audit-log


### To run the server
```
make run
```

#### .env file example
```
DB_HOST=localhost
DB_PORT=5432
DB_USERNAME=login
DB_PASSWORD=password
DB_SSLMODE=disable
DB_DBNAME=audit_logs

AMQP_URI=amqp://guest:guest@localhost:5672/
```