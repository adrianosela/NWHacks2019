FROM alpine:latest

RUN apk add --update bash curl && rm -rf /var/cache/apk/*

ADD nwhacks_api_linux /bin/nwhacks_api_linux

# Uncomment to enable serving HTTPS
#ADD ssl/cert.pem /cert.pem
#ADD ssl/key.pem /key.pem
#EXPOSE 443

# Uncomment to enable serving HTTP
EXPOSE 80

CMD ["/bin/nwhacks_api_linux"]
