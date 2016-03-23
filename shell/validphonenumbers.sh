#!/bin/bash
#author: mo2zie

awk '$0 ~ /^([0-9]{3}\-|\([0-9]{3}\) )[0-9]{3}\-[0-9]{4}$/ {print $0}' file.txt
