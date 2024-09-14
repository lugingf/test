package old_examples

import (
	"testing"
)

func BenchmarkSumGen(b *testing.B) {
	intMap := map[string]int64{
		"key1":  12,
		"key2":  24,
		"key3":  48,
		"key4":  32,
		"key5":  123,
		"key6":  536,
		"key7":  2345,
		"key8":  234,
		"key9":  12,
		"key10": 124,
	}
	floatMap := map[string]float64{
		"key1":  12.21,
		"key2":  24.42,
		"key3":  48.84,
		"key4":  32.23,
		"key5":  123.32,
		"key6":  536.65,
		"key7":  2345.43,
		"key8":  234.34,
		"key9":  12.1,
		"key10": 124.8,
	}
	for n := 0; n < b.N; n++ {
		SumNumbers(intMap)
		SumNumbers(floatMap)
	}
}

func BenchmarkSumInterface(b *testing.B) {
	intMap := map[string]interface{}{
		"key1":  12,
		"key2":  24,
		"key3":  48,
		"key4":  32,
		"key5":  123,
		"key6":  536,
		"key7":  2345,
		"key8":  234,
		"key9":  12,
		"key10": 124,
	}
	floatMap := map[string]interface{}{
		"key1":  12.21,
		"key2":  24.42,
		"key3":  48.84,
		"key4":  32.23,
		"key5":  123.32,
		"key6":  536.65,
		"key7":  2345.43,
		"key8":  234.34,
		"key9":  12.1,
		"key10": 124.8,
	}
	for n := 0; n < b.N; n++ {
		SumNumbersInterface(intMap)
		SumNumbersInterface(floatMap)
	}
}
