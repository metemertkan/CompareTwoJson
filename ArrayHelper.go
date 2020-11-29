package main

import (
	"fmt"
	"reflect"
	"sync"
	"time"

	. "github.com/ahmetb/go-linq"
)

func Equal(a, b []Pokemon) bool {
	start := time.Now()
	var result = true
	if len(a) != len(b) {
		result = false
	}

	for _, v := range a {
		var foundElement = From(b).Where(func(p interface{}) bool {
			return p.(Pokemon).Id == v.Id
		}).First()
		if !reflect.DeepEqual(v, foundElement) {
			result = false
		}
	}
	elapsed := time.Since(start)

	fmt.Println("Result =", result)
	fmt.Println("And it took ", elapsed)
	return result
}

func EqualWithMapParallel(a, b []Pokemon) bool {
	if len(a) != len(b) {
		return false
	}
	m1 := make(map[int]Pokemon)
	m2 := make(map[int]Pokemon)

	m1, m2 = convertArrayToMap(a, b)

	var result = true
	start := time.Now()
	var wg sync.WaitGroup

	wg.Add(len(a))
	for key, _ := range m1 {
		go func(key int) {
			defer wg.Done()
			if !reflect.DeepEqual(m1[key], m2[key]) {
				result = false
			}
		}(key)
	}

	wg.Wait()

	elapsed := time.Since(start)
	fmt.Println("Result =", result)
	fmt.Println("And it took ", elapsed)
	return result
}

func EqualWithMap(a, b []Pokemon) bool {

	if len(a) != len(b) {
		return false
	}
	m1 := make(map[int]Pokemon)
	m2 := make(map[int]Pokemon)

	m1, m2 = convertArrayToMap(a, b)

	start := time.Now()
	var result = true
	for key, element := range m1 {
		if !reflect.DeepEqual(element, m2[key]) {
			result = false
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Result =", result)
	fmt.Println("And it took ", elapsed)
	return result
}

func convertArrayToMap(a, b []Pokemon) (map[int]Pokemon, map[int]Pokemon) {
	m1 := make(map[int]Pokemon)
	m2 := make(map[int]Pokemon)

	for _, v := range a {
		m1[v.Id] = v
	}

	for _, x := range b {
		m2[x.Id] = x
	}
	return m1, m2
}

func reverse(array []string) []string {
	newArray := make([]string, len(array))
	for i, j := 0, len(array)-1; i <= j; i, j = i+1, j-1 {
		newArray[i], newArray[j] = array[j], array[i]
	}
	return newArray
}
