#!/bin/sh
LINUX=$1
EXECUTABLE=$2
VERSION=$3

echo "Set GOOS vars";
env GOOS=linux GOARCH=amd64;

echo "Build bins in go";
go build -o bin/$LINUX/$EXECUTABLE -ldflags="-s -w -X main.version=$VERSION";

echo "Copy app directories";

cp -r configs bin/$LINUX/;
cp -r scripts bin/$LINUX/;
cp -r queries bin/$LINUX/;
mkdir bin/$LINUX/data;
cp build/linux_install.sh bin/$LINUX/;
mkdir bin/$LINUX/log;
touch bin/$LINUX/log/dingo.log;
cd bin/ && tar -czvf ./$LINUX.tar.gz ./$LINUX/ && rm -rf ./$LINUX/;

echo "AsturDB Bins Done \n";
echo "*****************************";
echo "Wofgres - Postgres Enterprise";
echo "*****************************";
