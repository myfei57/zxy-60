FROM golang:1.23
ENV GOPROXY=off GOSUMDB=off
WORKDIR /app
COPY . .
RUN go build -mod=vendor -o /app/bagsort ./cmd/bagsort
CMD ["/app/bagsort"]
