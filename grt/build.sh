#!/bin/bash

g++ -std=c++11 -c wgrt.c
if [ "$?" -ne "0" ]; then
    echo "compilation failed"
    exit 1
fi

ar -x libgrt.a
ar rvs libwgrt.a *.o
rm *.o
rm __.SYMDEF