#!/usr/bin/expect

set timeout 30

set wang "wangchengheng"
set root "Aladin@aliyun"
set host "59.110.51.0"
if { [llength $argv] < 1} {
    puts "Usage:"
    puts "$argv0 username"
    exit 0
}

set user [lindex $argv 0]

if { $user == "wangch" } {
	set user $wang
	set passwd $wang
} elseif { $user == "root" } {
	set user "root"
	set passwd $root
}

spawn ssh ${user}@${host}
expect {
	"*password:*" {
		send "${passwd}\r"
	}
	"*yes/no*" {
		send "yes\r"
	}
}
interact
