#! /bin/bash

GITA_ID=$(curl -X 'GET' \
  'https://adam-jerusalem.nd.edu/git/api/v1/users/mnatsheh')
echo $GITA_ID