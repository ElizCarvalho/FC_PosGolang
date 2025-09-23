package tax

type Repository interface {
	SaveTax(amount float64) error
}

func CalculateTaxAndSave(amount float64, repository Repository) error {
	tax := CalculateTax(amount)
	return repository.SaveTax(tax)
}

func CalculateTax(amount float64) float64 {

	//pra conseguir ver o quanto de cobertura temos
	//usamos o comando $ go test -v -coverprofile=coverage.out (pra criar o arquivo de cobertura)
	//e depois usamos o comando $ go tool cover -html=coverage.out (pra ver a cobertura)
	if amount >= 20000 {
		return 20
	}

	if amount >= 1000 {
		return 10.0
	}

	if amount <= 0 {
		return 0
	}

	return 5.0
}
