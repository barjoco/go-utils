#!/bin/sh

for f in ttf/*.ttf; do
    base64 $f > "base64/$(basename $f).base64";
done