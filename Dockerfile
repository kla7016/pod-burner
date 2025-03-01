FROM golang:1.23.6 AS build
COPY ./ /app/
WORKDIR /app/
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/pod-burner

FROM alpine:3.12
COPY --from=build /bin/pod-burner /
CMD ["/pod-burner"]