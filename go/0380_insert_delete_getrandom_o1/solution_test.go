package insert_delete_getrandom_o1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SuccessCase1(t *testing.T) {
	randomizedSet := Constructor()

	assert.True(t, randomizedSet.Insert(1))  // Inserts 1 to the set. Returns true as 1 was inserted successfully
	assert.False(t, randomizedSet.Remove(2)) // Returns false as 2 does not exist in the set
	assert.True(t, randomizedSet.Insert(2))  // Inserts 2 to the set, returns true. Set now contains [1,2]

	assert.Contains(t, []int{1, 2}, randomizedSet.GetRandom()) // getRandom() should return either 1 or 2 randomly

	assert.True(t, randomizedSet.Remove(1))  // Removes 1 from the set, returns true. Set now contains [2]
	assert.False(t, randomizedSet.Insert(2)) // 2 was already in the set, so return false

	assert.Equal(t, 2, randomizedSet.GetRandom()) // Since 2 is the only number in the set, getRandom() will always return 2
}

func Test_SuccessCase2(t *testing.T) {
	randomizedSet := Constructor()

	assert.True(t, randomizedSet.Insert(0))
	assert.True(t, randomizedSet.Insert(1))
	assert.True(t, randomizedSet.Remove(0))

	assert.True(t, randomizedSet.Insert(2))
	assert.True(t, randomizedSet.Remove(1))

	assert.Equal(t, 2, randomizedSet.GetRandom())
}

func Test_SuccessCase3(t *testing.T) {
	randomizedSet := Constructor()

	assert.False(t, randomizedSet.Remove(0))
	assert.False(t, randomizedSet.Remove(0))
	assert.True(t, randomizedSet.Insert(0))

	assert.Equal(t, 0, randomizedSet.GetRandom())

	assert.True(t, randomizedSet.Remove(0))
	assert.True(t, randomizedSet.Insert(0))
}

func Test_SuccessCase4(t *testing.T) {
	randomizedSet := Constructor()

	assert.True(t, randomizedSet.Insert(3))
	assert.False(t, randomizedSet.Insert(3))

	assert.Equal(t, 3, randomizedSet.GetRandom())
	assert.Equal(t, 3, randomizedSet.GetRandom())

	assert.True(t, randomizedSet.Insert(1))
	assert.True(t, randomizedSet.Remove(3))

	assert.Equal(t, 1, randomizedSet.GetRandom())
	assert.Equal(t, 1, randomizedSet.GetRandom())

	assert.True(t, randomizedSet.Insert(0))
	assert.True(t, randomizedSet.Remove(0))
}
