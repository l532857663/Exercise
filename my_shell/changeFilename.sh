#!/bin/bash

changeFilename(){
	# 原文件路径
	filepath=$1
	# 目标文件路径
	airpath=$2
	# 操作数量 -1 无限
	if [ $3 -gt 0 ]
	then
		breakNum=$3
	else
		breakNum=-1
	fi
	num=0
	count=0
	for file in `ls $filepath`
	do
		# 操作数量
		if [ $num == $breakNum ]
		then
			break
		fi
		# 修改的新名字
		newFilename=`echo $file | sed 's/大/-b/g'`
		# 路径加文件名
		oldPath="$1/$file"
		newPath="$2"
		# 保存个数换文件夹
		if [ $(( $num % 100 )) == 0 ]
		then
			count=`expr $count + 1`
			`mkdir $newPath/$count`
		fi
		newFile="$newPath/$count/$newFilename"
		`mv $oldPath $newFile`
		num=`expr $num + 1`
	done
}


# 修改文件名
oldFile=$1
newFile=$2
num=$3

changeFilename $oldFile $newFile $num
