package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Flight struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Flight{}) //em ambiente de desenvolvimento

	//criar um novo voo
	//db.Create(&Flight{Name: "BSB-GIG", Price: 350})

	//criar em batch
	/*flighs := []Flight{
		{Name: "ABC-123", Price: 100},
		{Name: "DEF-456", Price: 200},
		{Name: "GHI-789", Price: 300},
	}
	db.Create(&flighs)*/

	//selecionar um voo
	var flight Flight
	fmt.Println("Selecionando voo com ID 2")
	db.Debug().First(&flight, 2)
	db.First(&flight, "id = ?", 2)
	fmt.Println(flight)

	//selecionar todos os voos
	var flights []Flight
	fmt.Println("Selecionando todos os voos")
	db.Debug().Find(&flights)
	fmt.Println(flights)

	//busca com limit
	fmt.Println("Selecionando 2 voos")
	db.Debug().Limit(2).Find(&flights)
	fmt.Println(flights)

	//busca com offset (paginação)
	fmt.Println("Selecionando 2 voos a partir do 2")
	db.Debug().Offset(2).Limit(2).Find(&flights)
	fmt.Println(flights)

	//busca com where
	fmt.Println("Selecionando voos com preço maior que 100")
	db.Debug().Where("price > ?", 100).Find(&flights)
	fmt.Println(flights)

	//busca com like
	fmt.Println("Selecionando voos com nome começando com B")
	db.Debug().Where("name LIKE ?", "B%").Find(&flights)
	fmt.Println(flights)

	//busca com where e AND
	fmt.Println("Selecionando voos com preço maior que 100 e menor que 300")
	db.Debug().Where("price > ? AND price < ?", 100, 300).Find(&flights)
	fmt.Println(flights)

	//busca com where e OR
	fmt.Println("Selecionando voos com preço maior que 100 ou menor que 300")
	db.Debug().Where("price > ? OR price < ?", 100, 300).Find(&flights)
	fmt.Println(flights)

	//atualizar um voo
	fmt.Println("Atualizando voo com ID 1")
	db.Debug().Model(&Flight{}).Where("id = ?", 1).Update("price", 150)
	fmt.Println("Voo atualizado com sucesso")

	//deletar um voo
	fmt.Println("Deletando voo com ID 2")
	db.Debug().Delete(&Flight{}, 2)
	fmt.Println("Voo deletado com sucesso")

}
