# Log Server

This is a demo for log server to illustrate how to directly show log and download log.

## Start the Server 

```shell
go run main.go
```

## Test

Open the browse, and directly request the address.

### Show log

URL: `http://localhost:8080/log` or `http://localhost:8080/log?download=false`

Result:

```text
Started by user robin
Building in workspace /Users/robin/.jenkins/workspace/job2
[job2] $ /bin/sh -xe /var/folders/gc/7sjlxfsx4p1f0qlrtx_g3_j80000gn/T/jenkins7178629777156498390.sh
+ echo 'hello world'
hello world
Finished: SUCCESS
```

### Download log

URL: `http://localhost:8080/log?download=true`

Result:

The log file `log.txt` will be downloaded.

### Error Param

URL: `http://localhost:8080/log?download=abc`

Result:

```text
strconv.ParseBool: parsing "abc": invalid syntax
```
