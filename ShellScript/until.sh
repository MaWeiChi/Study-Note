#!/bin/bash

counter=0
until [ $counter -gt 10 ]; do
    echo $counter
    counter=`expr $counter + 1`
done
