/*
Copyright © 2025 ElizCarvalho

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"github.com/ElizCarvalho/FC_PosGolang/15_Cobra_CLI/cmd"
	"github.com/ElizCarvalho/FC_PosGolang/15_Cobra_CLI/internal/config"
	"github.com/ElizCarvalho/FC_PosGolang/15_Cobra_CLI/internal/database"
	"github.com/ElizCarvalho/FC_PosGolang/15_Cobra_CLI/internal/service"
)

func main() {
	// Inicializar dependências usando DI
	db := config.GetDB()
	defer db.Close()

	// Criar repositório
	categoryRepo := database.NewCategory(db)

	// Criar service
	categoryService := service.NewCategoryService(categoryRepo)

	// Injetar service no comando
	cmd.InitializeCategoryService(categoryService)

	// Executar CLI
	cmd.Execute()
}
