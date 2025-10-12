package service

import (
	"github.com/ElizCarvalho/FC_PosGolang/15_Cobra_CLI/internal/database"
)

// CategoryService interface para injeção de dependência
type CategoryService interface {
	Create(name, description string) (database.Category, error)
	List() ([]database.Category, error)
	GetByID(id string) (database.Category, error)
	Update(id, name, description string) error
	Delete(id string) error
}

// categoryServiceImpl implementação do serviço de categoria
type categoryServiceImpl struct {
	repo database.CategoryRepository
}

// NewCategoryService cria uma nova instância do serviço de categoria
func NewCategoryService(repo database.CategoryRepository) CategoryService {
	return &categoryServiceImpl{
		repo: repo,
	}
}

// Create implementa a criação de categoria
func (s *categoryServiceImpl) Create(name, description string) (database.Category, error) {
	return s.repo.Create(name, description)
}

// List implementa a listagem de categorias
func (s *categoryServiceImpl) List() ([]database.Category, error) {
	return s.repo.List()
}

// GetByID implementa a busca de categoria por ID
func (s *categoryServiceImpl) GetByID(id string) (database.Category, error) {
	return s.repo.GetByID(id)
}

// Update implementa a atualização de categoria
func (s *categoryServiceImpl) Update(id, name, description string) error {
	return s.repo.Update(id, name, description)
}

// Delete implementa a deleção de categoria
func (s *categoryServiceImpl) Delete(id string) error {
	return s.repo.Delete(id)
}
