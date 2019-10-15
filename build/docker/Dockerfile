###################
## BUILDER IMAGE ##
###################

FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /var/build

COPY . .

RUN go build -o /bin/gojis ./cmd/gojis

#################
## FINAL IMAGE ##
#################

FROM scratch

COPY --from=builder /bin/gojis /bin/gojis

VOLUME [ "/src" ]

CMD ["/bin/gojis", "/src"]