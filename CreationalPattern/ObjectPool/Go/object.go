package objectpool

// Object represents an object in the pool.
type Object struct {
	ID int
}

// ObjectPool is a struct that represents an object pool.
type ObjectPool struct {
	objects chan *Object
}

// NewObjectPool creates a new object pool with the specified size.
func NewObjectPool(size int) *ObjectPool {
	// Create a buffered channel to hold the objects in the pool.
	pool := &ObjectPool{
		objects: make(chan *Object, size),
	}

	for i := 0; i < size; i++ {
		// Create a new object and add it to the pool.
		pool.objects <- &Object{ID: i}
	}

	return pool
}

func (p *ObjectPool) AcquireObject() *Object {
	// Wait for an object to become available in the pool.
	return <-p.objects
}

func (p *ObjectPool) ReleaseObject(obj *Object) {
	// Return the object to the pool.
	p.objects <- obj
}
