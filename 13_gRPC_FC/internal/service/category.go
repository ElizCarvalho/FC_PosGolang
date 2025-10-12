package service

import (
	"context"

	"github.com/ElizCarvalho/FC_PosGolang/13_gRPC_FC/internal/database"
	"github.com/ElizCarvalho/FC_PosGolang/13_gRPC_FC/internal/pb"
	"google.golang.org/grpc"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{CategoryDB: categoryDB}
}

func (c *CategoryService) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDB.Create(req.Name, req.Description)
	if err != nil {
		return nil, err
	}

	categoryPB := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return &pb.CategoryResponse{
		Category: categoryPB,
	}, nil
}

func (c *CategoryService) ListCategories(ctx context.Context, req *pb.Blank) (*pb.CategoryList, error) {
	categories, err := c.CategoryDB.List()
	if err != nil {
		return nil, err
	}

	categoryPB := make([]*pb.Category, 0, len(categories))
	for _, category := range categories {
		categoryPB = append(categoryPB, &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}

	return &pb.CategoryList{Categories: categoryPB}, nil
}

func (c *CategoryService) GetCategory(ctx context.Context, req *pb.GetCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDB.GetByID(req.Id)
	if err != nil {
		return nil, err
	}

	categoryPB := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return &pb.CategoryResponse{
		Category: categoryPB,
	}, nil
}

// CreateCategoryStream implementa server-side streaming
// Cria múltiplas categorias e envia em lotes via stream
func (c *CategoryService) CreateCategoryStream(req *pb.CreateCategoryRequest, stream grpc.ServerStreamingServer[pb.CategoryList]) error {
	// Simula criação de 10 categorias para demonstração
	categories, err := c.CategoryDB.CreateMultiple(req.Name, req.Description, 10)
	if err != nil {
		return err
	}

	// Processa em lotes de 3 categorias por vez
	batchSize := 3
	for i := 0; i < len(categories); i += batchSize {
		// Determina o fim do lote atual
		end := i + batchSize
		if end > len(categories) {
			end = len(categories)
		}

		// Cria o lote atual
		var batch []*pb.Category
		for j := i; j < end; j++ {
			batch = append(batch, &pb.Category{
				Id:          categories[j].ID,
				Name:        categories[j].Name,
				Description: categories[j].Description,
			})
		}

		// Envia o lote via stream
		err := stream.Send(&pb.CategoryList{
			Categories: batch,
		})
		if err != nil {
			return err
		}

		// Simula processamento (opcional)
		// time.Sleep(100 * time.Millisecond)
	}

	return nil
}
