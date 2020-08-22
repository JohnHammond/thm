# Wonderland

> John Hammond | August 19th, 2020

----------------------------------

view-source:http://10.10.101.240/r/a/b/b/i/t/

alice:HowDothTheLittleCrocodileImproveHisShiningTail

Needs a `random.py` in the home directory to privesc

sudo -l to switch to `rabbit` user and find `teaParty` setuid program.

It runs `date` and we can create a fake `date` to run first in the path.

This brings us to `hatter` user.

in the hatter user home directory there is a password.txt

```
WhyIsARavenLikeAWritingDesk?
```

As hatter, we can run `perl` and abuse the setuid capability

```
perl -e 'use POSIX qw(setuid); POSIX::setuid(0); exec "/bin/sh";'
```

1. Obtain the flag in user.txt

```
thm{"Curiouser and curiouser!"}
```


2. Escalate your privileges, what is the flag in root.txt?

```
thm{Twinkle, twinkle, little bat! How I wonder what you’re at!}
```