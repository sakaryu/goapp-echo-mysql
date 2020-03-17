#!/bin/sh
set -eu

cd `dirname $0`
cd ..
MYSQL_PWD=password mysql -h mysql -P 3306 -u root sample_dev < db/fixtures.sql
