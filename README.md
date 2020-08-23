# todo-api

Simple example of api for todo application

## Database configuration

### Connection

1. Copy .env.example file and rename in to the .env
2. Configure the database configuration in the .env file
3. Create the database specified in `POSTGRES_DB`

e.g.

```
POSTGRES_DB=todo

psql# create database todo;
```

### Migrations

There are all migration files in `migrations` folder.

Use [golang-migrate/migrate](https://github.com/golang-migrate/migrate) to run migrations.

Read [full documentation](https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md).

```
$ cd todo-api

$ migrate -database postgres://postgres:admin@localhost:5432/todo -path migrations up
```
