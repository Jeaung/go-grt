#!/bin/bash

g++ -std=c++11 -c wgrt.c
if [ "$?" -ne "0" ]; then
    echo "compilation failed"
    exit 1
fi

os=$(uname)
if [ "$os" = "Darwin" ]; then
    ar rvs darwin/libwgrt.a wgrt.o
elif [ "$os" = "Linux" ]; then
    ar rvs linux/libwgrt.so wgrt.o linux/libgrt.so
fi

rm wgrt.o
