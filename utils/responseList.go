package utils

// Response Air Pollution
type ResponseAirPollution struct {
	Coord struct {
		Lon float32
		Lat float32
	}
	List [1]struct {
		Main struct {
			Aqi float32
		}
		Components struct {
			Co    float32
			No    float32
			No2   float32
			O3    float32
			So2   float32
			Pm2_5 float32
			Pm10  float32
			Nh3   float32
		}
		Dt float32
	}
}
