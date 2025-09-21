package space

type Planet string

func Age(seconds float64, planet Planet) float64 {
    if planet == "Sun" {
        return -1.0
    }
    earthYear := 31557600.0
    planets := map[Planet]float64 {
        "Mercury": 0.2408467,
        "Venus": 0.61519726,
        "Earth": 1.0,
        "Mars": 1.8808158,
        "Jupiter": 11.862615,
        "Saturn": 29.447498,
        "Uranus": 84.016846,
        "Neptune": 164.79132,
    }
	return seconds / (earthYear * planets[planet])
}
