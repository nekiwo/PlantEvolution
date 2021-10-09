package plant

import "github.com/nekiwo/PlantEvolution/src/config"


func GenPlant(ID int) Plant {
	plant :=  Plant{
		len(config.AllGens),
		ID,
		3,
		make([]float64, 0),
		0,
	}

	return  plant
}