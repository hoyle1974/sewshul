############################
# STEP 1 build executable binary
############################
FROM docker.io/library/golang@sha256:d78cd58c598fa1f0c92046f61fde32d739781e036e3dc7ccf8fdb50129243dd8 as builder
# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
USER root
#RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates && apk add --no-cache upx
# Create appuser
ENV USER=appuser
ENV UID=10001
# See https://stackoverflow.com/a/55757473/12429735RUN 
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"
WORKDIR $GOPATH/src/mypackage/myapp/
COPY sewshul/account /tmp/account
COPY sewshul/list /tmp/list
COPY sewshul/login /tmp/login
RUN chmod u+x /tmp/account
RUN chmod u+x /tmp/list
RUN chmod u+x /tmp/login
############################
# STEP 2 build a small image
############################
FROM ubuntu:18.04
# Import from builder.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
# Copy our static executable
USER appuser:appuser
COPY --from=builder /tmp/account /go/bin/account
COPY --from=builder /tmp/list /go/bin/list
COPY --from=builder /tmp/login /go/bin/login
RUN chmod ugo+x /go/bin/*
# Use an unprivileged user.
# Expose the port
EXPOSE 8080
# Run the binary.
#ENTRYPOINT /go/bin/"$EXE"
ENTRYPOINT "$EXE"
