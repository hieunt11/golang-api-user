FROM golang:1.13.8 as builder

VOLUME /api-create-user/src/api-user
ENV GOPATH /api-create-user
WORKDIR "/api-create-user/src/api-user"

COPY . .
ADD go.mod .
#ADD go.sum .

RUN go mod download
RUN go mod verify
RUN go build -o api-create-user

FROM golang:1.13.8
COPY --from=builder "/api-create-user/src/api-user/api-create-user" .
EXPOSE 5000
CMD ["./api-create-user"]




