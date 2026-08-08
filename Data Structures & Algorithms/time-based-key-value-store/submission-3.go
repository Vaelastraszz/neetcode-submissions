type TimeMap struct {
	db map[string][]Entry
}

type Entry struct {
	value string
	timestamp int
}

func Constructor() TimeMap {
	db := make(map[string][]Entry)
	return TimeMap{db:db}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	newEntry := Entry{
		value : value,
		timestamp : timestamp,
	}
	this.db[key] = append(this.db[key], newEntry)
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if listEntries, ok := this.db[key] ; !ok {
		return ""
	} else {

		left, right := 0, len(listEntries) - 1
		var mid int
		var prevEntry Entry
		
		if timestamp > listEntries[right].timestamp {
			return listEntries[right].value
		}

		for left <= right {
			mid = (left + right)/2

			if listEntries[mid].timestamp > timestamp {
				right = mid - 1
			} else if listEntries[mid].timestamp < timestamp {
				left = mid + 1 
				prevEntry = listEntries[mid]
			} else {
				return listEntries[mid].value
			} 
		}

		return prevEntry.value
	}
	
}
