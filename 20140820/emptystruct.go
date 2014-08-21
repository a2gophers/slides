package main

var used = make(map[string]struct{})

func UnusedName(s string) bool {
	_, ok := used[s]
	return ok
}
