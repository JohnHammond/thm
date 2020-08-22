#!/usr/bin/env python3

import periodictable



elements = [
	"Ag",
	"Hg",
	"Ta",
	"Sb",
	"Po",
	"Pd",
	"Hg",
	"Pt",
	"Lr",
]

message = ''.join([ chr(vars(periodictable)[element].number) for element in elements ])
print(message)
	