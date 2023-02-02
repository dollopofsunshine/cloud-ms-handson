#!/bin/sh
for file in `ls source`
do 
  mv source/$file source/$file.bak
done
