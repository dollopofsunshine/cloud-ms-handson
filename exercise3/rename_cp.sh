#!/bin/sh
ls /bfg/demofiles
for file in /bfg/demofiles/*.txt
do 
 #echo $file  
 mv $file $file.bak
done
echo "after rename"
ls /bfg/demofiles
