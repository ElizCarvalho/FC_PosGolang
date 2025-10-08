module calculator-app

go 1.23

require (
	calculator-math v0.0.0
	calculator-utils v0.0.0
)

require github.com/google/uuid v1.6.0 // indirect

replace calculator-math => ../math

replace calculator-utils => ../utils
