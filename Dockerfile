FROM sonarsource/sonar-scanner-cli:4.6

LABEL version="0.0.1" \
      repository="https://github.com/dewidyabagus/altastore-service2" \
      homepage="https://github.com/dewidyabagus/altastore-service2" \
      maintainer="SonarSource" \
      com.github.actions.name="SonarCloud Scan" \
      com.github.actions.description="Scan your code with SonarCloud to detect bugs, vulnerabilities and code smells in more than 25 programming languages." \
      com.github.actions.icon="check" \
      com.github.actions.color="green"


# Step Peratama
FROM golang:1.16-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download -x

COPY . .

RUN go build -o main

# Step Kedua
FROM alpine:3.14

WORKDIR /app/webservice

COPY --from=builder /app/.env ./.env
COPY --from=builder /app/main .

EXPOSE 8000

CMD ["./main"] 
