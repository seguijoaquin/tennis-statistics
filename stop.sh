#!/bin/bash
docker-compose -f docker-compose-dev.yaml stop
docker-compose -f docker-compose-dev.yaml down
docker rmi $(docker images | grep -ir "<none>" | awk '{print $4}')
