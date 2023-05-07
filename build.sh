#!/bin/bash

set -e

# Set the working directory to the script's path
my_abs_path=$(readlink -f "$0")
root_dir=$(dirname "$my_abs_path")

mkdir -p "$root_dir/bin"

cd "$root_dir/entry/cmd/entry_server" && go build -o "$root_dir/bin/entry_svc/entry_server" || { echo "entry_server build failed" ; exit 1; }
cd "$root_dir/entry/cmd/entry_server-cli" && go build -o "$root_dir/bin/entry_svc/entry_cli" || { echo "entry_cli build failed" ; exit 1; }
echo "entry service build done"

cd "$root_dir/characters/cmd/characters_server" && go build -o "$root_dir/bin/character_svc/character_server" || { echo "characters_server build failed" ; exit 1; }
cd "$root_dir/characters/cmd/characters_server-cli" && go build -o "$root_dir/bin/character_svc/character_cli" || { echo "characters_cli build failed" ; exit 1; }
echo "character service build done"

cd "$root_dir/inventories/cmd/inventory_server" && go build -o "$root_dir/bin/inventory_svc/inventory_server" || { echo "inventory_server build failed" ; exit 1; }
cd "$root_dir/inventories/cmd/inventory_server-cli" && go build -o "$root_dir/bin/inventory_svc/inventory_cli" || { echo "inventory_cli build failed" ; exit 1; }
echo "inventory service build done"

cd "$root_dir/items/cmd/item_server" && go build -o "$root_dir/bin/item_svc/item_server" || { echo "item_server build failed" ; exit 1; }
cd "$root_dir/items/cmd/item_server-cli" && go build -o "$root_dir/bin/item_svc/item_cli" || { echo "item_cli build failed" ; exit 1; }
echo "item service build done"

cd "$root_dir" || exit 1

echo "All builds succeeded."
