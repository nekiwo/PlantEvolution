package plant

import "github.com/nekiwo/PlantEvolution/src/config"


func GenPlant(ID int) Plant {
	plant := Plant{
		len(config.AllGens),
		ID,
		3,
		GenSegments(make([]float64, 0), 3),
		0,
	}

	return  plant
}