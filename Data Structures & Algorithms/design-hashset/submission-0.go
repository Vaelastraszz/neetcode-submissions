type MyHashSet struct {
	data map[int]bool
}

func Constructor() MyHashSet {
    myData := make(map[int]bool)
	return MyHashSet{
		data : myData,
	}
}

func (this *MyHashSet) Add(key int) {
    this.data[key] = true
}

func (this *MyHashSet) Remove(key int) {
    delete(this.data, key)
}

func (this *MyHashSet) Contains(key int) bool {
    if _, ok := this.data[key]; ok {
		return true
	}

	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
 