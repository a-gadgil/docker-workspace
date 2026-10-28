FROM --platform=$BUILDPLATFORM alpine:3.18
ARG TARGETPLATFORM
ARG BUILDPLATFORM
RUN echo "I am running on platform: $TARGETPLATFORM (built on $BUILDPLATFORM)" > /arch.txt
CMD ["cat", "/arch.txt"]
