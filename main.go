package main

import (
	"fmt"
	"sort"
)

func main() {
	arr1 := []string{"a", "a", "a", "b", "b", "c"}
	arr2 := []string{"b", "c", "d", "d", "d"}
	fmt.Println("Input array : ")
	fmt.Println("array 1", arr1)
	fmt.Println("array 2", arr2)
	arr1 = removeDuplicates(arr1)
	arr2 = removeDuplicates(arr2)
	fmt.Println("Unique array : ")
	fmt.Println("array 1", arr1)
	fmt.Println("array 2", arr2)
	result1, result2 := cal_array(arr1, arr2)
	fmt.Println("Union     :", result1)
	fmt.Println("Difference:", result2)
}
func removeDuplicates(arr []string) []string {
	m := make(map[string]bool)
	uniqueArr := []string{}
	for _, item := range arr {
		m[item] = true
	}
	for key := range m {
		uniqueArr = append(uniqueArr, key)
	}
	sort.Strings(uniqueArr)
	return uniqueArr
}
func cal_array(ar1 []string, ar2 []string) ([]string, []string) {
	//union
	m := make(map[string]bool)
	union := ar1
	for _, item := range union {
		m[item] = true
	}
	for _, item := range ar2 {
		if _, ok := m[item]; !ok {
			union = append(union, item)

		}
	}
	//diff
	different := []string{}
	m = make(map[string]bool)

	for _, item := range ar2 {
		m[item] = true
	}
	for _, item := range ar1 {
		if _, ok := m[item]; !ok {
			different = append(different, item)
			m[item] = true
		}
	}

	m = make(map[string]bool)
	for _, item := range ar1 {
		m[item] = true
	}
	for _, item := range ar2 {
		if _, ok := m[item]; !ok {
			different = append(different, item)
			m[item] = true
		}
	}
	return union, different
}
