package main

import "fmt"

func StringTests() {

	// in Go there is no such thing as an uninitialized variable.
	var s string
	fmt.Printf("'%s'", s)

	var bb, ff, ss = true, 2.3, "four" // bool, float64, string
	fmt.Println(bb, ff, ss)

	x := 1
	p := &x
	fmt.Printf("Value pointed from p=%d\n", *p)
	*p = 2
	fmt.Printf("Now x=%d\n", x)

}

func StructTests() {
	var tfNetworkSecurityGroups = map[string]map[string][]map[string]interface{}{
		"subnet-02-nsg": {
			"security_rule": {
				{
					"name":                       "test-01",
					"description":                "Network security group for test",
					"priority":                   110,
					"direction":                  "Inbound",
					"access":                     "Allow",
					"protocol":                   "*",
					"source_port_range":          "*",
					"destination_port_range":     "",
					"destination_port_ranges":    []string{"80", "443"},
					"source_address_prefix":      "10.0.1.0/24",
					"destination_address_prefix": "10.0.2.1",
				},
			},
		},
	}

	fmt.Print(tfNetworkSecurityGroups)
}

func sliceTests() {
	ar := [5]int{0, 1, 2, 3, 4}
	s_1 := ar[0:3]

	fmt.Println("\n===== SLICE EXAMPLES =====")
	fmt.Printf("arr: %v, len(a)=%d, cap(a)=%d\n", ar, len(ar), cap(ar))
	fmt.Printf("s_1: %v, len(s_1)=%d, cap(s_1)=%d\n", s_1, len(s_1), cap(s_1))

	s_2 := append(s_1, 100, 200, 300, 400)
	fmt.Printf("\narr: %v, len(a)=%d, cap(a)=%d\n", ar, len(ar), cap(ar))
	fmt.Printf("s_1: %v, len(s_1)=%d, cap(s_1)=%d\n", s_1, len(s_1), cap(s_1))
	fmt.Printf("s_2: %v, len(s_2)=%d, cap(s_2)=%d\n", s_2, len(s_2), cap(s_2))

	fmt.Println("\nAppend slice to another slice: s_3 := append(s_1, s_2...)")
	s_3 := append(s_1, s_2...)
	fmt.Printf("s_3: %v, len(s_3)=%d, cap(s_3)=%d\n", s_3, len(s_3), cap(s_3))

	fmt.Println()
	arr := []int{10, 20, 90, 70, 60}
	slice := make([]int, 10)
	num := copy(slice, arr)
	fmt.Println(slice)
	fmt.Println(num)

}
