#!/bin/bash
for i in {1..5}; do
    ./main_build
done

cat logtimes.json
cp logtimes.json logtimes_old.json
echo "cleaning old log_times for new start"
rm logtimes.json