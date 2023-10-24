#!/bin/bash

checkPath=$1
checkSymbol=$2
backName="backupA_$2_*"

existStr() {
	res=`echo $1 | grep "${2}"`
	if [ "$res" != "" ]
	then
		echo "error! $1"
		return 1
	fi
	return 0
}

checkSql() {
	echo "checkSql START"
	cd $checkPath && ls $backName
	dataArr=(`cat $backName | tr ',' ' ' | awk '{print $1}'`)
	echo "get ps data length: "${#dataArr[*]}
	i=0
	for data in ${dataArr[*]}
	do
		flag=0
		res=`grep -rn $data ./wallet`
		i=`expr $i + 1`
		if [ $i -lt 5 ]
		then
			reg="$i, 1, 2"
			existStr $res $reg
			flag=$?
		elif [ $i -lt 101 ]
		then
		 	reg="0, 0, 1"
			existStr $res $reg
			flag=$?
		else
			reg="0, 0, 0"
			existStr $res $reg
			flag=$?
		fi
		echo $i
		if [ $flag != 0 ]
		then
			return $flag
		fi
	done
	echo "checkSql END"
	return
}

main() {
	echo "START"
	checkSql
	echo "END"
}

if [ ! -n "$1" ]
then
	echo "Please input params! Eg: ./checkSql.sh bgw1 BTC"
else
	echo "path $checkSymbol"
	main
fi
