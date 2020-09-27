# Throwback

> John Hammond | September 5th, 2020

----------------------------

## Hosts:

```
10.200.24.138 - pfSense THROWBACK-FW01
10.200.24.219 - THROWBACK-PROD
10.200.24.232 - THROWBACK-MAIL

10.200.24.176 -THROWBACK-TIME
```


## People Names

```
Rikka Foxx
Hugh Gongo
Jeff Davies
Summers Winters
W Humphrey
Summers Winters
Rikka Foxx
Nana Daiba
Mr Peanutbutter
Jon Peters
J Davies
J Blaire
Hugh Gongo
Frank Murphy
D Jeffers
BoJack Horseman
Frank Murphy
```

## Emails

```
hello@TBHSecurity.com
```

## Flags

1. What is the flag in the guess email within the email server?

```
TBH{ede543c628d365ab772078b0f6880677}
```

2. 


```
TBH{4060a70860f0a1648e5a991de1739888}
```

3. What is the root flag on THROWBACK-FW01?

```
TBH{b6f17a9c06e75ea4a09b79e8d89f9749}
```

4. What is the log flag on THROWBACK-FW01?

```
TBH{c9cf8b688a9b8677a4546781527e4484}
```


5. What is the flag from the poisoned user on THROWBACK-PROD?

```
TBH{277c5929d176569338ce0cff02f328c0}
```

6. What is the second user flag on THROWBACK-PROD?

```
TBH{9b56df4dc5cbda864a246ebfe4964d6c}
```

7. What is the root flag on THROWBACK-PROD?

```
TBH{4d6945c0b80283b875fc7c3a5a057da6}
```

8. What is the user flag on THROWBACK-WS01?

```
TBH{813e2c2709ceb02041891acaec55121d}
```

9. What is the root flag on THROWBACK-WS01?

```
TBH{9c5e361a2368723e042924180be7c958}
```

10. What is the password reset flag on THROWBACK-TIME?

```
TBH{326e71e82d2cfc439ee513340b8d9222}
```

11. What is the root flag on THROWBACK-TIME?

```
TBH{2898c692926188884bf508efe560588f}
```

12. What is the SQL flag on THROWBACK-TIME?

```
TBH{ac3f61048236fd398da9e2289622157e}
```

13. What is the user flag on THROWBACK-DC01?

```
TBH{e6119f456f5107d655be3682559f720f}
```

14. What is the root flag on THROWBACK-DC01?

```
TBH{1b9b614a505017c6fa34cb188581db65}
```

15. What is the account description flag on THROWBACK-DC01?

MercerH -- Hans Mercer has a flag in his description:

```
TBH{b89d9a1648b62a7f2ed01038ac47796b}
```

16. What is the user flag on CORP-DC01?

```
TBH{773e16d57284363e68a4db254860aed1}
```


17. What is the root flag on CORP-DC01?

```
TBH{d2368a76214103ac670a7984b4dba5a3}
```

18. What is the flag on GitHub?

```
TBH{19fa56ead6f82d8c4abc664e2e56f0b1}
```

19. What is the user flag on CORP-ADT01?

```
TBH{250fd11eadbd01e7ed14196611d7b255}
```

20. What is the root flag on CORP-ADT01?

```
TBH{7defa0d5b36c72a48e5966fd2493e19e}
```


21. What is the flag on Twitter?

```
TBH{ca57861454b195f6a5c951a634e05f9e}
```

22. What is the flag on LinkedIn?

```
TBH{2913c22315f3ce3c873a14e4862dd717}
```

23. What is the flag in the source code of Breach GTFO?

```
TBH{53f3a6cb77f633edd9749926b9a9217b}
```

24. What is the flag on the Corporate Mail server?

```
TBH{19b6ca4281bbef3ee060aaf1c2eb4021}
```

25. What is the user flag on TBSEC-DC01?

```
TBH{3efabe3366172f3f97d1123f2cc6dfb5}
```

Cracked password for SQL Service account?

```
mysql337570
```

Password spraying domain users against the domain controller:

```
THROWBACK.local\JeffersD:Throwback2020
```

On the DC:

```
throwback\jeffersd@THROWBACK-DC01 C:\Users\jeffersd\Documents>type 
backup_notice.txt 
As we backup the servers all staff are to use the backup account fo
r replicating the servers
Don't use your domain admin accounts on the backup servers.        

The credentials for the backup are:
TBH_Backup2348!

Best Regards,
Hans Mercer
Throwback Hacks Security System Administrator
```

Using the `backup` account, (which has DCSync rights) we could use `SecretsDump.py` from impacket to dump all the hashes on the domain. 

With that, we found `MercerH` password and we were able to crack with the `colabcat` and `OneRuleToRuleThemAll.rules` rules file.

```
pikapikachu7
```

We could now SSH in and get the `root.txt` as the `MercerH` user had local admin rights

```
TBH{1b9b614a505017c6fa34cb188581db65}
```

Finding credentials on the Github
https://github.com/RikkaFoxx/Throwback-Time


```
define('DB_SRV', 'localhost');
define('DB_PASSWD', "Management2018");
define('DB_USER', 'DaviesJ');
define('DB_NAME', 'timekeepusers');
```

So, CORP-ADT01 is `10.200.24.243` and we can use the above credentials to log in.


```
HRE-KDoiser@TBHSecurity.com
```



Final hash crack:

```
securityadmin284650
```


```
TBH{ec08be8aa9113b47f321b5032a27b220}
```