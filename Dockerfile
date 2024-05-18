FROM golang:latest
COPY ./ ./
RUN go build
ENTRYPOINT [ "./testYADRO" ]