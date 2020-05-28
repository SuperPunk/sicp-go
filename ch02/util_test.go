package ch02

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccumulate(t *testing.T) {
	list := EnumerateInterval(1, 100)
	sum := Accumulate(func(i interface{}, i2 interface{}) interface{} {
		intI, _ := i.(int)
		intI2, _ := i2.(int)
		return intI + intI2
	}, 0, list)
	assert.Equal(t, 5050, sum)
}

// 2->3->4   =>
// (1)->(1->2)->(1->2->3)  =>
// ((1,2))->((1,3)->(2,3))->((1,4)->(2,4)->(3,4))
// FlatMap的操作，通过调用Append将每一个嵌套的Node连在一起
func TestGeneratePair(t *testing.T) {
	list := MapList(func(i interface{}) interface{} {
		intI, _ := i.(int)
		return MapList(func(j interface{}) interface{} {
			intJ, _ := j.(int)
			return NewPair(intJ, intI)
		}, EnumerateInterval(1, intI-1))
	}, EnumerateInterval(2, 4))

	var node = ListRef(list, 0).(*Node)
	var pair = ListRef(node, 0).(Pair)
	assert.Equal(t, 1, Left(pair))
	assert.Equal(t, 2, Right(pair))

	node = ListRef(list, 1).(*Node)
	pair = ListRef(node, 0).(Pair)
	assert.Equal(t, 1, Left(pair))
	assert.Equal(t, 3, Right(pair))
	pair = ListRef(node, 1).(Pair)
	assert.Equal(t, 2, Left(pair))
	assert.Equal(t, 3, Right(pair))

	node = ListRef(list, 2).(*Node)
	pair = ListRef(node, 0).(Pair)
	assert.Equal(t, 1, Left(pair))
	assert.Equal(t, 4, Right(pair))
	pair = ListRef(node, 1).(Pair)
	assert.Equal(t, 2, Left(pair))
	assert.Equal(t, 4, Right(pair))
	pair = ListRef(node, 2).(Pair)
	assert.Equal(t, 3, Left(pair))
	assert.Equal(t, 4, Right(pair))

	// 此处是其实是一个flatMap的操作
	var pairList = Accumulate(func(i interface{}, j interface{}) interface{} {
		NodeI, _ := i.(*Node)
		NodeJ, _ := j.(*Node)
		return Append(NodeI, NodeJ)
	}, nil, list)
	pairListRoot := pairList.(*Node)

	assert.Equal(t, 1, Left(ListRef(pairListRoot, 0).(Pair)))
	assert.Equal(t, 2, Right(ListRef(pairListRoot, 0).(Pair)))
	assert.Equal(t, 1, Left(ListRef(pairListRoot, 1).(Pair)))
	assert.Equal(t, 3, Right(ListRef(pairListRoot, 1).(Pair)))
	assert.Equal(t, 2, Left(ListRef(pairListRoot, 2).(Pair)))
	assert.Equal(t, 3, Right(ListRef(pairListRoot, 2).(Pair)))
	assert.Equal(t, 1, Left(ListRef(pairListRoot, 3).(Pair)))
	assert.Equal(t, 4, Right(ListRef(pairListRoot, 3).(Pair)))
	assert.Equal(t, 2, Left(ListRef(pairListRoot, 4).(Pair)))
	assert.Equal(t, 4, Right(ListRef(pairListRoot, 4).(Pair)))
	assert.Equal(t, 3, Left(ListRef(pairListRoot, 5).(Pair)))
	assert.Equal(t, 4, Right(ListRef(pairListRoot, 5).(Pair)))

	fmt.Println(pairList)
}
