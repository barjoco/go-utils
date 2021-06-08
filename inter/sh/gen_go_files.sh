#!/bin/sh

for f in base64/*.base64; do
    bn=$(basename $f)
    fontName="${bn%%.*}"
    fontName=$(echo $fontName | sed s,-,,)
    gofile='package inter'
    gofile+='\n\n'
    gofile+="// $fontName is a base64 encoding of the font\n"
    gofile+='var '$fontName' = `'$(cat $f)'`'
    echo -e $gofile >"$fontName.go"
done
