#
# Copyright SecureKey Technologies Inc. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

ARG GO_VER
ARG ALPINE_VER

FROM golang:${GO_VER}-alpine${ALPINE_VER} as golang
RUN apk add --no-cache \
	gcc \
	musl-dev \
	git \
	libtool \
	bash \
	make;
ADD . $GOPATH/src/github.com/trustbloc/edge-adapter
WORKDIR $GOPATH/src/github.com/trustbloc/edge-adapter
ENV EXECUTABLES go git

FROM golang as edge-adapter
ARG GO_TAGS
ARG GOPROXY
RUN GO_TAGS=${GO_TAGS} GOPROXY=${GOPROXY} make mock-webhook


FROM alpine:${ALPINE_VER} as base
COPY --from=edge-adapter /go/src/github.com/trustbloc/edge-adapter/build/bin/webhook-server /usr/local/bin
CMD WEBHOOK_PORT=${WEBHOOK_PORT} webhook-server
