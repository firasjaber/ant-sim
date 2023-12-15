package scene

import (
	"github.com/firasjaber/ant-sim/config"
	"github.com/firasjaber/ant-sim/entity"
)

type Entities struct {
	Ants []*entity.Ant
	Food []*entity.Food
	Home *entity.Home
}

func GetEntitiesByMapId(mapId int) *Entities {
	switch mapId {
	case 1:
		return prepareEntitiesOneNest()
	case 2:
		return prepareEntitiesTwoNests()
	case 3:
		return prepareEntitiesFourNests()
	case 4:
		return prepareEntitiesSurrounding()
	default:
		return prepareEntitiesOneNest()
	}
}

func prepareEntitiesOneNest() *Entities {
	// span entites
	// spawn home
	homeXPos := config.WindowWidth/2 - (config.WindowWidth / 4)
	homeYPos := config.WindowHeight / 2
	// create the home
	home := entity.NewHome(int32(homeXPos), int32(homeYPos))

	// spawn ants
	// loop through the ants number range and create new ants
	ants := []*entity.Ant{}
	for i := 0; i < config.AntsCount; i++ {
		// create a new ant
		ant := entity.NewAnt(int32(homeXPos), int32(homeYPos))
		// add the ant to the ants list
		ants = append(ants, ant)
	}

	// spawn food
	// loop through the food number range and create new food on random positions
	food := []*entity.Food{}

	initFoodSpawnXPos := config.WindowWidth/2 + config.WindowWidth/4
	initFoodSpawnYPos := config.WindowHeight / 4
	lastFoodSpawnXPos := initFoodSpawnXPos + config.FoodCount
	lastFoodSpawnYPos := initFoodSpawnYPos + config.FoodCount
	for i := initFoodSpawnXPos; i < lastFoodSpawnXPos; i++ {
		for j := initFoodSpawnYPos; j < lastFoodSpawnYPos; j++ {
			// create a new food
			f := entity.NewFood(int32(i), int32(j))
			// add the food to the food list
			food = append(food, f)
		}
	}

	return &Entities{Ants: ants, Food: food, Home: home}
}

func prepareEntitiesTwoNests() *Entities {
	// span entites
	// spawn home
	homeXPos := config.WindowWidth/2 - (config.WindowWidth / 4)
	homeYPos := config.WindowHeight / 2
	// create the home
	home := entity.NewHome(int32(homeXPos), int32(homeYPos))

	// spawn ants
	// loop through the ants number range and create new ants
	ants := []*entity.Ant{}
	for i := 0; i < config.AntsCount; i++ {
		// create a new ant
		ant := entity.NewAnt(int32(homeXPos), int32(homeYPos))
		// add the ant to the ants list
		ants = append(ants, ant)
	}

	// spawn food
	// loop through the food number range and create new food on random positions
	food := []*entity.Food{}

	initFoodSpawnXPos := config.WindowWidth/2 + config.WindowWidth/4
	initFoodSpawnYPos := config.WindowHeight / 4
	lastFoodSpawnXPos := initFoodSpawnXPos + config.FoodCount
	lastFoodSpawnYPos := initFoodSpawnYPos + config.FoodCount
	for i := initFoodSpawnXPos; i < lastFoodSpawnXPos; i++ {
		for j := initFoodSpawnYPos; j < lastFoodSpawnYPos; j++ {
			// create a new food
			f := entity.NewFood(int32(i), int32(j))
			// add the food to the food list
			food = append(food, f)
		}
	}

	initFoodSpawnTwoXPos := (config.WindowWidth / 2) + config.WindowWidth/4
	initFoodSpawnTwoYPos := (config.WindowHeight / 4) + config.WindowHeight/2
	lastFoodSpawnTwoXPos := initFoodSpawnTwoXPos + config.FoodCount
	lastFoodSpawnTwoYPos := initFoodSpawnTwoYPos + config.FoodCount
	for i := initFoodSpawnTwoXPos; i < lastFoodSpawnTwoXPos; i++ {
		for j := initFoodSpawnTwoYPos; j < lastFoodSpawnTwoYPos; j++ {
			// create a new food
			f := entity.NewFood(int32(i), int32(j))
			// add the food to the food list
			food = append(food, f)
		}
	}

	return &Entities{Ants: ants, Food: food, Home: home}
}

