#! /bin/bash

#! curl https://api.github.com/users/dauren2089 | jq '.name'

#! curl https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json | jq '.[52] | .name'

curl https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json | jq ' .[] | select( .id == 70 ) | .name'