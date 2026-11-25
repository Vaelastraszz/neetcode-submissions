func badSensor(sensor1 []int, sensor2 []int) int {
	pointer1, pointer2 := 0, 0
	maxLen := len(sensor1) - 1 

	for pointer1 < maxLen && sensor1[pointer1] == sensor2[pointer2]  {
		pointer1 ++
		pointer2 ++ 
	}

	if pointer1 == maxLen {
		return -1
	}

	pointer1 = maxLen
	
	if sensor1[pointer1] == sensor2[pointer1 - 1] {
		return 2
	} else {
		return 1
	}

}
