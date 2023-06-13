#!/bin/bash

ARG=$1

if [ -z "$ARG" ]
then
    echo "No argument specified"
    echo "testall : Test all available APIs"
    exit 0
fi


if [ "$ARG" == "testall" ]
then
    echo "Not yet"
fi