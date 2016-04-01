#! /bin/bash -ex

pushd `dirname $0`/.. > /dev/null
root=$(pwd -P)
popd > /dev/null

export GOPATH=$root/gopath
mkdir -p $GOPATH

source $root/ci/vars.sh

go get -v github.com/venicegeo/$APP/...

ls -alh
src=$GOPATH/bin
ls -alh $src
cp $src/bf-service $root
#mv $src $root/$APP.$EXT

tar -czf $APP.$EXT bf-service
