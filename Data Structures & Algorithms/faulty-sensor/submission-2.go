func badSensor(sensor1 []int, sensor2 []int) int {
    pointer := 0
    maxLen := len(sensor1) - 1

    // Trouver la première divergence
    for pointer < maxLen && sensor1[pointer] == sensor2[pointer] {
        pointer++
    }

    if pointer == maxLen {
        return -1
    }

    // Vérifier que la valeur ajoutée diffère de la valeur perdue
    canBeFaulty1 := sensor1[maxLen] != sensor2[pointer]
    canBeFaulty2 := sensor2[maxLen] != sensor1[pointer]

    // Vérifier les deux hypothèses
    for i := pointer; i < maxLen; i++ {
        if sensor1[i] != sensor2[i+1] {
            canBeFaulty1 = false
        }

        if sensor2[i] != sensor1[i+1] {
            canBeFaulty2 = false
        }
    }

    if canBeFaulty1 && !canBeFaulty2 {
        return 1
    }

    if canBeFaulty2 && !canBeFaulty1 {
        return 2
    }

    return -1
}
