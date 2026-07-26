type LRUCache struct {
    myCache map[int]*Node
	capacity int
	head *Node
	tail *Node
}

type Node struct {
	key int
	value int
	prev *Node
	next *Node
}

func Constructor(capacity int) LRUCache {
    return LRUCache{myCache :make(map[int]*Node,capacity), 
	capacity: capacity,
	head: nil,
	tail: nil,
	}
}

func (this *LRUCache) Get(key int) int {

    node, ok := this.myCache[key]
    if !ok {
        return -1
    }

    // déjà le plus récent
    if node == this.tail {
        return node.value
    }

    // retirer node de sa position actuelle
    if node == this.head {
        this.head = node.next
        this.head.prev = nil
    } else {
        node.prev.next = node.next
        node.next.prev = node.prev
    }

    // placer node en tail
    node.prev = this.tail
    node.next = nil
    this.tail.next = node
    this.tail = node

    return node.value
}

func (this *LRUCache) Put(key int, value int) {

    // update existing
    if node, ok := this.myCache[key]; ok {
        node.value = value

        // déplacer node en tail
        if node != this.tail {

            if node == this.head {
                this.head = node.next
                this.head.prev = nil
            } else {
                node.prev.next = node.next
                node.next.prev = node.prev
            }

            node.prev = this.tail
            node.next = nil
            this.tail.next = node
            this.tail = node
        }

        return
    }


    // création nouveau node
    newNode := &Node{
        key:key,
        value:value,
    }

    this.myCache[key] = newNode


    // premier élément
    if this.head == nil {
        this.head = newNode
        this.tail = newNode
        return
    }


    // ajouter en tail
    this.tail.next = newNode
    newNode.prev = this.tail
    this.tail = newNode


    // éjection
    if len(this.myCache) > this.capacity {

        old := this.head

        this.head = old.next
        this.head.prev = nil

        delete(this.myCache, old.key)
    }
}

