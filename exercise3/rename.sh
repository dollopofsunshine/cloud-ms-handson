#!/bin/sh
ls source
for file in source/*.txt
do 
 #echo $file  
 mv $file $file.bak
done
echo "after rename"
ls source
