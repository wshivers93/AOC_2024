# Advent of Code 2024

This repo contains solutions for [Advent of Code 2024](https://adventofcode.com/2024). I've been wanting to learn Go, so this year that's what I'm using to solve the problems.

### Day 1

Day 1 was pretty straightforward, just a matter of splitting the input into two sorted slices and comparing the two. I did find it interesting Go doesn't have a built-in method for finding all occurrences of a value in a slice (at least I couldn't find one). My Typescript brain wanted to write `arr.filter((item) => item === val)` and then just get the length of the returned array. Instead wrote my own method to achieve this.

```
func countAll(val int64, arr []int64) int64 {
	var count int64

	for _, i := range arr {
		if val == i {
			count++
		}

		if val < i {
			break
		}
	}

	return count
}
```
