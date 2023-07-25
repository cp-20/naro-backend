package main

func calculatePopulation(cities []City) map[string]int {
	populations := make(map[string]int)

	for _, city := range cities {
		if !city.CountryCode.Valid {
			continue
		}
		if _, ok := populations[city.CountryCode.String]; !ok {
			populations[city.CountryCode.String] = 0
		}
		populations[city.CountryCode.String] += int(city.Population.Int64)
	}

	return populations
}
