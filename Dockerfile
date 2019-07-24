FROM golang:alpine AS build-env
ADD . /src
RUN apk update && apk upgrade && \
apk add --no-cache bash git openssh
RUN cd /src && go build entry.go

FROM alpine
WORKDIR /app
COPY --from=build-env /src/entry /app/
COPY --from=build-env /src/settings/* /app/settings/
ENTRYPOINT ./entry