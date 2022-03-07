package tutorial

import "testing"  

func FuzzBrokenMethod(f *testing.F) {
  f.Add("FUZ")
  
  f.Fuzz(func(t *testing.T, str string) {
    BrokenMethod(str)
  })
}

func FuzzBrokenMethodNoSeed(f *testing.F) {
  f.Add("F")
  f.Fuzz(func(t *testing.T, str string) {
    BrokenMethod(str)
  })
}

func FuzzTestNumberInputMethod(f *testing.F) {
  f.Add(true, int(1), "FU")
  f.Fuzz(func(t *testing.T, data1 bool, data2 int, data3 string) {
    NumberInputMethod(data1, data2, data3)
  })
}
