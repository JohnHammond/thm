# Anthem

> John Hammond | August 24th, 2020

-------------------------------------


This task involves you, paying attention to details and finding the 'keys to the castle'.

This room is designed for beginners, however, everyone is welcomed to try it out!

Enjoy the Anthem.

In this room, you don't need to brute force any login page. Just your preferred browser and Remote Desktop.

Please give the box up to 5 minutes to boot and configure.


# Task 1

1. 	Let's run nmap and check what ports are open.


No answer needed.

2. 	What port is for the web server?

```
80
```

3. 	What port is for remote desktop service?

```
3389
```

4. What is a possible password in one of the pages web crawlers check for?

```
UmbracoIsTheBest!
```

5. What CMS is the website using?

```
Umbraco
```

6. What is the domain of the website?


```
anthem.com
```

7. What's the name of the Administrator

Found from researching the poem on one of the posts

```
Solomon Grundy
```

Can we find the email address of the administrator?

Change the initials to match that on the "we are hiring" post

```
SG@anthem.com
```

# Task 2

Found on the first webpage source

```
THM{L0L_WH0_US3S_M3T4}
```

2. What is flag 2?

```
THM{G!T_G00D}
```

3. What is flag 3?

```
THM{L0L_WH0_D15}
```

4. What is flag 4?

Found on one of the blog posts

```
THM{AN0TH3R_M3TA}
```

# Task 3

1. Let's figure out the username and password to log in to the box.(The box is not on a domain)

Remmina log in with `SG` and `UmbracoIsTheBest!`

2. Gain initial access to the machine, what is the contents of user.txt?

```
THM{N00T_NO0T}
```

3. Can we spot the admin password?

Go to C:/ and make sure hidden files are visible. Find `backup` and give yourself access to the `restore.txt` file.


```
ChangeMeBaby1MoreTime
```


4. 	Escalate your privileges to root, what is the contents of root.txt?

```
THM{Y0U_4R3_1337}
```