FROM alpine
COPY gopath/bin/spinnaker /go/bin/spinnaker
ENTRYPOINT /go/bin/spinnaker
