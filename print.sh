#!/bin/bash

if [[ -z $1 ]]; then
  to_grade="graderlist.txt"
else
  to_grade=$1
fi

echo ""
column "$to_grade" -t -s "|"
echo ""