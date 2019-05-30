FROM golang:alpine AS build-env
ADD . /src
RUN apk update && apk upgrade && \
        apk add --no-cache bash git openssh
RUN cd /src && go build main.go

#Stage to runnable docker container
FROM alpine
WORKDIR /app
COPY --from=build-env src/main /app/
COPY --from=build-env src/config.json /app/config.json

# RUN ping -c 3 google.com

# CMD [ "ping", "google.com" ]
ENTRYPOINT ./main
