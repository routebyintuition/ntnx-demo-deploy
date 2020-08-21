#!/bin/bash

killall server-one
cd /
nohup ./server-one > /dev/null 2> /dev/null < /dev/null &
