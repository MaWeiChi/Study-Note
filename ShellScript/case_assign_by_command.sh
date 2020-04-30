#!/bin/bash

a=$(./demo2.sh) 

case $a in
	demo.sh2*) echo 3
		;;
	demo.sh*) echo 2
esac
