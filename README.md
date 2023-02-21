# REST API
REST API with Swagger, based on `gin` framework,
using `postgresql` database.

## prepare database
launch postgres in docker container
```bash
$ docker run --name postgres -p 6500:5432 -e POSTGRES_PASSWORD=adminpswd -d postgres
```

use `psql` CLI to create and fill table:
```bash
$ docker exec -it postgres psql -U postgres
postgres=# \c postgres
You are now connected to database "postgres" as user "postgres".
postgres=# ... execute below sql ...
```

```sql
CREATE TABLE IF NOT EXISTS albums (
    id SERIAL PRIMARY KEY,
    title VARCHAR ( 255 ) UNIQUE NOT NULL,
    artist VARCHAR ( 255 ) NOT NULL,
    price VARCHAR ( 255 ) NOT NULL
);

INSERT INTO albums(title,artist,price) VALUES('Blue Train','John Coltrane', 56.99);
INSERT INTO albums(title,artist,price) VALUES('Jeru','Gerry Mulligan', 17.99);
INSERT INTO albums(title,artist,price) VALUES('Sarah Vaughan and Clifford Brown','Sarah Vaughan', 39.99);

SELECT id,title,artist,price FROM albums;
```


## api
launch restapi service (run with all files related to package 'main')
```bash
go run .
```
open swagger in browser http://localhost:8080/docs/index.html
