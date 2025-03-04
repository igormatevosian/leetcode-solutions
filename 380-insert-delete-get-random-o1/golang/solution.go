package golang

import (
	"math/rand"
)

type RandomizedSet struct {
	indexes map[int]int
	arr     []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{make(map[int]int), make([]int, 0)}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.indexes[val]; ok {
		return false
	}
	this.indexes[val] = len(this.arr)
	this.arr = append(this.arr, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.indexes[val]
	if !ok {
		return false
	}
	lastIndex := len(this.arr) - 1
	lastEl := this.arr[lastIndex]
	this.arr[index] = lastEl
	this.indexes[lastEl] = index

	this.arr = this.arr[:lastIndex]
	delete(this.indexes, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.arr[rand.Intn(len(this.arr))]
}
