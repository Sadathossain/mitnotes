#!/bin/bash
function clean_up {
    sudo docker-compose kill && sudo docker-compose rm -f
}

clean_up
sudo docker-compose up -d

go test -v -tags integration
RET=$?
clean_up

exit ${RET}
