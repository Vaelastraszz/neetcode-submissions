type Car struct {
	position int
	speed    int
}

func carFleet(target int, position []int, speed []int) int {
	cars := make([]Car, 0, len(position))
	stack := make([]float64, 0, len(position))

	for i := 0; i < len(position); i++ {
		cars = append(cars, Car{
			position: position[i],
			speed:    speed[i],
		})
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].position > cars[j].position
	})

	for _, car := range cars {
		timeToTarget := float64(target-car.position) / float64(car.speed)

		if len(stack) == 0 {
			stack = append(stack, timeToTarget)
			continue
		}

		previousFleetTime := stack[len(stack)-1]

		if timeToTarget <= previousFleetTime {
			continue
		}

		stack = append(stack, timeToTarget)
	}

	return len(stack)
}