# Gaming Server

> John Hammond | August 30th, 2020

-------------------------------

> Can you gain access to this gaming server built by amateurs with no experience of web development and take advantage of the deployment system.

1. What is the user flag?

Username `john` indicated by an HTML comment in the webpage

`dict.list` found in `/uploads`, found by gobuster

`/secret/secretKey` has SSH private key with a password for `john` user.


```
Note: This format may emit false positives, so it will keep trying even after finding a
possible candidate.
Using default input encoding: UTF-8
Loaded 1 password hash (SSH [RSA/DSA/EC/OPENSSH (SSH private keys) 32/64])
Cost 1 (KDF/cipher [0=MD5/AES 1=MD5/3DES 2=Bcrypt/AES]) is 0 for all loaded hashes
Cost 2 (iteration count) is 1 for all loaded hashes
Will run 12 OpenMP threads
Press 'q' or Ctrl-C to abort, almost any other key for status
letmein          (secretKey)
1g 0:00:00:00 DONE (2020-09-01 18:20) 100.0g/s 22200p/s 22200c/s 22200C/s 2003..starwars
Session completed. 
```

password is `letmein`.


user flag:

```
a5c2ff8b9c2e3d4fe9d4ff2f1a5a6e7e
```


2. What is the root flag?

```
2e337b8c9f3aff0c2b3e8d4e6a7c88fc
```