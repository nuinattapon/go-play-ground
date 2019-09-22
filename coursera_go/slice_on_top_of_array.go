package main

import "fmt"

func main() {
	// arr index        0    1    2    3    4    5    6
	arr := [...]string{"a", "b", "c", "d", "e", "f", "g"}
	s1 := arr[1:3]
	s2 := arr[2:5]

	fmt.Printf("arr = %v, %T - len %v - cap %v\n", arr, arr, len(arr), cap(arr))
	fmt.Printf(" s1 = %v, %T - len %v - cap %v\n", s1, s1, len(s1), cap(s1))
	fmt.Printf(" s2 = %v, %T - len %v - cap %v\n", s2, s2, len(s2), cap(s2))
	s1 = append(s1, "u")
	s1 = append(s1, "v")
	s1 = append(s1, "w")
	s1 = append(s1, "x")
	s1 = append(s1, "y")
	s1 = append(s1, "z")
	s1[0] = "s"
	s1[1] = "t"
	fmt.Printf(" s1 = %v, %T - len %v - cap %v\n", s1, s1, len(s1), cap(s1))
	fmt.Printf("arr = %v, %T - len %v - cap %v\n", arr, arr, len(arr), cap(arr))

}
