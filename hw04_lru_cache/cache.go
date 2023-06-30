package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItemValue struct {
	k Key
	v interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	item, ok := c.items[key]

	civ := cacheItemValue{k: key, v: value}
	if ok {
		item.Value = civ
		c.queue.MoveToFront(item)
	} else {
		item = c.queue.PushFront(civ)
		c.items[key] = item
		if c.queue.Len() > c.capacity {
			toDelete := c.queue.Back()
			c.queue.Remove(toDelete)
			delete(c.items, toDelete.Value.(cacheItemValue).k)
		}
	}

	return ok
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	var result interface{}

	item, ok := c.items[key]
	if ok {
		result = item.Value.(cacheItemValue).v
		c.queue.MoveToFront(item)
	}

	return result, ok
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}
