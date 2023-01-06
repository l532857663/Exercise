#!/bin/bash

passwd=12345678
path=/root/geth-linux-amd64-1.10.23-d901d853/keystore
nulink=/root/data/nulink
seed=/root/seed.log
filename=""
address=""
run2="docker run --restart on-failure -d \
--name nulink01 \
-p 9151:9151 \
-v $nulink:/code \
-v $nulink:/home/circleci/.local/share/nulink \
-e NULINK_KEYSTORE_PASSWORD \
-e NULINK_OPERATOR_ETH_PASSWORD \
nulink/nulink nulink ursula run --no-block-until-ready"

getBash() {
	echo "getBash START"
	apt-get update
	# apt-get install expect
	apt-get install docker.io
	# wget https://gethstore.blob.core.windows.net/builds/geth-linux-amd64-1.10.23-d901d853.tar.gz && tar -xvzf geth-linux-amd64-1.10.23-d901d853.tar.gz
	# echo export NULINK_KEYSTORE_PASSWORD=$passwd >> ~/.profile
	# echo export NULINK_OPERATOR_ETH_PASSWORD=$passwd >> ~/.profile
	# source ~/.profile
	echo "getBash END"
}

getUtc() {
	echo "getUTC START"
	expect <<EOF
set timeout 30

cd geth-linux-amd64-1.10.23-d901d853
spawn ./geth account new --keystore ./keystore
expect {
	"Password:*" {send "$passwd\n";exp_continue}
	"Repeat password:*" {send "$passwd\n";exp_continue}
}
EOF
	echo "getUTC END"
}

getAddr() {
	echo "getAddr START"
	filename=`ls $path`
	file="$path/$filename"
	mkdir -p $nulink && cp $file $nulink && sudo chmod -R 777 $nulink
	echo "" > $seed
	addr=`sudo cat $file | python3 -c "import sys, json; print(json.load(sys.stdin)['address'])"`
	address="0x$addr"
	echo "getAddr END"
}

getDocker() {
	echo "getDocker START"
	echo "get filename: $filename"
	echo "get address: $address"
	seedValue=""
	if [ ! -n "$filename" ]
	then
		echo "未获取到秘钥文件和钱包地址"
		exit 0
	fi
	expect<<EOF
set timeout 200
set seedValue ""
log_file $seed

spawn docker run -it --rm \
-p 9151:9151 \
-v $nulink:/code \
-v $nulink:/home/circleci/.local/share/nulink \
-e NULINK_KEYSTORE_PASSWORD \
nulink/nulink nulink ursula init \
--signer keystore:///code/$filename \
--eth-provider https://data-seed-prebsc-1-s3.binance.org:8545 \
--network horus \
--payment-provider https://data-seed-prebsc-1-s3.binance.org:8545 \
--payment-network bsc_testnet \
--operator-address $address \
--max-gas-price 21000
expect {
	"*Is this the public-facing address of Ursula*" {send "y\n";exp_continue}
	"*backed up your seed*" {send "y\n";set seedValue `cat $seed | sed -n '11p'`;exp_continue}
	"Confirm seed words*" {send_user "$seedValue\n"}
}
EOF
	echo "seedValue:"$seedValue
	echo "getDocker END"
}

#getBash
#getUtc
getAddr
getDocker
