type LRUCache struct {
    capacity int
    indexMap map[int]*CacheNode
    head, tail *CacheNode
}

type CacheNode struct {
    key, value int
    next, prev *CacheNode
}

func Constructor(capacity int) LRUCache {
    var cache LRUCache
    var head, tail CacheNode
    cache.capacity = capacity
    cache.indexMap = make(map[int]*CacheNode)
    cache.head, cache.tail = &head, &tail
    cache.head.next = cache.tail
    cache.tail.prev = cache.head
    return cache
}


func (this *LRUCache) Get(key int) int {
    if vnode, ok := this.indexMap[key]; ok {
        vnode.prev.next = vnode.next
        vnode.next.prev = vnode.prev
        vnode.next = this.tail
        vnode.prev = this.tail.prev
        vnode.next.prev = vnode
        vnode.prev.next = vnode
        return vnode.value
    } else {
        return -1
    }
}


func (this *LRUCache) Put(key int, value int)  {
    if vnode, ok := this.indexMap[key]; ok {
        vnode.prev.next = vnode.next
        vnode.next.prev = vnode.prev
        vnode.next = this.tail
        vnode.prev = this.tail.prev
        vnode.next.prev = vnode
        vnode.prev.next = vnode
        vnode.value = value
    } else {
        var node CacheNode
        p := &node
        p.next = this.tail
        p.prev = this.tail.prev
        p.next.prev = p
        p.prev.next = p
        p.key, p.value = key, value
        this.indexMap[key] = p
        if len(this.indexMap) > this.capacity {
            delete(this.indexMap, this.head.next.key)
            this.head.next = this.head.next.next
            this.head.next.prev = this.head
        }
    }
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */