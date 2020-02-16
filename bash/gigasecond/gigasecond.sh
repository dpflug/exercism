#!/usr/bin/env bash
date -ud "@$(($(date -ud "$1" '+%s')+1000000000))" '+%FT%T'
