#!/usr/bin/python

def testParam(arr):
	arr.append("tail")
	arr[0]="0"
	print arr

arr = ["first", "second", "third"]
print arr
testParam(arr)
print arr
