#!/bin/sh

x=20
y=30

if [ $x -gt $y ]; then
   echo "value x is greater than value y"
else
   echo "value x is not greater than value y"
fi

if [ $x -lt $y ]; then
   echo "value x is not less than value y"
else
   echo "value x is not less than value y"
fi

if [ $x -ge $y ]; then
   echo "value x is greater or equal than value y"
else
   echo "value x is not greater than value y"
fi

if [ $x -le $y ]; then
   echo "value x is not less or equal than value y"
else
   echo "value x is not less or equal than value y"
fi
