#!/usr/bin/env bash

[ $# -ne 1 ] && echo "Usage: ./error_handling <greetee>" && exit 1
echo "Hello, $1"
