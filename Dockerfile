FROM golang:1.20
WORKDIR /app
COPY . ./
RUN go mod download
WORKDIR /app/cmd/api
RUN CGO_ENABLED=0 GOOS=linux go build -o /greenlight
EXPOSE 4000
CMD ["/greenlight"]
