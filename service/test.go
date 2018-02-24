package main

import "fmt"

func main() {
	m := map[string]bool{ "key1":true, "key2":false }
	var k string
	var v bool
	for k,v = range m{
		fmt.Printf("key[%s] value[%d]\n", k, v)
	}

	nums := []int{2, 3, 4}
	var sum  = 0
	for _, num := range(nums){
		sum += num
	}
	fmt.Printf("sum is %d", sum)
	array := []string{"1","2","3"}
	x ,y ,z := array[0],array[1],array[2]
	fmt.Printf("%s %s %s\n", x, y, z)

}