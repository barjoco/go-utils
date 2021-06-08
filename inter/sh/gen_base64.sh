#!/bin/sh

for f in ttf/*.ttf; do
    base64 $f | tr -d '\n' > "base64/$(basename $f).base64";
done
