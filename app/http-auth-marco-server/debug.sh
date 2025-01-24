#!/bin/bash
go run -ldflags "-X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=warn" . \
    -conf ../http-auth-server/conf/trpc_go.yaml
