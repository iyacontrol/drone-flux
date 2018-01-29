#
# ----- Go Builder Image ------
#
FROM golang:1.8-alpine AS builder

# set working directory
RUN mkdir -p /go/src/drone-flux
WORKDIR /go/src/drone-flux

# copy sources
COPY . .

# build binary
RUN go build -v -o "/drone-flux"

FROM alpine:3.6
MAINTAINER iyacontrol <gaohj2015@yeah.net>

COPY fluxctl /usr/bin/fluxctl
RUN chmod +x /usr/bin/fluxctl 
COPY --from=builder /drone-flux /bin/drone-flux


ENTRYPOINT [ "/bin/drone-flux" ]