FROM golang:1.14 AS build-env
LABEL maintainer="kroos.chen" \
      build-date={BUILD-DATE} \
      description="A EDA for model input & output process." \
      distribution-scope="private" \
      name={IMAGE} \
      release="0" \
      summary="A EDA for model input & output process." \
      vcs-ref={VCS-REF} \
      vcs-type="git" \
      vendor="kroos.chen" \
      version={VERSION}
EXPOSE 8080
ADD . /src
RUN cd /src && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o service main.go
CMD ["/src/service"]