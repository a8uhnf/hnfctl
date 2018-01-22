#!/usr/bin/env bash

rm -rf /usr/local/bin/hnfctl
go build -x ../hnfctl
mv hnfctl /usr/local/bin/
