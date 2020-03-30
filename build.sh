#! /bin/bash

# Build web UI

cd ~/Documents/go-video/web
go install
cp $GOPATH/bin/web ~/Documents/go-video/bin/web
cp -R ~/Documents/go-video/templates ~/Documents/go-video/bin
cd ..
./bin/web