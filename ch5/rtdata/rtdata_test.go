package rtdata

import (
	"testing"
	"math/rand"
	"fmt"
	"math"
)

func TestLocalVal(t *testing.T) {
	const testLen = 10
	f := Frame{
		localVar: NewLocalVars(testLen),
	}
	// Test int
	ints := make([]int32, testLen)
	for i := range ints {
		ints[i] = int32(rand.Uint32()) // int32
		f.localVar.SetInt(ints[i], uint(i))
	}
	for i := range ints {
		res, want := f.localVar.GetInt(uint(i)), ints[i]
		if want != res {
			t.Errorf("unmatched %s @ %d, want: %d, result: %d", "int", i, want, res)
		}
	}

	// Test float
	fs := make([]float32, testLen)
	for i := range fs {
		fs[i] = rand.Float32() // float32
		f.localVar.SetFloat(fs[i], uint(i))
	}
	for i := range fs {
		res, want := f.localVar.GetFloat(uint(i)), fs[i]
		if want != res {
			t.Errorf("unmatched %s @ %d, want: %f, result: %f", "float", i, want, res)
		}
	}

	// Test long
	ls := make([]int64, testLen/2)
	for i := range ls {
		ls[i] = int64(rand.Uint64()) // int 64
		f.localVar.SetLong(ls[i], uint(i*2))
		high, low := int32(ls[i]>>32), int32(ls[i])
		fmt.Printf("%d %d \n", ls[i], int64(high)<<32|int64(low)&0xFFFFFFFF)
	}
	for i := range ls {
		res, want := f.localVar.GetLong(uint(i*2)), ls[i]
		if want != res {
			t.Errorf("unmatched %s @ %d, want: %d, result: %d", "long", i, want, res)
		}
	}

	// Test double
	ds := make([]float64, testLen/2)
	for i := range ds {
		ds[i] = rand.Float64() // int 64
		f.localVar.SetDouble(ds[i], uint(i*2))
		bs := math.Float64bits(ds[i])
		high, low := int32(bs>>32), int32(bs)
		fmt.Printf("%d %d\n", math.Float64bits(ds[i]), uint64(high)<<32|uint64(low)&0xFFFFFFFF)
	}
	for i := range ds {
		res, want := f.localVar.GetDouble(uint(i*2)), ds[i]
		if want != res {
			t.Errorf("unmatched %s @ %d, want: %f, result: %f", "double", i, want, res)
		}
	}
}

func TestOperandStack(t *testing.T) {
	const testLen = 10
	f := Frame{
		operandStack: NewOperandStack(uint(testLen)),
	}
	// Test int
	ints := make([]int32, testLen)
	for i := range ints {
		ints[i] = int32(rand.Uint32()) // int32
		f.operandStack.PushInt(ints[i])
	}
	fmt.Println(f.operandStack)
	for i := range ints {
		res, want := f.operandStack.PopInt(), ints[testLen-1-i]
		if want != res {
			t.Errorf("unmatched %s @ %d, want: %d, result: %d", "int", i, want, res)
		}
	}

	// Test float
	fs := make([]float32, testLen)
	for i := range fs {
		fs[i] = rand.Float32() // float32
		f.operandStack.PushFloat(fs[i])
	}
	for i := range fs {
		res, want := f.operandStack.PopFloat(), fs[testLen-1-i]
		if want != res {
			t.Errorf("unmatched %s @ %d, want: %f, result: %f", "float", i, want, res)
		}
	}

	// Test long
	ls := make([]int64, testLen/2)
	for i := range ls {
		ls[i] = int64(rand.Uint64()) // int 64
		f.operandStack.PushLong(ls[i])
	}
	for i := range ls {
		res, want := f.operandStack.PopLong(), ls[testLen/2-1-i]
		if want != res {
			t.Errorf("unmatched %s @ %d, want: %d, result: %d", "long", i, want, res)
		}
	}

	// Test double
	ds := make([]float64, testLen/2)
	for i := range ds {
		ds[i] = rand.Float64() // int 64
		f.operandStack.PushDouble(ds[i])
	}
	for i := range ds {
		res, want := f.operandStack.PopDouble(), ds[testLen/2-1-i]
		if want != res {
			t.Errorf("unmatched %s @ %d, want: %f, result: %f", "double", i, want, res)
		}
	}
}
