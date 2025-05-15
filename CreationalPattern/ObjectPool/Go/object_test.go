package objectpool

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Object(t *testing.T) {
	assert := assert.New(t)

	pool := NewObjectPool(3)

	obj1 := pool.AcquireObject()
	assert.Equal(0, obj1.ID)

	obj2 := pool.AcquireObject()
	assert.Equal(1, obj2.ID)

	obj3 := pool.AcquireObject()
	assert.Equal(2, obj3.ID)
}
