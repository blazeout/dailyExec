package main

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	mp := map[string]bool{}
	for _, v := range supplies {
		mp[v] = true
	}
	res := []string{}
	flag := false
	for i := 0; i < len(ingredients); i++ {
		for _, v := range ingredients[i] {
			if _, ok := mp[v]; !ok {
				flag = false
				break
			} else {
				flag = true
			}
		}
		if flag {
			res = append(res, recipes[i])
			mp[recipes[i]] = true
		}
	}
	return res
}
