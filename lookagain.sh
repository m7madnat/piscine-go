#! /bin/bash

find . \( -name "*.sh"  \) | sed 's/.sh//g' | sed 's/^.*\///g' | sort -rn 