package main

func main() {
	/* ERROR CODE
	var m1 map[int]bool
	m1[5] = true

	m2 := map[int]int{3: 10, 4: 40}
	m3 := map[int]int{3: 10, 4: 40}

	fmt.Println(m2 == m3)
	*/

	/*VALID CODE*/
	var m1 map[int]bool

	// ERROR -> panic: assignment to entry in nil map
	// m1[5] = true

	m2 := map[int]int{3: 10, 4: 40}
	m3 := map[int]int{3: 10, 4: 40}

	// ERROR -> invalid operation: m2 == m3 (map can only be compared to nil)
	// fmt.Println(m2 == m3)

	_, _, _ = m1, m2, m3

}
