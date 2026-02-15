package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

// main reads numbers from stdin and predicts a range for each next value.
// Strategy: keep a sliding window, compute trimmed percentiles, and use
// the median ± scaled IQR for a tight, adaptive range.
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)

	const winSize = 50
	window := make([]int, 0, winSize)

	for scanner.Scan() {
		line := scanner.Text()
		var num int
		_, err := fmt.Sscanf(line, "%d", &num)
		if err != nil {
			continue
		}

		// Maintain sliding window (FIFO).
		if len(window) < winSize {
			window = append(window, num)
		} else {
			copy(window, window[1:])
			window[winSize-1] = num
		}

		n := len(window)
		sorted := make([]int, n)
		copy(sorted, window)
		sort.Ints(sorted)

		var lower, upper int

		if n == 1 {
			padding := absInt(num)/2 + 50
			lower = num - padding
			upper = num + padding
		} else if n <= 4 {
			lo, hi := sorted[0], sorted[n-1]
			spread := hi - lo
			if spread < 2 {
				spread = 2
			}
			padding := spread/2 + 10
			lower = lo - padding
			upper = hi + padding
		} else {
			// Use trimmed data: remove top and bottom 10% to filter outliers.
			trimLo := n / 10
			trimHi := n - n/10
			if trimHi <= trimLo {
				trimLo = 0
				trimHi = n
			}
			trimmed := sorted[trimLo:trimHi]

			med := percentile(trimmed, 50)
			q1 := percentile(trimmed, 25)
			q3 := percentile(trimmed, 75)
			iqr := q3 - q1
			if iqr < 1 {
				iqr = 1
			}

			// Use 1.5x IQR — classic Tukey range covers most inlier data.
			margin := int(math.Round(1.5 * float64(iqr)))

			// Safety: margin should be at least 5% of trimmed spread.
			trimSpread := trimmed[len(trimmed)-1] - trimmed[0]
			minMargin := trimSpread / 20
			if minMargin < 1 {
				minMargin = 1
			}
			if margin < minMargin {
				margin = minMargin
			}

			lower = med - margin
			upper = med + margin
		}

		fmt.Println(lower, upper)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
}

// percentile returns the p-th percentile of a sorted slice using linear interpolation.
func percentile(arr []int, p int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return arr[0]
	}
	idx := float64(p) / 100.0 * float64(n-1)
	lo := int(math.Floor(idx))
	hi := int(math.Ceil(idx))
	if lo == hi || hi >= n {
		return arr[lo]
	}
	frac := idx - float64(lo)
	return int(math.Round(float64(arr[lo])*(1-frac) + float64(arr[hi])*frac))
}

func absInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
