FROM gcr.io/distroless/static:nonroot
COPY manager /
USER 65532:65532

ENTRYPOINT ["/manager"]
