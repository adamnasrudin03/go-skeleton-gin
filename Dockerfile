FROM golang:1.22.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o go-skeleton .

FROM alpine:latest

#Set Timezone 
ENV TZ=Asia/Jakarta
RUN apk add -U tzdata

# set working directory 
WORKDIR /app

COPY --from=builder /app/go-skeleton .

EXPOSE 8000

CMD ["./go-skeleton"]