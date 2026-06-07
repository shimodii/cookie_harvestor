#!/bin/bash

DB_HOST="127.0.0.1"
DB_PORT="3306"
DB_USER="root"
DB_PASS="password"
DB_NAME="ads_harvesting"
DB_TABLE="divar_cookie"
PWD=$(pwd)

OUT_DIR="output/"
cd $OUT_DIR

for i in $(ls)
do
  access_token=$(cat $i | jq '.["sAccessToken"]')
  front_token=$(cat $i | jq '.["sFrontToken"]')
  parsed_access=$(echo $access_token | cut -c2- | rev | cut -c2- | rev)
  parsed_front=$(echo $front_token | cut -c2- | rev | cut -c2- | rev)
  file_name=$(echo $i | rev | cut -c5- | rev)
  cookie="sAccessToken=$parsed_access; sFrontToken=$parsed_front"
  echo "$cookie"
  # echo "access token: $access_token"
  # echo "front token: $front_token"
  # echo "phone_number: $file_name"
#  mysql -h "$DB_HOST" -u "$DB_USER" -p"$DB_PASS" "$DB_NAME" <<EOF
#INSERT INTO divar_cookie (cookie, phone_number)
#VALUES ('$cookie', '$file_name')
#EOF
done

cd $PWD
