#!/usr/bin/env bash
protoc --proto_path=./proto  --csharp_out=client/Assets/Scripts/Net/proto ./proto/*.proto
protoc --proto_path=./proto --go_out=server/msg ./proto/*.proto

python3 msgIdTool.py