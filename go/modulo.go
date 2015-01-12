package main

func getNextNumber(start, current, max int) int {
	var newMax = current + 1

	if newMax > max {
		return start
	}

	return current
}

func main() {
	var currentVersion = getCurrentVersion()
	var number = getNextNumber(1, currentVersion, 3)
}
