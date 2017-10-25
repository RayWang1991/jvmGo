package rtdt

import (
	"fmt"
	"jvmGo/jvm/marea"
	"math"
	"math/rand"
	"testing"
)

func TestLocalVal(t *testing.T) {
	const testLen = 10
	f := Frame{
		LocalVar: marea.NewVars(testLen),
	}
	// Test int
	ints := make([]int32, testLen)
	for i := range ints {
		ints[i] = int32(rand.Uint32()) // int32
		f.LocalVar.SetInt(ints[i], uint(i))
	}
	for i := range ints {
		res, want := f.LocalVar.GetInt(uint(i)), ints[i]
		if want != res {
			t.Errorf("unmatched %s @ %d, want: %d, result: %d", "int", i, want, res)
		}
	}

	// Test float
	fs := make([]float32, testLen)
	for i := range fs {
		fs[i] = rand.Float32() // float32
		f.LocalVar.SetFloat(fs[i], uint(i))
	}
	for i := range fs {
		res, want := f.LocalVar.GetFloat(uint(i)), fs[i]
		if want != res {
			t.Errorf("unmatched %s @ %d, want: %f, result: %f", "float", i, want, res)
		}
	}

	// Test long
	ls := make([]int64, testLen/2)
	for i := range ls {
		ls[i] = int64(rand.Uint64()) // int 64
		f.LocalVar.SetLong(ls[i], uint(i*2))
		high, low := int32(ls[i]>>32), int32(ls[i])
		fmt.Printf("%d %d \n", ls[i], int64(high)<<32|int64(low)&0xFFFFFFFF)
	}
	for i := range ls {
		res, want := f.LocalVar.GetLong(uint(i*2)), ls[i]
		if want != res {
			t.Errorf("unmatched %s @ %d, want: %d, result: %d", "long", i, want, res)
		}
	}

	// Test double
	ds := make([]float64, testLen/2)
	for i := range ds {
		ds[i] = rand.Float64() // int 64
		f.LocalVar.SetDouble(ds[i], uint(i*2))
		bs := math.Float64bits(ds[i])
		high, low := int32(bs>>32), int32(bs)
		fmt.Printf("%d %d\n", math.Float64bits(ds[i]), uint64(high)<<32|uint64(low)&0xFFFFFFFF)
	}
	for i := range ds {
		res, want := f.LocalVar.GetDouble(uint(i*2)), ds[i]
		if want != res {
			t.Errorf("unmatched %s @ %d, want: %f, result: %f", "double", i, want, res)
		}
	}
}

func TestOperandStack(t *testing.T) {
	const testLen = 10
	f := Frame{
		OperandStack: NewOperandStack(uint(testLen)),
	}
	// Test int
	ints := make([]int32, testLen)
	for i := range ints {
		ints[i] = int32(rand.Uint32()) // int32
		f.OperandStack.PushInt(ints[i])
	}
	fmt.Println(f.OperandStack)
	for i := range ints {
		res, want := f.OperandStack.PopInt(), ints[testLen-1-i]
		if want != res {
			t.Errorf("unmatched %s @ %d, want: %d, result: %d", "int", i, want, res)
		}
	}

	// Test float
	fs := make([]float32, testLen)
	for i := range fs {
		fs[i] = rand.Float32() // float32
		f.OperandStack.PushFloat(fs[i])
	}
	for i := range fs {
		res, want := f.OperandStack.PopFloat(), fs[testLen-1-i]
		if want != res {
			t.Errorf("unmatched %s @ %d, want: %f, result: %f", "float", i, want, res)
		}
	}

	// Test long
	ls := make([]int64, testLen/2)
	for i := range ls {
		ls[i] = int64(rand.Uint64()) // int 64
		f.OperandStack.PushLong(ls[i])
	}
	for i := range ls {
		res, want := f.OperandStack.PopLong(), ls[testLen/2-1-i]
		if want != res {
			t.Errorf("unmatched %s @ %d, want: %d, result: %d", "long", i, want, res)
		}
	}

	// Test double
	ds := make([]float64, testLen/2)
	for i := range ds {
		ds[i] = rand.Float64() // int 64
		f.OperandStack.PushDouble(ds[i])
	}
	for i := range ds {
		res, want := f.OperandStack.PopDouble(), ds[testLen/2-1-i]
		if want != res {
			t.Errorf("unmatched %s @ %d, want: %f, result: %f", "double", i, want, res)
		}
	}
}
