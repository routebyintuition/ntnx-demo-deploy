#!/bin/bash

killall server-one
cd /
nohup /home/ec2-user/app/server-one > /dev/null 2> /dev/null < /dev/null &
