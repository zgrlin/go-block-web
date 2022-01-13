## Go-based Ethereum New Block REST API

Go-based Ethereum new Block alert REST Interface

This REST API support Gin framework.

Handler Method Support:

- Post Block
- Get Block
- Alerting (PING)

Field:

- Height  (Block number)
- Minedby (Block Miner)

## Run

```shell
$ go mod init
$ go run main.go
```

## Web Side

```shell
http://localhost:8080/web Web Interface 
```

## REST Side

```shell
GET http://localhost:8080/blockfeed
POST http://localhost:8080/blockfeed
```

## Unit Test
if you want to reproduce new field use a Unit Test

```shell
$ cd platform
$ go test
```

## JS Injection

```shell
await fetch('/blockfeed', {
	method: 'POST',
	headers: {'content-type': 'application/json'},
	body: JSON.stringify({
		number: '1',
		minedby: '2Miners'
	})
})
```
