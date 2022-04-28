#!/bin/bash

for i in {1..5}
do
    echo '{"user_id":"jakob"}' | fluent-cat debug.test
done