# pit - Postgres Initialization tool

<div align="center">

<img src="./assets/pit_icon.png" align="center" width="300px" height="300px"/>  

A small devops tool that lets you run quick psql initializations in automated environments.  

</div>

## Usage

```bash
pit init --sqlfile init.sql --conn postgres://USER:PASSWORD@HOST:5432/DB
```

## Local test
```bash
docker run --name some-postgres -p 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword -e POSTGRES_USER=demo01 -e POSTGRES_DB=db01 -d postgres
```

```bash
export DB_NAME=db01
envsubst < init.sql.template > init.sql
pit init --sqlfile init.sql --conn postgres://demo01:mysecretpassword@localhost:5432/db01
```