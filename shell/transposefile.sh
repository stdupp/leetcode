#!/bin/bash
#author: mo2zie

awk '{for(i = 1; i <= NF;i++){if(a[i])a[i]=a[i]" "$i;else a[i]=$i}}END{for(k in a)print a[k]}' file.txt
