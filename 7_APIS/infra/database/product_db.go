package database

import (
	"github.com/ElizCarvalho/FC_PosGolang/7_APIS/internal/entity"
	"gorm.io/gorm"
)

type ProductDB struct {
	db *gorm.DB
}

func NewProductDB(db *gorm.DB) *ProductDB {
	return &ProductDB{db: db}
}

func (p *ProductDB) Create(product *entity.Product) error {
	return p.db.Create(product).Error
}

func (p *ProductDB) FindById(id string) (*entity.Product, error) {
	var product entity.Product
	err := p.db.First(&product, "id = ?", id).Error
	return &product, err
}

func (p *ProductDB) Update(product *entity.Product) error {
	_, err := p.FindById(product.ID.String())
	if err != nil {
		return err
	}
	return p.db.Save(product).Error
}

func (p *ProductDB) Delete(id string) error {
	_, err := p.FindById(id)
	if err != nil {
		return err
	}
	return p.db.Delete(&entity.Product{}, "id = ?", id).Error
}

func (p *ProductDB) FindAll(page, limit int, sort string) ([]entity.Product, error) {
	if sort != "" && sort != "desc" && sort != "asc" {
		sort = "asc"
	}

	var products []entity.Product
	var err error
	if page != 0 && limit != 0 {
		err = p.db.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&products).Error
		return products, err
	} else {
		err = p.db.Order("created_at " + sort).Find(&products).Error
		return products, err
	}
}
