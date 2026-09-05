# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
#
# NOTE: this expects a pre-built "manager" binary in the build context (see
# `make build` / the goreleaser `builds` section) rather than compiling it
# here, since goreleaser's docker pipe builds this image from a context that
# only contains the release binary, not the full repo source.
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY manager .
USER 65532:65532

ENTRYPOINT ["/manager"]
