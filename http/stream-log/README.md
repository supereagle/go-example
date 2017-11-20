# Stream Log Server

This is a demo for stream log server to illustrate how to provide stream log produced from server.
The log content can be inputted through console after the server.

## Start the Server

```shell
# go run main.go
```

## Test

1. Run the test script

```shell
# ./curl-websocket.sh
```

2. Input the log consent from the console of the server

## Result

* Result from server

```
# go run main.go
I1119 12:30:09.571329    4749 main.go:29] Start listen on :8080
I1119 12:32:47.566021    4749 main.go:42] Upgrade to websocket
hello
aaaaaaaaadsef
30454356-=0235123
$@#$%@#$&^%&&^*
```

* Result from test script

```
./curl-websocket.sh
HTTP/1.1 101 Switching Protocols
Upgrade: websocket
Connection: Upgrade
Sec-WebSocket-Accept: qGEgH3En71di5rrssAZTmtRTyFk=

�hello
�aaaaaaaaadsef
�30454356-=0235123
�$@#$%@#$&^%&&^*
```