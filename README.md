# Travels
Travels - an application that allows you to choose any travel.

### Features:
* Clean Architecture
* Gin Framework
* PostgreSQL
* Migrations(Migrate CLI)
* Unit Tests
* Interaction with Cloudinary Service

### Installation

```bash
$ cd server
```

1. Create your own **.env** file from example and put it into server folder.

2. Run Postgres container(if you don't have any):
```bash
$ make createdb
```
3. Connect to this container and create databases(**travels** and **travels_test**).
4. Apply migrations:
```bash
$ migrate -path "migrations" -database=<your-db-url> up
```
