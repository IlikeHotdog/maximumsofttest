package main

import (
	"fmt"
)

func main() {
	arr1 := []interface{}{"a", "a", "a", "b", "b", "c"}
	arr2 := []interface{}{"b", "c", "d", "d", "d"}
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
func removeDuplicates(arr []interface{}) []interface{} {
	m := make(map[interface{}]bool)
	uniqueArr := []interface{}{}
	for _, item := range arr {
		m[item] = true
	}
	for key := range m {
		uniqueArr = append(uniqueArr, key)
	}
	return uniqueArr
}
func cal_array(ar1 []interface{}, ar2 []interface{}) ([]interface{}, []interface{}) {
	//union
	m := make(map[interface{}]bool)
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
	different := []interface{}{}
	m = make(map[interface{}]bool)

	for _, item := range ar2 {
		m[item] = true
	}
	for _, item := range ar1 {
		if _, ok := m[item]; !ok {
			different = append(different, item)
			m[item] = true
		}
	}

	m = make(map[interface{}]bool)
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
