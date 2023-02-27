# fake traffic generator
A Go app to make fake UDP traffic
</br>
usage:

```bash
fTraffic -t 1
```
-t1 send 100Gb </br>
-t2 send 200Gb </br>
-t3 send 300Gb </br>
-t4 send 400Gb </br>
-t5 send 500Gb </br>

</br>

build:
```bash
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main main.go
```
