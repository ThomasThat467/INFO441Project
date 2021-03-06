FROM alpine
COPY gateway /gateway
RUN apk update && apk add --no-cache ca-certificates
EXPOSE 443
ENTRYPOINT ["/gateway"]