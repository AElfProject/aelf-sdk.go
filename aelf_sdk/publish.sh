#!/bin/bash



protoc --proto_path=./aelf_sdk --go_out=./aelf_sdk ./aelf_sdk/proto/client.proto
