#!/usr/bin/env python

import pwn
import binascii

with open("tweedledee-humptydumpty.txt") as h:
	dee = h.readlines()
with open("tweedledum-humptydumpty.txt") as h:
	dum = h.readlines()

for i, x in enumerate(dum):
	e = dum[i].strip()
	o = binascii.unhexlify(e)
	r = o[::-1]
	print(pwn.xor(o,r))
	# print(pwn.xor(o,binascii.unhexlify(dum[-(i+1)].strip())))