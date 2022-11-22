### Migrations

#### Download a tool to handle migrations

[Migrate](https://github.com/golang-migrate/migrate)

#### Create a migration

`migrate create -seq -ext=.sql -dir=./migrations create_some_table`

#### Run migration

`migrate -path=./migrations -database=$API_DB_DSN up`

#### Rollback migration

`migrate -path=./migrations -database=$API_DB_DSN down`
