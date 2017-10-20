#! /usr/bin/python

def globalTest():
	#print outVal
	print outVal2
	print outVal3
	global outVal2
	outVal = 2
	outVal2 = 10
	print outVal
	print outVal2
	print outVal3

outVal = 3
outVal2 = 9
outVal3 = 100
print outVal
print outVal2
print outVal3
globalTest()
print outVal
print outVal2
print outVal3
