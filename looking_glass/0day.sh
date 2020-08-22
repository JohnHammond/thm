#!/usr/bin/expect
set ip [lindex $argv 0]
set port [lindex $argv 1]

while {$port < 13456} {

spawn ssh -p $port 0day@$ip -o StrictHostKeyChecking=no
expect "?"
send -- "yes\r"
set port [expr $port+1]
}
expect eof