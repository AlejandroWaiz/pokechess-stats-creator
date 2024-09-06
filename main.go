package main

import (
	"fmt"
	"log"
	"math"

	"github.com/mtslzr/pokeapi-go"
	"github.com/xuri/excelize/v2"
)

type pokemon struct {
	Name string
	Hp, Attack, Defense, SpecialAttack, SpecialDefense, Speed float64
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
			pokemonToCreate.Hp = float64(stat.BaseStat)

		case "attack":
			pokemonToCreate.Attack = float64(stat.BaseStat)

		case "defense":
			pokemonToCreate.Defense = float64(stat.BaseStat)

		case "special-attack":
			pokemonToCreate.SpecialAttack = float64(stat.BaseStat)

		case "special-defense":
			pokemonToCreate.SpecialDefense = float64(stat.BaseStat)

		case "speed":
			pokemonToCreate.Speed = float64(stat.BaseStat)
			
		}

	}

	allpokemons = append(allpokemons, pokemonToCreate)

	}


	f := excelize.NewFile()
	
	sheet := "sheet1"

	var createHpStat = func(hp float64)(finalHp float64){

		finalHp = ((((2 * hp) + 94) * 100) / 100) + 110
		return math.Round(finalHp/5)

	}

	var createNonHpStat = func(nonHp float64) (finalNonHp float64){

		finalNonHp = ((((2*nonHp)+94)*100)/100) + 5
		return math.Round(finalNonHp/10)

	}

	for pokemonIndex, pokemon := range allpokemons {

		for i := 0; i < 8; i++{

			switch i {

			case 0: f.SetCellValue(sheet, fmt.Sprintf("%v%v", excelColumnsForStats[i], pokemonIndex+2) , pokemonIndex+1)

			case 1: f.SetCellValue(sheet, fmt.Sprintf("%v%v", excelColumnsForStats[i], pokemonIndex+2) , pokemon.Name)

			case 2: f.SetCellValue(sheet, fmt.Sprintf("%v%v", excelColumnsForStats[i], pokemonIndex+2) , createHpStat(pokemon.Hp))

			case 3: f.SetCellValue(sheet, fmt.Sprintf("%v%v", excelColumnsForStats[i], pokemonIndex+2) , createNonHpStat(pokemon.Attack))

			case 4: f.SetCellValue(sheet, fmt.Sprintf("%v%v", excelColumnsForStats[i], pokemonIndex+2) , createNonHpStat(pokemon.Defense))

			case 5: f.SetCellValue(sheet, fmt.Sprintf("%v%v", excelColumnsForStats[i], pokemonIndex+2) , createNonHpStat(pokemon.SpecialAttack))

			case 6: f.SetCellValue(sheet, fmt.Sprintf("%v%v", excelColumnsForStats[i], pokemonIndex+2) , createNonHpStat(pokemon.SpecialDefense))

			case 7: f.SetCellValue(sheet, fmt.Sprintf("%v%v", excelColumnsForStats[i], pokemonIndex+2) , createNonHpStat(pokemon.Speed))
				
			}

		}
	
	}

	f.SaveAs("AllPokemonsWithStats.xlsx")

}

var excelColumnsForStats []string = []string{"A","B","C","D","E","F","G","H"}
var defaultColumnsNameForStats []string = []string{"ID","Name","HP","Attack","Defense","Special-Attack","Special-Defense","Speed"}