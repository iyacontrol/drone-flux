FROM alpine:3.6
MAINTAINER iyacontrol <gaohj2015@yeah.net>

COPY fluxctl /bin/fluxctl

ENTRYPOINT [ "/bin/fluxctl" ]