package plant

import "github.com/nekiwo/PlantEvolution/src"


func GenPlant(ID int) Plant {
	plant :=  Plant{
		len(main.AllGens),
		ID,
		3,
		make([]float64, 0),
		0,
	}

	return  plant
}
