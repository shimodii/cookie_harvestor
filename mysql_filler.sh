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
  access_token=$(cat $i | jq '.["sAccessToken"]' | rev | cut -c2- | rev | cut -c2-)
  front_token=$(cat $i | jq '.["sFrontToken"]' | rev | cut -c2- | rev | cut -c2-)
  refresh_token=$(cat $i | jq '.["sRefreshToken"]' | rev | cut -c2- | rev | cut -c2-)
  # parsed_access=$(echo $access_token | cut -c2- | rev | cut -c2- | rev)
  # parsed_front=$(echo $front_token | cut -c2- | rev | cut -c2- | rev)
  phone_number=$(echo $i | rev | cut -c5- | rev)
  # cookie="sAccessToken=$parsed_access; sFrontToken=$parsed_front"
  # echo "$cookie"
  echo "access token: $access_token"
  echo "front token: $front_token"
  echo "phone number: $phone_number"
  echo "refresh token: $refresh_token"
#  mysql -h "$DB_HOST" -u "$DB_USER" -p"$DB_PASS" "$DB_NAME" <<EOF
#INSERT INTO divar_cookie (phone_number, refresh_token, access_token, front_token, updated_at, is_active)
#VALUES ('$phone_number', '$refresh_token', '$access_token', '$front_token', $(date "+%s"), '1')
#EOF
done

cd $PWD
