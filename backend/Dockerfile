FROM golang:1.14-alpine as build

WORKDIR /app

COPY . /app
RUN go build -o api-server

FROM alpine as runtime 
COPY --from=build /app/api-server /
CMD ./api-server