import argparse
import sys
#sys.setrecursionlimit(1500000)
parser = argparse.ArgumentParser(description='QuickSort.')
parser.add_argument('method', choices=['first', 'last', 'median'])
parser.add_argument('file', type=str, help='input file')
parser.add_argument('--test', help="Test getPivot")

comparisons = 0;
method = None

def readFile(fileName):
  f = open(fileName, 'r')
	result = []
	for line in f:
		result.append(int(line))
	return result

def partition(A, l, r):
	global comparisons
	comparisons += (r-l)
	p = A[l]
	i = l + 1
	for j in range(i, r+1):
		if A[j] < p:
			A[j], A[i] = A[i], A[j]
			i += 1
	A[l], A[i-1] = A[i-1], A[l]
	return i-1

def quickSort(A, l, r):
	if r<l:
		return
	p_id = getPivot(A, l, r)
	A[l], A[p_id] = A[p_id], A[l] #swapping
	p_id = partition(A, l, r)
	quickSort(A, l, p_id-1)
	quickSort(A, p_id+1, r)

def getPivot(A, l, r):
	if method == 'first':
		return l
	elif method == 'last':
		return r
	elif method == 'median':
		m = int((r-l)/2) + l
		med = sorted([A[l], A[r], A[m]])[1]
		if med == A[l]:
			return l
		elif med == A[m]:
			return m
		else:
			return r

def testPivot(A, expect):
	assert getPivot(A, 0, len(A)-1) == expect

def test():
	testPivot([1,2,3,4], 1)
	testPivot([1,2,3,4,5], 2)
	testPivot([4,3,2,1], 1)
	testPivot([1,2,3], 1)
	xx = getPivot([23, 45, 3, 46, 231, 2315, 87, 544, 85, 75], 2, 8)
	print xx
	assert xx == 8, xx


if __name__ == "__main__":
    args = parser.parse_args()
    method = args.method
    if args.test:
    	test()
    	exit()
    inputA = readFile(args.file)
    
    quickSort(inputA, 0, len(inputA)-1)
    #print inputA
    print comparisons
    #print getPivot([1, 2, 8, 5, 7, 4], 0, 5)
