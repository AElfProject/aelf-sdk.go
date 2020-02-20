#!/bin/bash
client_path=client/
dto_path=dto/
model_path=model/
utils_path=utils/
protobuf_path=protobuf/generated/
for i in $client_path $dto_path $model_path $utils_path $protobuf_path;
do
    cd $i && go build && cd ..
    sleep 1
done
cd ../test && go test
