#!/bin/bash
#author: mo2zie

awk '{for(i=1;i<=NF;i++){a[$i]++}}END{for(i=1;i<= asort(a,b);i++)print k,a[k]}' words.txt| sort -r -n -k2
