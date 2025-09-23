module calculator-app

go 1.23

require (
	calculator-math v0.0.0
	calculator-utils v0.0.0
)

replace calculator-math => ../math
replace calculator-utils => ../utils
