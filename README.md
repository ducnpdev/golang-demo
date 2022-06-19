# golang-demo

## Init project golang:

```bash
go mod init golang-demo
```

## Build Images

```bash
docker build -t golang-demo .
```

## Run container 

```bash
docker run --name   golang-demo -p 8100:8100 -h 0.0.0.0 golang-demo 
```
