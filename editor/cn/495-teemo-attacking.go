package main

// 495 teemo-attacking 2022-12-09 15:38:22

// leetcode submit region begin(Prohibit modification and deletion)
func findPoisonedDuration(timeSeries []int, duration int) int {
	if len(timeSeries) == 0 {
		return 0
	}
	sum := duration
	for i := 1; i < len(timeSeries); i++ {
		diff := timeSeries[i] - timeSeries[i-1]
		if diff > duration {
			diff = duration
		}
		sum += diff
	}
	return sum
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
