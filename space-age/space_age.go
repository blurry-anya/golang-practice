package space

type Planet string

func Age(seconds float64, planet Planet) float64 {
	switch planet {
	case "Mercury":
		return seconds / getSeconds(0.2408467)
	case "Venus":
		return seconds / getSeconds(0.61519726)
	case "Earth":
		return seconds / getSeconds(1)
	case "Mars":
		return seconds / getSeconds(1.8808158)
	case "Jupiter":
		return seconds / getSeconds(11.862615)
	case "Saturn":
		return seconds / getSeconds(29.447498)
	case "Uranus":
		return seconds / getSeconds(84.016846)
	case "Neptune":
		return seconds / getSeconds(164.79132)
	}

	return -1
}

func getSeconds(earthYears float64) float64 {
	return earthYears * 365.25 * 86400
}
