# docker-compose-sample

docker compose service sample

## How to build myimage

Build my container image.

```bash
$ docker compose build
```

Confirm the built image.

```bash
$ docker images | grep myimage
docker-compose-sample-myimage   latest   c093bc6d209d   9 seconds ago  242MB
```

## How to run compose

```bash
$ docker compose up -V
```

## Attach a running container

```bash
$ docker exec -it mycontainer /bin/bash
f3e8733b0bd9:/app#
```
