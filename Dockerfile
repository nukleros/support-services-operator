FROM ubuntu:latest
COPY ./bin/manager /manager
RUN chmod +x /manager
RUN chown -R 65532:65532 /manager
USER 65532:65532

ENTRYPOINT ["/manager"]
