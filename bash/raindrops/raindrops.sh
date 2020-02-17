#!/usr/bin/env bash
export n="$1"
( (( n % 105 == 0 )) && echo -n "PlingPlangPlong") ||
    ( (( n % 35 == 0 )) && echo -n "PlangPlong") ||
    ( (( n % 21 == 0 )) && echo -n "PlingPlong") ||
    ( (( n % 15 == 0 )) && echo -n "PlingPlang") ||
    ( (( n % 7 == 0 )) && echo -n "Plong") ||
    ( (( n % 5 == 0 )) && echo -n "Plang") ||
    ( (( n % 3 == 0 )) && echo -n "Pling") ||
    echo -n "$1"
echo ""
