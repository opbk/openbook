#!/bin/bash

# make migrate
# sudo -u postgres psql openbook < ./integration/resources/openbook.sql
cp -r ./integration/resources/upload ./build/frontend/var/lib/openbook/frontend
./build/frontend/usr/lib/openbook/frontend/frontend -config ./build/frontend/etc/openbook/frontend/config.gcfg
