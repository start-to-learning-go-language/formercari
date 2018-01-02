for i := 1; i < 100; i++ {
	if i/2 != 0 {
		fmt.Println(i)
	}
}

i := 0
for {
	i++
	if i >= 100 {
		break
	} else if i/2 == 0 {
		continue
	}
	fmt.Println(i)
}

dayOfWeeks := [...]string{"月", "火", "水", "木", "金", "土", "日"}
for arrayIndex, dayOfWeek := range dayOfWeeks {
	fmt.Printf("%d番目の曜日は%s曜日です。\n", arrayIndex+1, dayOfWeek)
}