func prepareEntitiesFourNests() *Entities {
	// span entites
	// spawn home
	homeXPos := config.WindowWidth/2 - 15
	homeYPos := config.WindowHeight/2 - 15
	// create the home
	home := entity.NewHome(int32(homeXPos), int32(homeYPos))

	// spawn ants
	// loop through the ants number range and create new ants
	ants := []*entity.Ant{}
	for i := 0; i < config.AntsCount; i++ {
		// create a new ant
		ant := entity.NewAnt(int32(homeXPos), int32(homeYPos))
		// add the ant to the ants list
		ants = append(ants, ant)
	}

	// spawn food
	// loop through the food number range and create new food on random positions
	food := []*entity.Food{}

	initFoodSpawnXPos := 0
	initFoodSpawnYPos := 0
	lastFoodSpawnXPos := initFoodSpawnXPos + config.FoodCount
	lastFoodSpawnYPos := initFoodSpawnYPos + config.FoodCount
	for i := initFoodSpawnXPos; i < lastFoodSpawnXPos; i++ {
		for j := initFoodSpawnYPos; j < lastFoodSpawnYPos; j++ {
			// create a new food
			f := entity.NewFood(int32(i), int32(j))
			// add the food to the food list
			food = append(food, f)
		}
	}

	initFoodSpawnTwoXPos := 0
	initFoodSpawnTwoYPos := config.WindowHeight - config.FoodCount
	lastFoodSpawnTwoXPos := config.FoodCount
	lastFoodSpawnTwoYPos := config.WindowHeight
	for i := initFoodSpawnTwoXPos; i < lastFoodSpawnTwoXPos; i++ {
		for j := initFoodSpawnTwoYPos; j < lastFoodSpawnTwoYPos; j++ {
			// create a new food
			f := entity.NewFood(int32(i), int32(j))
			// add the food to the food list
			food = append(food, f)
		}
	}

	initFoodSpawnThreeXPos := config.WindowWidth - config.FoodCount
	initFoodSpawnThreeYPos := 0
	lastFoodSpawnThreeXPos := config.WindowWidth
	lastFoodSpawnThreeYPos := config.FoodCount
	for i := initFoodSpawnThreeXPos; i < lastFoodSpawnThreeXPos; i++ {
		for j := initFoodSpawnThreeYPos; j < lastFoodSpawnThreeYPos; j++ {
			// create a new food
			f := entity.NewFood(int32(i), int32(j))
			// add the food to the food list
			food = append(food, f)
		}
	}

	initFoodSpawnFourXPos := config.WindowWidth - config.FoodCount
	initFoodSpawnFourYPos := config.WindowHeight - config.FoodCount
	lastFoodSpawnFourXPos := config.WindowWidth
	lastFoodSpawnFourYPos := config.WindowHeight
	for i := initFoodSpawnFourXPos; i < lastFoodSpawnFourXPos; i++ {
		for j := initFoodSpawnFourYPos; j < lastFoodSpawnFourYPos; j++ {
			// create a new food
			f := entity.NewFood(int32(i), int32(j))
			// add the food to the food list
			food = append(food, f)
		}
	}

	return &Entities{Ants: ants, Food: food, Home: home}
}

func prepareEntitiesSurrounding() *Entities {
	// spawn home
	homeXPos := config.WindowWidth/2 - 15
	homeYPos := config.WindowHeight/2 - 15
	// create the home
	home := entity.NewHome(int32(homeXPos), int32(homeYPos))

	// spawn ants
	// loop through the ants number range and create new ants
	ants := []*entity.Ant{}
	for i := 0; i < config.AntsCount; i++ {
		// create a new ant
		ant := entity.NewAnt(int32(homeXPos), int32(homeYPos))
		// add the ant to the ants list
		ants = append(ants, ant)
	}

	// spawn food
	// loop through the food number range and create new food on random positions
	food := []*entity.Food{}

	ax := 20
	ay := 20
	bx := config.WindowWidth - 20
	by := 22
	for i := ax; i < bx; i++ {
		for j := ay; j < by; j++ {
			// create a new food
			f := entity.NewFood(int32(i), int32(j))
			// add the food to the food list
			food = append(food, f)
		}
	}

	cx := 20
	cy := config.WindowHeight - 23
	dx := config.WindowWidth - 20
	dy := config.WindowHeight - 20
	for i := cx; i < dx; i++ {
		for j := cy; j < dy; j++ {
			// create a new food
			f := entity.NewFood(int32(i), int32(j))
			// add the food to the food list
			food = append(food, f)
		}
	}

	ex := 20
	ey := 23
	fx := 23
	fy := config.WindowHeight - 25
	for i := ex; i < fx; i++ {
		for j := ey; j < fy; j++ {
			// create a new food
			f := entity.NewFood(int32(i), int32(j))
			// add the food to the food list
			food = append(food, f)
		}
	}

	gx := config.WindowWidth - 23
	gy := 23
	hx := config.WindowWidth - 20
	hy := config.WindowHeight - 25
	for i := gx; i < hx; i++ {
		for j := gy; j < hy; j++ {
			// create a new food
			f := entity.NewFood(int32(i), int32(j))
			// add the food to the food list
			food = append(food, f)
		}
	}

	return &Entities{Ants: ants, Food: food, Home: home}
}
