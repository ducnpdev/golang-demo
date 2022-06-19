# golang-demo

## Init project golang:

```bash
go mod init golang-docker-demo
```

## Build Images

```bash
docker build -t golang-docker-demo .
```

## Run container 

```bash
docker run --name   golang-docker-demo -p 8100:8100 -h 0.0.0.0 golang-docker-demo 
```
