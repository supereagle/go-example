# Graceful Shutdown

This is an example for graceful shutdown of Go application.

## Result

```shell
$ go run main.go
2019/04/20 19:49:46 Starting HTTP server!
^C2019/04/20 19:49:48 Catch signal: interrupt
2019/04/20 19:49:48 HTTP server graceful shutdowns!
```

# References

- [peterhellberg/graceful.go](https://gist.github.com/peterhellberg/38117e546c217960747aacf689af3dc2)
- [Golang: Gracefully stop application](https://medium.com/@kpbird/golang-gracefully-stop-application-23c2390bb212)
- [HTTP.Server.Shutdown](https://golang.org/pkg/net/http/#Server.Shutdown)
