#!/bin/bash
rm -rf build

mkdir build
cd client
yarn build
cp -rf build ../build/app

cd ../server
go build
cp mini-oa-server ../build/mini-oa-server

cp -rf config ../build/