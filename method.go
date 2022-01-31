package tutorial

// BrokenMethod has a bug - it will try to read the 4th
// index of Data even when it only has a length of 3.
func BrokenMethod(Data string) bool {
	return len(Data) >= 3 &&
		Data[3] == 'F' &&
		Data[2] == 'I' &&
		Data[1] == 'T' &&
		Data[0] == 'V'
}
