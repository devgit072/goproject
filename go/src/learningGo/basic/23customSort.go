package main

import (
	"sort"
	"fmt"
)

type byLength []string // its just a alias, byLength is alias of []string
// sort.Interface has three methods: Len, Less, Swap, u can implement your own version

func (l byLength) Len() int {// its is just a way to register that function in []string so that u can call s.Len()
	return len(l)
}

func (s byLength) Less(i,j int) bool {
	return len(s[i]) < len(s[j])
}
func (s byLength) Swap(i, j int) {
	s[i],s[j] = s[j],s[i]
}
func main()  {
	strs := []string{"ddd", "ccddddddddddddaasdasda", "abce", "kkjdasask", "ssaadasd445ds", "hhtstr", "i"}
	strs1 := []string{"ddd", "ccddddddddddddaasdasda", "abce", "kkjdasask", "ssaadasd445ds", "hhtstr", "i"}
	sort.Strings(strs)
	fmt.Println(strs)
	sort.Sort(byLength(strs1))
	fmt.Println(byLength(strs1))
}