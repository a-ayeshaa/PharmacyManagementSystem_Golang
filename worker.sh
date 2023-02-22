#! /bin/sh
go build -v .
export PHARMACY_CONSUL_URL="127.0.0.1:8500" &&
export PHARMACY_CONSUL_PATH="pharmacy" &&
echo "ENV: PHARMACY_CONSUL_URL=" $PHARMACY_CONSUL_URL
echo "ENV: PHARMACY_CONSUL_PATH=" $PHARMACY_CONSUL_PATH

echo "putting consul config ..."
curl --request PUT --data-binary @config.yml http://localhost:8500/v1/kv/pharmacy

./PharmaProject worker
