# Directions Parser 

#### Performs tasks:
* Gets universities from the database (including links to lists of directions), 
parses these lists and pushes them into the database.

#### Dependencies:
* sqlalchemy - ORM for database;
* beautifulsoup4 - for parsing html;
* urllib3 - for getting lists of directions by URL.

#### Configuration:
* Configuration .yaml file:
```yaml
db:
  host: "rlmp-db"
  port: "5433"
  username: "ythosa"
  password: "qwerty"
  dbname: "rlmp"
  sslmode: "disable"
```
