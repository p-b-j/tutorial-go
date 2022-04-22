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


func NumberInputMethod(Data1 bool, Data2 int, Data3 string) int {
    if Data1 {
        return 1
    }

    if Data2 == 0 {
        return 5 / Data2
    }

    if Data2 == 1 {
        return 5 / (Data2 - 1)
    }

    if len(Data3) == 1 {
        return int(Data3[1])
    }

    if len(Data3) == 0 {
        return int(Data3[0])
    }

    return 0
}

