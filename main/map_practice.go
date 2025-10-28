package main

import (
	"fmt"
	"sort"
)

func make_maps() {
	var mapLit map[string]int
	var mapAssigned map[string]int
	mapLit = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32)
	mapAssigned = mapLit
	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.1415926
	mapAssigned["two"] = 3
	value, ok := mapLit["ten"]
	fmt.Printf("mapLit 的以一个值为:%d\n", mapLit["one"])
	fmt.Printf("mapCreated的第二个值为:%f\n", mapCreated["two"])
	fmt.Printf("mapAssigned的第二个值为:%d\n", mapAssigned["key2"])
	fmt.Printf("mapLit的第10个值为:%d\n", mapLit["ten"])
	fmt.Printf("MapLit['ten'] = %d,是否存在=%t\n", value, ok)
}

func testelement() {
	var map1 map[string]int
	map1 = map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6}
	delete(map1, "ten")
	if _, ok := map1["one"]; ok {

	}
	for key, value := range map1 {
		fmt.Printf("key:%s\tvalue:%d\n", key, value)
	}

}
func map_days() {
	weekdays := map[int]string{
		1: "星期一", // 或 "Monday"
		2: "星期二", // 或 "Tuesday"
		3: "星期三", // 或 "Wednesday"
		4: "星期四", // 或 "Thursday"
		5: "星期五", // 或 "Friday"
		6: "星期六", // 或 "Saturday"
		7: "星期日", // 或 "Sunday"
	}
	for key, value := range weekdays {
		if value == "星期二" {
			fmt.Printf("星期二是第:%d天\n", key)
		}
	}
}

func maps_forrange2() {
	items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][i] = i
	}
	fmt.Println(items)
}

func sort_maps() {
	var (
		barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
			"delta": 87, "echo": 56, "foxtrot": 12,
			"golf": 34, "hotel": 16, "indio": 87,
			"juliet": 65, "kili": 43, "lima": 98}
	)
	fmt.Println("未排序的:")
	for k, _ := range barVal {
		fmt.Println(k)
	}
	fmt.Println("==============")
	keys := make([]string, len(barVal))
	i := 0
	for k, _ := range barVal {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println()
	fmt.Println("排序后")
	for _, k := range keys {
		fmt.Println(k)
	}
}
func map_drink() {
	drinkmap := map[string]string{
		"Coffee":   "咖啡",
		"Tea":      "茶",
		"Water":    "水",
		"Milk":     "牛奶",
		"Juice":    "果汁",
		"Beer":     "啤酒",
		"Wine":     "葡萄酒",
		"Smoothie": "冰沙",
		"Cocktail": "鸡尾酒",
		"Lemonade": "柠檬水",
		"Soda":     "苏打水",
		"Espresso": "浓缩咖啡",
	}
	fmt.Println("排序前")
	for k, v := range drinkmap {
		fmt.Printf("英文:%v,中文:%v\n", k, v)
	}
	engsort := make([]string, len(drinkmap))
	i := 0
	for k, _ := range drinkmap {
		engsort[i] = k
		i++
	}
	fmt.Println(engsort)
	sort.Strings(engsort)
	fmt.Println("排序后:")
	for _, k := range engsort {
		fmt.Printf("英文:%v,中文:%v\n", k, drinkmap[k])
	}
}
