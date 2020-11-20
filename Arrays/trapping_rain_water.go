package arrays

// Problem : Given n-non negative integers representing an elevation
// map where the width of each bar is 1, compute how much water it
// is able to trap after raining.

// TrappingRainWater1 : Naive O(n^2) solution.
// 1. Find the maximum left boundary.
// 2. Find the maximum right boundary.
// 3. Find the minimum of the 1, and 2.
// 4. Minus 3 with arr[i]
func TrappingRainWater1(arr []int, n int) (res int) {
	for i := 1; i < n-1; i++ {
		var lmax int = arr[i]
		for j := 0; j < i; j++ {
			if arr[j] > lmax {
				lmax = arr[j]
			}
		}
		var rmax int = arr[i]
		for j := i + 1; j < n; j++ {
			if arr[j] > rmax {
				rmax = arr[j]
			}
		}
		if lmax <= rmax {
			res += (lmax - arr[i])
		} else if rmax <= lmax {
			res += (rmax - arr[i])
		}
	}
	return
}

// TrappingRainWater2 : Efficient O(n) solution.
// Simillar to O(n^2) solution but precompute
// lmax and rmax.
func TrappingRainWater2(arr []int, n int) (res int) {

	lmax := make([]int, n, n)
	rmax := make([]int, n, n)

	// precompute lmax
	lmax[0] = arr[0]
	for i := 1; i < n; i++ {
		if arr[i] >= lmax[i-1] {
			lmax[i] = arr[i]
		} else {
			lmax[i] = lmax[i-1]
		}
	}

	// precompute rmax
	rmax[n-1] = arr[n-1]
	for i := n - 2; i > 0; i-- {
		if arr[i] >= rmax[i+1] {
			rmax[i] = arr[i]
		} else {
			rmax[i] = rmax[i+1]
		}
	}

	for i := 1; i < n-1; i++ {
		if lmax[i] <= rmax[i] {
			res += (lmax[i] - arr[i])
		} else if rmax[i] <= lmax[i] {
			res += (rmax[i] - arr[i])
		}
	}

	return
}
