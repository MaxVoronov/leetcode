package insert_delete_getrandom_o1

import "math/rand"

type RandomizedSet struct {
	items []int
	table map[int]int
}

func Constructor() RandomizedSet {
	const initSize = 25

	return RandomizedSet{
		items: make([]int, 0, initSize),
		table: make(map[int]int, initSize),
	}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, found := rs.table[val]; found {
		return false
	}

	rs.items = append(rs.items, val)
	rs.table[val] = len(rs.items) - 1

	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	idx, found := rs.table[val]
	if !found {
		return false
	}

	delete(rs.table, val)                     // Remove val element from map
	rs.items[idx] = rs.items[len(rs.items)-1] // Swap val with last element in array
	rs.items[len(rs.items)-1] = 0             // Avoid memory leak
	rs.items = rs.items[:len(rs.items)-1]     // Remove last element from array

	if idx < len(rs.items) {
		rs.table[rs.items[idx]] = idx // Update index for moved last element in map
	}

	return true
}

func (rs *RandomizedSet) GetRandom() int {
	randIdx := rand.Intn(len(rs.items))

	return rs.items[randIdx]
}
