FROM golang:1.14

WORKDIR /go/src/api.todo
COPY . .
RUN go build ./cmd/todo-api

CMD [ "./todo-api" ]