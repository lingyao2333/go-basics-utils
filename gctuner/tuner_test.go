package gctuner

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testHeap []byte

func TestTuner(t *testing.T) {
	is := assert.New(t)
	memLimit := uint64(100 * 1024 * 1024) //100 MB
	threshold := memLimit / 2
	tn := newTuner(threshold)
	is.Equal(tn.threshold, threshold)
	is.Equal(defaultGCPercent, tn.getGCPercent())

	// no heap
	for i := 0; i < 100; i++ {
		runtime.GC()
		is.Equal(maxGCPercent, tn.getGCPercent())
	}

	// 1/4 threshold
	testHeap = make([]byte, threshold/4)
	for i := 0; i < 100; i++ {
		runtime.GC()
		// ~= 300
		is.GreaterOrEqual(tn.getGCPercent(), uint32(250))
		is.LessOrEqual(tn.getGCPercent(), uint32(300))
	}

	// 1/2 threshold
	testHeap = make([]byte, threshold/2)
	runtime.GC()
	for i := 0; i < 100; i++ {
		runtime.GC()
		// ~= 100
		is.GreaterOrEqual(tn.getGCPercent(), uint32(50))
		is.LessOrEqual(tn.getGCPercent(), uint32(100))
	}

	// 3/4 threshold
	testHeap = make([]byte, threshold/4*3)
	runtime.GC()
	for i := 0; i < 100; i++ {
		runtime.GC()
		is.Equal(minGCPercent, tn.getGCPercent())
	}

	// out of threshold
	testHeap = make([]byte, threshold+1024)
	runtime.GC()
	for i := 0; i < 100; i++ {
		runtime.GC()
		is.Equal(minGCPercent, tn.getGCPercent())
	}
}

func TestCalcGCPercent(t *testing.T) {
	is := assert.New(t)
	const gb = 1024 * 1024 * 1024
	// use default value when invalid params
	is.Equal(defaultGCPercent, calcGCPercent(0, 0))
	is.Equal(defaultGCPercent, calcGCPercent(0, 1))
	is.Equal(defaultGCPercent, calcGCPercent(1, 0))

	is.Equal(maxGCPercent, calcGCPercent(1, 3*gb))
	is.Equal(maxGCPercent, calcGCPercent(gb/10, 4*gb))
	is.Equal(maxGCPercent, calcGCPercent(gb/2, 4*gb))
	is.Equal(uint32(300), calcGCPercent(1*gb, 4*gb))
	is.Equal(uint32(166), calcGCPercent(1.5*gb, 4*gb))
	is.Equal(uint32(100), calcGCPercent(2*gb, 4*gb))
	is.Equal(minGCPercent, calcGCPercent(3*gb, 4*gb))
	is.Equal(minGCPercent, calcGCPercent(4*gb, 4*gb))
	is.Equal(minGCPercent, calcGCPercent(5*gb, 4*gb))
}

func TestChangeGCPercent(t *testing.T) {
	is := assert.New(t)

	old := minGCPercent
	retOld := SetMinGCPercent(10)
	is.Equal(old, retOld)
	is.Equal(uint32(10), minGCPercent)

	old = maxGCPercent
	retOld = SetMaxGCPercent(1000)
	is.Equal(old, retOld)
	is.Equal(uint32(1000), maxGCPercent)
}
