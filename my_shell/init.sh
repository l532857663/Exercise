#!/bin/bash

gateway_server="gateway"
project_server="project_svc"

server="./build/bin/innov-%s-local"
server_log="%s.log"
server_pid=0

startServer() {
	server=`printf "${server}" $1`
	server_log=`printf "${server_log}" $1`
	echo "startServer START"
	# ps aux -> USER PID %CPU %MEM VSZ RSS TTY STAT START TIME COMMAND
	# 获取的结果生成数组
	dataArr=(`ps aux | grep ${server} | awk '{print $2"|"$11}'`)
	echo "get ps data length: "${#dataArr[*]}
	for val in ${dataArr[*]}
	do
		tmp=(${val//|/ })
		if [ ${tmp[1]} == ${server} ]
		then
			server_pid=${tmp[0]}
			echo "get old server pid: "${server_pid}
			# 如果服务已启动，杀死
		    `kill -15 ${server_pid}`
		fi
	done
	if [ ${server_pid} != 0 ]
	then
		echo "stop the old server..."
	fi
	# 启动服务
	echo "start"
	`nohup ${server} > ${server_log} 2>&1 &`

	echo "startServer END"
}

if [ ! -n "$1" ]
then
    echo "eg: ./init.sh [gateway|project|ucenter]"
    exit 0
fi

echo "init.sh START:"
case $1 in
"gateway")
	startServer ${gateway_server}
	;;
"project")
	startServer ${project_server}
	;;
esac
