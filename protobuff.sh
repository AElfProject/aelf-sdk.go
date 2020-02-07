#!/bin/bash

SOURCE_FOLDER="./aelf_sdk/protobuf/proto"
GOLANG_COMPILER_PATH="protoc"
IMP_FOLDER="./aelf_sdk/protobuf/generated"
for i in $(ls ${SOURCE_FOLDER}/*.proto)
do 
$GOLANG_COMPILER_PATH --proto_path=${SOURCE_FOLDER} --go_out=$IMP_FOLDER $i
done
