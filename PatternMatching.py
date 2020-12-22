
import numpy as np
NO_OF_CHARS = 128

# write Pattern matrix
def matrix_to_txt(Mat, name):
	w = open(name + '.txt', 'w')

	for i in range(Mat.shape[0]):
		row = [str(x) for x in Mat[i, :]]
		w.write(' '.join(row) + '\n')
	w.close()


def getNextState(pat, M, state, x):
	'''
	calculate the next state
	'''

	# If the character c is same as next character
	# in pattern, then simply increment state
	if state < M and x == ord(pat[state]):
		return state+1

	i=0
	# ns stores the result which is next state

	# ns finally contains the longest prefix
	# which is also suffix in "pat[0..state-1]c"

	# Start from the largest possible value and
	# stop when you find a prefix which is also suffix
	for ns in range(state,0,-1):
		if ord(pat[ns-1]) == x:
			while(i<ns-1):
				if pat[i] != pat[state-ns+1+i]:
					break
				i+=1
			if i == ns-1:
				return ns
	return 0

def computeTF(pat, M):
	'''
	This function builds the TF table which
	represents Finite Automata for a given pattern
	'''
	global NO_OF_CHARS

	TF = [[0 for i in range(NO_OF_CHARS)]\
		for _ in range(M+1)]

	for state in range(M+1):
		for x in range(NO_OF_CHARS):
			z = getNextState(pat, M, state, x)
			TF[state][x] = z

	return TF

def search(pat, txt):
	'''
	Prints all occurrences of pat in txt
	'''
	global NO_OF_CHARS
	M = len(pat)
	N = len(txt)

	TF = computeTF(pat, M)

	TF_copy = TF
	txt_copy = txt
	matrix_to_txt(np.array(TF), 'test/pattern')
	txt = make_txt_as_matrix(txt)
	matrix_to_txt(np.array(txt), 'test/txt')

	###### For testing purpose ######
	state=0
	for i in range(N):
		state = TF_copy[state][ord(txt_copy[i])]
		if state == M:
			print("Pattern found at index: {}". \
			  format(i-M+1))
	####################################



def make_txt_as_matrix(txt):
	global NO_OF_CHARS
	list_num = [ord(letter) for letter in txt]
	# for each char I create a list of NO_OF_CHARS length,
	# where the ith elements of this list is equal to 1
	l_t = []
	for i in range(len(list_num)):
		l_i = [0 if list_num[i] != j else 1 for j in range(NO_OF_CHARS)]
		l_t.append(l_i)
	return l_t


# Driver program to test above function
if __name__ == '__main__':
    txt = "AABAACAADAABAAABAA"
    pat = "AABA"
    search(pat, txt)

# See PyCharm help at https://www.jetbrains.com/help/pycharm/