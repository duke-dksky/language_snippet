#!/bin/bash

curl "http://localhost:9200/" -i --header "Content-Type: application/json" --request POST --data @- << EOF
{
    "Backend": "xxxx",
    "Action": "xxxx",
}
EOF


