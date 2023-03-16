# to-do-app

# pull the official docker image
FROM golang:1.19

WORKDIR /to-do-app

# install dependencies
COPY ./go.mod .
COPY ./go.sum .

RUN go mod download

COPY . .

RUN go build -o .

CMD [ "./to-do-app" ]