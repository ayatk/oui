package main

func searchOUI(data []MAC, address string) (MAC, bool) {
	var ret MAC
	var isFound = false
	for i := range data {
		if data[i].Hex == address {
			ret = data[i]
			isFound = true
		}
	}
	return ret, isFound
}
