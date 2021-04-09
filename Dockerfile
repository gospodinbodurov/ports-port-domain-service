FROM golang:alpine
RUN mkdir /app 
ADD . /app/
VOLUME ["/app"]
WORKDIR /app 
RUN go build -o main .
RUN adduser -S -D -H -h /app appuser
USER appuser
CMD ./main --grpcAddress=$GRPC_ADDRESS
EXPOSE 6666