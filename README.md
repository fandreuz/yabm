# yabm

**Y**et **A**nother **B**oomark **M**anager

## Install

```bash
git clone git@github.com:fandreuz/yabm.git
cd yabm
go install .
```

### Install shell completion

See [here](https://github.com/spf13/cobra/blob/main/site/content/completions/_index.md).

## Database

This CLI is backed by a database. To run a local Postgres DB locally:
```bash
podman run -it --rm \
    -p 5432:5432 \
    -e POSTGRES_DB=admin \
    -e POSTGRES_USER=admin \
    -e POSTGRES_PASSWORD=pwd \
    -v ./dev/db:/docker-entrypoint-initdb.d \
    postgres
```

Image docs [here](https://hub.docker.com/_/postgres).

### Persisting data

The command above will spin up a local DB which won't persist any data upon restart.
You can optionally mount a volume to persist DB data:
```bash
    PGDATA=/var/lib/postgresql/data
    ...
    -e "PGDATA=$PGDATA" \
    -v "./.postgres-data:$PGDATA" \
    ...
```

### Script

A script to run a local DB via Podman with persistence in `./.postgres-data` is 
provided [here](dev/run_db_podman.sh).
