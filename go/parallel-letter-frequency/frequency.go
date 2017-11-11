package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func sumFrequencyMaps(a, b FreqMap) FreqMap{

	for k, v := range b {
		a[k] += v
	}

	return a
}


func ConcurrentFrequency(txt []string) FreqMap {

	c := make(chan FreqMap)

	for _, s := range txt {

		go func(s string, c chan FreqMap){
			c <- Frequency(s)
		}(s, c)
		
	}
	
	res := FreqMap{}
	for range txt {
		res = sumFrequencyMaps(res, <-c)
	}

	close(c)

	return res
}