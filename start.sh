#!/bin/bash


./build.sh

# Check exit status of build script
if [ $? -eq 0 ]; then
    echo "Start running services..."
    ./bin/entry_svc/entry_server &
    ./bin/character_svc/character_server &
    ./bin/inventory_svc/inventory_server &
    ./bin/item_svc/item_server &
    echo "All service are running"
    wait
else
    echo "Build failed."
fi
