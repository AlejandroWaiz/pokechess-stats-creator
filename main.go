package main

import (
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

	file := excelize.NewFile()

	log.Println(file.SetCellValue("sheet1", "A1", 3))

	log.Println(file.SaveAs("asjkdaksdjask.xlsx"))
	

}

func main2(){

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

	for _, pokemon := range allpokemons {
		
	}

}

var excelColumnsForStats []string = []string{"A","B","C","D","E","F","G","H"}
var defaultColumnsNameForStats []string = []string{"ID","Name","HP","Attack","Defense","Special-Attack","Special-Defense","Speed"}