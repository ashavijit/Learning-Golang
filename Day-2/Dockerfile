FROM golang:1.16.5 as development

WORKDIR /go_app

COPY go.mod go.sum ./

RUN go mod download

RUN go build -o go_app

COPY ./go_app /


RUN go install github.com/gorilla/mux
RUN go install github.com/gorilla/handlers
RUN go install github.com/gorilla/context

EXPOSE 8080

CMD reflex -g -l -p 8080 -c /go_app

