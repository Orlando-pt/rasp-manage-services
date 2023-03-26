#!/bin/bash

sudo systemctl stop rasp-manager.service

sudo chown root:root bin/server
sudo cp bin/server /usr/bin/server

sudo cp rasp-manager.service /lib/systemd/system/
sudo systemctl daemon-reload
sudo systemctl restart rasp-manager.service

sudo cp index.html /var/www/html/
sudo systemctl restart nginx.service

