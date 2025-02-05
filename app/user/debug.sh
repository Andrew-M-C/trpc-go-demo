#!/bin/bash
go run -ldflags "-X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=warn" . -conf test/trpc_go.yaml
