package main

import (
	"fmt"
	"log"

	"github.com/mtslzr/pokeapi-go"
	"github.com/xuri/excelize/v2"
)

type pokemon struct {
	Name string
	Hp, Attack, Defense, SpecialAttack, SpecialDefense, Speed int
}

var allpokemons []pokemon

func main(){

	results , err := pokeapi.Resource("pokemon", 0, 2000)

	if err != nil {
		log.Println(err)
	}

	var pokemonToCreate pokemon

	for _, r := range results.Results {
  

	pokemonInfo, err :=	pokeapi.Pokemon(r.Name)

		if err != nil {
			log.Println(err)
		}

	pokemonToCreate.Name = r.Name

	for _, stat := range pokemonInfo.Stats {
		
		switch stat.Stat.Name {
		case "hp":
			pokemonToCreate.Hp = stat.BaseStat

		case "attack":
			pokemonToCreate.Attack = stat.BaseStat

		case "defense":
			pokemonToCreate.Defense = stat.BaseStat

		case "special-attack":
			pokemonToCreate.SpecialAttack = stat.BaseStat

		case "special-defense":
			pokemonToCreate.SpecialDefense = stat.BaseStat

		case "speed":
			pokemonToCreate.Speed = stat.BaseStat
			
		}

	}

	allpokemons = append(allpokemons, pokemonToCreate)

	}


	f := excelize.NewFile()
	
	sheet := "sheet1"

	for pokemonIndex, pokemon := range allpokemons {

		for i := 0; i < 8; i++{

			switch i {

			case 0: f.SetCellValue(sheet, fmt.Sprintf("%v%v", excelColumnsForStats[i], pokemonIndex+2) , pokemonIndex+1)

			case 1: f.SetCellValue(sheet, fmt.Sprintf("%v%v", excelColumnsForStats[i], pokemonIndex+2) , pokemon.Name)

			case 2: f.SetCellValue(sheet, fmt.Sprintf("%v%v", excelColumnsForStats[i], pokemonIndex+2) , pokemon.Hp)

			case 3: f.SetCellValue(sheet, fmt.Sprintf("%v%v", excelColumnsForStats[i], pokemonIndex+2) , pokemon.Attack)

			case 4: f.SetCellValue(sheet, fmt.Sprintf("%v%v", excelColumnsForStats[i], pokemonIndex+2) , pokemon.Defense)

			case 5: f.SetCellValue(sheet, fmt.Sprintf("%v%v", excelColumnsForStats[i], pokemonIndex+2) , pokemon.SpecialAttack)

			case 6: f.SetCellValue(sheet, fmt.Sprintf("%v%v", excelColumnsForStats[i], pokemonIndex+2) , pokemon.SpecialDefense)

			case 7: f.SetCellValue(sheet, fmt.Sprintf("%v%v", excelColumnsForStats[i], pokemonIndex+2) , pokemon.Speed)
				
			}

		}
	
	}

	f.SaveAs("AllPokemonsWithStats.xlsx")

}

var excelColumnsForStats []string = []string{"A","B","C","D","E","F","G","H"}
var defaultColumnsNameForStats []string = []string{"ID","Name","HP","Attack","Defense","Special-Attack","Special-Defense","Speed"}