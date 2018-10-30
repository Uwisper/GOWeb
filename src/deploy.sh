#! /bin/sh

kill -9 $(pgrep webserver)
cd ~/GOWeb/
git pull https://github.com/Uwisper/GOWeb.git
cd src/webserver/
chmod 777 webserver
./webserver &