#!/usr/bin/env bash
set -ev
go build
./goto add test "http://example.com"
./goto add test test test "http://baidu.com"
./goto get test test test
./goto get test
./goto list tes