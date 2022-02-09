package tutorial

import "testing"  

func FuzzBrokenMethod(f *testing.F) {
  f.Add("FUZ")
  
  f.Fuzz(func(t *testing.T, str string) {
    BrokenMethod(str)
  })
}

func FuzzBrokenMethodNoSeed(f *testing.F) {
  f.Fuzz(func(t *testing.T, str string) {
    BrokenMethod(str)
  })
}

func FuzzTestNumberInputMethod(f *testing.F) {
  f.Fuzz(func(t *testing.T, d1 int32, d2 uint64, d3 float32) {
    NumberInputMethod(d1, d2, d3)
  })
}
