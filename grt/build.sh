#!/bin/bash

g++ -std=c++11 -c wgrt.c
if [ "$?" -ne "0" ]; then
    echo "compilation failed"
    exit 1
fi

os=$(uname)
if [ "$os" = "Darwin" ]; then
    ar rvs libwgrt.a wgrt.o
elif [ "$os" = "Linux" ]; then
    ar rvs libwgrt.so wgrt.o
fi

rm wgrt.o