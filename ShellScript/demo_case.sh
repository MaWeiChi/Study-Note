#!/bin/bash

language='Java'

case $language in
    Java*) echo "是 Java！"
            ;;
    Python*) echo "是 Python！"
            ;;
    C*)     echo "是 C！"
            ;;
    *)      echo "沒 match 到！"
esac
