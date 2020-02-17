#!/usr/bin/env bash

[ $# -ne 2 ] && echo "Usage: hamming.sh <string1> <string2>" && exit 1
[ ${#1} -ne ${#2} ] && echo "left and right strands must be of equal length" && exit 1
[ ${#1} -ne ${#2} ] && echo "left and right strands must be of equal length" && exit 1

count=0

for (( i=0; i<${#1}; i++ )) {
        [[ ${1:$i:1} == "${2:$i:1}" ]] || ((count++))
    }

echo $count
