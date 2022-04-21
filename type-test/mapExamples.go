package main
import ("fmt")

// https://go101.org/article/container.html

func mapExamples(){

	// 
	my_map := make(map[int]int)
	my_map[2] = 4
	my_map[4] = 16
	my_map[8] = 64
	delete(my_map, 4)
	fmt.Println(my_map)
	fmt.Printf("================================================\n")

	// Initialize the map object to append slices to it later
	var tfNetworkSecurityGroups = make(map[string]map[string][]map[string]interface{})

	// A map{string->interface} will be a record of the slice
	rule_01 := map[string]interface{}{
		"name":                       "test",
		"description":                "Network security group for test",
		"priority":                   "110",
		"direction":                  "Inbound",
		"access":                     "Allow",
		"protocol":                   "*",
		"source_port_range":          "*",
		"destination_port_ranges":    []string{"80", "443"},
		"source_address_prefix":      "10.0.1.0/24",
		"destination_address_prefix": "10.0.2.1",
	}

	tfNetworkSecurityGroups["subnet-02-nsg"] = map[string][]map[string]interface{}{}
	tfNetworkSecurityGroups["subnet-02-nsg"]["security_rule"] = []map[string]interface{}{}
	fmt.Println(tfNetworkSecurityGroups["subnet-02-nsg"]["security_rule"])
	fmt.Println(rule_01)

	tfNetworkSecurityGroups["subnet-02-nsg"]["security_rule"] = append(
		[]map[string]interface{}{rule_01},
		tfNetworkSecurityGroups["subnet-02-nsg"]["security_rule"]...,
	)
	fmt.Println(tfNetworkSecurityGroups["subnet-02-nsg"]["security_rule"])
	fmt.Printf("%v \n", tfNetworkSecurityGroups)
	fmt.Printf("================================================\n")

	// -------- Super object inline
	var tfNetworkSecurityGroups_02 = map[string]map[string][]map[string]interface{}{
		"subnet": {
			"security_rule": {
				{
					"name":                       "test-01",
					"description":                "Network security group for test",
					"priority":                   "110",
					"direction":                  "Inbound",
					"access":                     "Allow",
					"protocol":                   "*",
					"source_port_range":          "*",
					"destination_port_ranges":    []string{"80", "443"},
					"source_address_prefix":      "10.0.1.0/24",
					"destination_address_prefix": "10.0.2.1",
				},
				{
					"name":                       "test-02",
					"description":                "Deny access",
					"priority":                   111,
					"direction":                  "Inbound",
					"access":                     "Deny",
					"protocol":                   "*",
					"source_port_range":          "*",
					"destination_port_ranges":    "*",
					"source_address_prefix":      "10.0.1.0/24",
					"destination_address_prefix": "10.0.2.1",
				},
			},
		},
	}
	fmt.Printf("%v \n", tfNetworkSecurityGroups_02)
	fmt.Printf("================================================\n")
}