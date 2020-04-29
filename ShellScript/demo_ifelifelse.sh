#!/bin/bash

value1=20
value2=30
value3=30

if [ $value1 -gt $value2 ]; then
   echo "value1 is greater than value2"
elif [ $value1 == $value3 ]; then
   echo "value1 is equal to value3"
else
   echo "other result"
fi
