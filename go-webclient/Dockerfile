FROM golang:1.14.13-alpine3.12 as builder
WORKDIR /module
COPY . .
# `skaffold debug` sets SKAFFOLD_GO_GCFLAGS to disable compiler optimizations
ARG SKAFFOLD_GO_GCFLAGS
RUN cd cmd && go build -gcflags="${SKAFFOLD_GO_GCFLAGS}" -o /app

FROM alpine:3.12
# Define GOTRACEBACK to mark this container as using the Go language runtime
# for `skaffold debug` (https://skaffold.dev/docs/workflows/debug/).
ENV GOTRACEBACK=single
CMD ["./app"]
COPY --from=builder /app .
