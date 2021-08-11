# API 

#### Performs tasks:
* Registration & authorization;
* Getting user data;
* Work with universities: 
  * Get all;
  * Get by ID;
  * Get for user;
  * Set for user.
* Work with university directions:
  * Get all per university;
  * Get by ID;
  * Get for user;
  * Get for user with rating;
  * Set for user.

#### Some information about service:
* For authorization using JWT tokens: access and refresh tokens.
* Swagger Open API documentation: ```host:port/api/docs/index.html```
* Makefile for fast using commands: ```./Makefile```

#### Dependencies:
* Gin - Go REST framework;
* sqlx - adapter for database;
* Golangcilint - a set of linters for writing good code in Go;
* [Migrate](https://github.com/golang-migrate/migrate) to up and down migrates on DB;
* Redis - for storing refresh tokens and temporary items such as recovery codes;
* PostgreSQL - as DBMS;
* Prometheus - for getting API metrics;
* Grafana - for visualizing prometheus API metrics.

#### Configuration:
* Environment variables:
```bash
RLMP_CONFIGS_PATH=/usr/src/app/configs # Path to configs folder
RLMP_CONFIG_NAME=config # Config name from configs folder
RLMP_DOTENV_PATH=/usr/src/app/.env # Path to .env file
GIN_MODE=release # Gin mode: release / debug
```
* Environment file (.env):
```bash
DB_PASSWORD=qwerty
```
* Configuration .yaml file: 
```yaml
server:
  port: "8001"
  max_header_bytes: 1048576
  read_timeout: "10s"
  write_timeout: "10s"

db:
  driver: "postgres"
  host: "rlmp-db"
  port: "5433"
  username: "ythosa"
  dbname: "rlmp"
  sslmode: "disable"
  migrations_path: "/usr/src/app/schema"

cache:
  address: "rlmp-cache:6379"
  password: ""
  db: 0

auth:
  access_token:
    ttl: "60m"
    signing_key: "81hJ!*@#Y&12yN#UI!Yjfklsjdf"
  refresh_token:
    ttl: "43200m"
    signing_key: "410fj12fjhsdfjksaj(UY^JIJ98adsuJIKDiHA&*"

parsing:
  rating_list_ttl: "120m"
  read_timeout: "10s"
  read_buffer_size: 6291456
  max_response_body_size: 16777216
  max_conns_per_host: 10
```
