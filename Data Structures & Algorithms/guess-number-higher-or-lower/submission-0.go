func guessNumber(n int) int {
    left, right := 1, n

    for left <= right {
        mid := left + (right-left)/2

        switch guess(mid) {
        case 0:
            return mid
        case -1:
            right = mid - 1
        case 1:
            left = mid + 1
        }
    }

    return -1
}