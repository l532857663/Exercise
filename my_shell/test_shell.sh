#!/bin/bash

# 常量数据 NOTE: root路径需要改为服务所在文件夹
rootDir="/Users/mac/work"
logFile="$rootDir/checkService.log"
cmdDir="cmd/app"
flag="------------------------------------"

# NOTE: 选择调用文件夹
serviceList=(api-service blockchain builder notify wallet-admin wallet-gateway wallet-risk wallet-cron)

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
		return
	fi
	# 编译服务代码
	res=`make clean server`
	writeLog "$res"
	return
}

writeLog() {
	if [ -e $logFile ]
	then
		echo $1 >> $logFile
	else
		touch $logFile
		echo $1 > $logFile
	fi
}

main() {
	echo "Start $rootDir"
	cd $rootDir
	# 新开始时，清空日志
	echo "" > $logFile
	# 循环列表
	for service in ${serviceList[@]}
	do
		# 进入检测路径文件夹
		cd $rootDir
		echo $flag"START"
		# 进入服务代码进行编译
		checkServiceByPath $service
		echo $flag"【$service END】"
		echo ""
	done
	echo "End"
}

# 操作
main
