#!/bin/bash

swag fmt \
    && swag init \
    -g "./cmd/api/main.go" \
    --quiet \
    --parseDependency \
    --parseInternal \
    --parseDepth 1
