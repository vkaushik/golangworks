package backSpaceStringCompare

func backspaceCompare(S string, T string) bool {
	sLastHashIndex := nextHashIndex(S, len(S)-1)
	tLastHashIndex := nextHashIndex(T, len(T)-1)
	sLastCharIndex := nextCharIndex(S, len(S)-1)
	tLastCharIndex := nextCharIndex(T, len(T)-1)

	for true {
		if (sLastCharIndex == -1) && (tLastCharIndex == -1) {
			// both strings are ended
			return true
		} else {
			// both string are not ended
			// ignore character comparisons until a character is found which won't be deleted by backspace
			for sLastHashIndex > sLastCharIndex && sLastCharIndex >= 0 {
				sLastHashIndex = nextHashIndex(S, sLastHashIndex-1)
				sLastCharIndex = nextCharIndex(S, sLastCharIndex-1)
			}
			for tLastHashIndex > tLastCharIndex && tLastCharIndex >= 0 {
				tLastHashIndex = nextHashIndex(T, tLastHashIndex-1)
				tLastCharIndex = nextCharIndex(T, tLastCharIndex-1)
			}
			if (sLastCharIndex == -1) && (tLastCharIndex == -1) {
				// both strings are ended
				return true
			} else if (sLastCharIndex == -1) || (tLastCharIndex == -1) {
				// one string is ended but another is not
				return false
			} else {
				// both strings are not ended. Both characters have to be equal at this point
				if S[sLastCharIndex] != T[tLastCharIndex] {
					return false
				} else {
					sLastCharIndex = nextCharIndex(S, sLastCharIndex-1)
					tLastCharIndex = nextCharIndex(T, tLastCharIndex-1)
				}
			}
		}
	}

	return false
}

func nextHashIndex(str string, from int) (index int) {
	for index = from; index >= 0; index-- {
		if string(str[index]) == "#" {
			break
		}
	}

	return
}

func nextCharIndex(str string, from int) (index int) {
	for index = from; index >= 0; index-- {
		if string(str[index]) != "#" {
			break
		}
	}

	return
}


Solution {
    public boolean backspaceCompare(String S, String T) {
        int i = S.length() - 1, j = T.length() - 1;
        int skipS = 0, skipT = 0;

        while (i >= 0 || j >= 0) { // While there may be chars in build(S) or build (T)
            while (i >= 0) { // Find position of next possible char in build(S)
                if (S.charAt(i) == '#') {skipS++; i--;}
                else if (skipS > 0) {skipS--; i--;}
                else break;
            }
            while (j >= 0) { // Find position of next possible char in build(T)
                if (T.charAt(j) == '#') {skipT++; j--;}
                else if (skipT > 0) {skipT--; j--;}
                else break;
            }
            // If two actual characters are different
            if (i >= 0 && j >= 0 && S.charAt(i) != T.charAt(j))
                return false;
            // If expecting to compare char vs nothing
            if ((i >= 0) != (j >= 0))
                return false;
            i--; j--;
        }
        return true;
    }
}