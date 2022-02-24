package tutorial

// BrokenMethod has a bug - it will try to read the 4th
// index of Data even when it only has a length of 3.
func BrokenMethod(Data string) bool {
	return len(Data) >= 3 &&
		Data[0] == 'F' &&
		Data[1] == 'I' &&
		Data[2] == 'Z'
}


func NumberInputMethod(Data bool) int {
    if Data {
        return -1
    }

    return 0
}

