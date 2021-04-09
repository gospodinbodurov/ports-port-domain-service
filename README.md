# Port Domain Service

A sample GRPC service implemented in GO.

You can run 

```
go build -o main .
```

And after that

```
./main
```

A GRPC service will be started on port 6666

If you want to change the hostname you can do

```
./main --grpcAddress=localhost:6667
```