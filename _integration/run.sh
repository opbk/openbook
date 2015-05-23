#!/bin/bash

make migrate
sudo -u postgres psql openbook < ./_integration/resources/openbook.sql
cp -r ./_integration/resources/upload ./_build/frontend/var/lib/openbook/frontend
./_build/frontend/usr/lib/openbook/frontend/frontend -config ./_build/frontend/etc/openbook/frontend/config.gcfg
./_build/frontend/usr/lib/openbook/utils/subscription-notifyer -config ./_build/frontend/etc/openbook/frontend/config.gcfg
