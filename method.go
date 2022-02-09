package tutorial

// BrokenMethod has a bug - it will try to read the 4th
// index of Data even when it only has a length of 3.
func BrokenMethod(Data string) bool {
	return len(Data) >= 3 &&
		Data[0] == 'F' &&
		Data[1] == 'O' &&
		Data[2] == 'U' &&
		Data[3] == 'R'
}


func NumberInputMethod(Data1 int32, Data2 uint64, Data3 float32) int {
    if Data1 < 0 || Data3 < 0 {
        return -1
    }

    if Data2 > 10 {
        return 1
    }

    return 0
}

