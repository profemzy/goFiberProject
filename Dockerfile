FROM golang AS builder
WORKDIR /k8s-fiber
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/k8s-fiber .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN addgroup -S k8s-fiber && adduser -S k8s-fiber -G k8s-fiber
USER k8s-fiber
WORKDIR /home/k8s-fiber
COPY --from=builder /bin/k8s-fiber ./
EXPOSE 3000
ENTRYPOINT ["./k8s-fiber"]