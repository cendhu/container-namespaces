#!/bin/bash
echo "This is the start-up script for the app"
echo $TEST_ENV_VAR >> /persistData/datelog.txt
while true; do sleep 1; date | tee -a /persistData/datelog.txt; done
