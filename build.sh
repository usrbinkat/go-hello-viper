#!/bin/bash
# cobra init --pkg-name github.com/CodeSparta/gohi-go
# cobra add mirror
# cobra add bundle
# go build
# gitup devel

goCmd=$(which go)

rm /bin/gohi 2>/dev/null
rm -rf /root/gohi 2>/dev/null
mkdir -p /tmp/bin

plugins="
	"github.com/docktermj/go-hello-viper/configuration" \
	"github.com/docktermj/go-hello-viper/etc" \
	"github.com/docopt/docopt-go" \
"
for i in ${plugins}; do
  ${goCmd} get -u ${i};
done

${goCmd} build

cp -f ./dev /tmp/bin/gohi 2>/dev/null
ls -lah /tmp/bin

