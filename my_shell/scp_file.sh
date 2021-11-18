#!/usr/bin/expect

set timeout 30

set local_file [lindex $argv 0]
set wang "wangchengheng"
set host "59.110.51.0"

spawn scp ${local_file} "${wang}@${host}:/home/${wang}"
expect {
	"*yes/no*" {
		send "yes\r"
		exp_continue
	}
	"*password*" {
		send "${wang}\r"
	}
}
expect eof
