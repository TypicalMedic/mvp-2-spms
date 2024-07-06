FROM golang
WORKDIR /app
ENV CGO_ENABLED=0
CMD ["make", "run-commands-to-build-go"]
