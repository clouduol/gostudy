#!/usr/bin/python

def param(param1,param2,param3,param4,param5):
	param1 = 100
	param2 = "hehe"
	param3[0] = 1000
	param4["go"]="world"
	#param5[0] = 999
	print param1
	print param2
	print param3
	print param4
	print param5

val1 = 3
val2 = "test"
val3 = [7,8,9]
val4 = dict()
val4["go"]="hello"
val5 = (1,2,3)
print val1
print val2
print val3
print val4
print val5
param(val1, val2, val3, val4, val5)
print val1
print val2
print val3
print val4
print val5

