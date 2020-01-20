# !/bin/bash

# MySQLサーバーが起動するまでmain.goを実行せずにループで待機する
# until mysqladmin ping -h mysql --silent; do
#   echo 'waiting for mysqld to be connectable...'
#   sleep 2
# done

sleep 15

echo "app is starting...!"
# exec ./go-app
exec /usr/local/bin/go-app