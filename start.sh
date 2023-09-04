#!/bin/bash

for i in $(ls -d -r */); 
    do echo ${i%%/};
    cd ${i%%/}; docker-compose down -v; docker-compose up -d; cd ../
done