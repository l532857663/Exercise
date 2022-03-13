#!/bin/bash

# 选择调用文件夹
serviceList=(gateway)
# 常量数据
rootDir="/Users/halou/work"
logFile="$rootDir/checkService.log"
cmdDir="cmd"
flag="------------------------------------"

# 编译服务，输出结果
checkServiceByPath() {
	echo "check [$1]"
	# 检测是否存在cmd目录
	cmdPath="$1/$cmdDir"
	if [ -d "$cmdPath" ]
	then
		cd $1
	else
		echo "the service not have cmd dir!"
		return 1
	fi
	# 编译服务代码
	res=`make clean gateway`
	writeLog "$res"
	return 0
}

writeLog() {
	if [ -e $logFile ]
	then
		echo $1 > $logFile
	else
		touch $logFile
		echo $1 > $logFile
	fi
}

main() {
	echo "Start $rootDir"
	cd $rootDir
	# 循环列表
	for service in $serviceList
	do
		# 进入检测路径文件夹
		cd $rootDir
		echo $flag
		# 进入服务代码进行编译
		checkServiceByPath $service
		if [ ! $? ]
		then
			echo "【$service ERROR】"
		else
			echo "【$service OK】"
		fi
		echo $flag
	done
	echo "End"
}

# 操作
main
