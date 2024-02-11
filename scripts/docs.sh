#!/bin/bash

swag fmt \
    && swag init \
    -g "./cmd/api/main.go" \
    --parseDependency \
    --parseInternal \
    --parseDepth 1
