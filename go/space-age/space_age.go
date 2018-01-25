package space

type Planet = string

func Age(s float64, p Planet) float64 {
	offset := 1.0
	switch p {
	case "Earth":
		offset = 1.0
		break
	case "Mercury":
		offset = 0.2408467
		break
	case "Venus":
		offset = 0.61519726
		break
	case "Mars":
		offset = 1.8808158
		break
	case "Jupiter":
		offset = 11.862615
		break
	case "Saturn":
		offset = 29.447498
		break
	case "Uranus":
		offset = 84.016846
		break
	case "Neptune":
		offset = 164.79132
		break
	}
	return s / 365.25 / 86400 / offset
}
