/* [5:07 PM] Phuong Nguyen
Given an integer array nums and an integer k, return the k most frequent elements.
You may return the answer in descending order. In case that 2 numbers have the same occurencies, the bigger number will be selected first.

Example:
Input: nums = [1], k = 1
Output: [1]   -  1 appears 1 time
Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]  -  1 appears 3 times, 2 appears 2 times
Input: nums = [8,2,3,5,5,2,1,5,6,8], k = 3
Output: [5, 8, 2]  -  5 appears 3 times, 2 and 8 appear 2 times each. Since 8 is bigger than 2, so 2 comes after 8.

Input: nums = [8,2,3,5,5,2,1,5,6,8, 6], k = 3
Output: [5, 8, 6]  -  5 appears 3 times, 2, 6 and 8 appear 2 times each. Since 8 > 6 and 6 > 2, so 8 and 6 are selected.

Input: nums = [1], k = 2
Output: [1]   -  1 appears 1 time. There are no more numbers in nums, so we cannot find 2 most frequent elements and we can return only [1]
Function Description:
Complete the function topKFrequent which has the following parameters:
nums: array of integers, values are duplicated
k: number of most frequent values in the response
Returns: array of k most frequent integers in nums array
[5:11 PM] Phuong Nguyen
fmt.Println(...) */

package main

import (
	"fmt"
	"sort"
)

type pair struct {
	value     int
	frequency int
}

// input should be array and k

func topKFrequent(number []int, k int) []int{

	// we have to count the frequency of each number
	// we will use map to do that

	mapfrequency := make(map[int]int)

	for i := 0; i < len(number); i++ {
		mapfrequency[number[i]] = mapfrequency[number[i]] + 1
	}  /// O(n)

	fmt.Println(mapfrequency)

	// we can convert the map into slice

	elements := make([]pair, 0, len(mapfrequency))
	for value, frequency := range mapfrequency {
		elements = append(elements, pair{value, frequency})
	} //o(n)

	fmt.Println(elements)

	// sort the slice (we have to compare frequency and return the bigger one)

	sort.Slice(elements, func(i, j int) bool {
		if elements[i].frequency == elements[j].frequency {
			return elements[i].value > elements[j].value
		}
		return elements[i].frequency > elements[j].frequency
	}) /// o(n log n)

	fmt.Println(elements)

	// return the k most frequent elements

	result := make([]int, 0, k)

	for i := 0; i < k; i++ {
		result = append(result, elements[i].value)
	} // o(n)
	fmt.Println(result)
	return result

}

//O(n log n)
//space On(n)

func main() {
	topKFrequent([]int{1, 1, 1, 2, 2, 3, 3}, 2)

	fmt.Println("hello one minute")
}
