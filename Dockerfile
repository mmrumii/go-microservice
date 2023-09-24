    FROM golang:1.21-alpine as dependencies

    WORKDIR /app
    COPY go.mod go.sum ./

    RUN go mod download

    FROM dependencies AS build
    COPY . ./
    RUN CG0_ENABLE=0 go build -o /main -ldflags="-w -s"

    FROM golang:1.21-alpine 
    COPY --from=build /main /main
    CMD [ "/main" ]