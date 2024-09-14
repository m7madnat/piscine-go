#!/bin/bash


NAME=$(curl https://adam-jerusalem.nd.edu/assets/superhero/all.json | jq ' .[] | select( .id == 70 ) | .name ')

echo $NAME