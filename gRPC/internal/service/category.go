package service

import (
	"context"
	"io"

	"github.com/FelpsCorrea/GoExpertPostgraduation/gRPC/internal/database"
	"github.com/FelpsCorrea/GoExpertPostgraduation/gRPC/internal/pb"
)

// CategoryService represents the service for managing categories.
type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

// NewCategoryService creates a new instance of CategoryService.
func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

// CreateCategory creates a new category.
func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}

	return &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, nil
}

func (c *CategoryService) ListCategories(ctx context.Context, in *pb.Blank) (*pb.CategoryList, error) {
	categories, err := c.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}

	var pbCategories []*pb.Category
	for _, category := range categories {
		pbCategories = append(pbCategories, &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}

	return &pb.CategoryList{Categories: pbCategories}, nil
}

func (c *CategoryService) GetCategory(ctx context.Context, in *pb.GetCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.FindByID(in.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, nil
}

// CreateCategoryStream is a streaming RPC method that creates multiple categories.
func (c *CategoryService) CreateCategoryStream(stream pb.CategoryService_CreateCategoryStreamServer) error {

	// Create an empty CategoryList to store the created categories.
	categories := &pb.CategoryList{}

	// Loop until the client finishes sending categories.
	for {

		// Receive a category from the client.
		category, err := stream.Recv()

		// If the client finished sending categories, send the CategoryList and close the stream.
		if err == io.EOF {
			return stream.SendAndClose(categories)
		}

		// If there was an error receiving the category, return the error.
		if err != nil {
			return err
		}

		// Create the category in the database.
		categoryResult, err := c.CategoryDB.Create(category.Name, category.Description)

		if err != nil {
			return err
		}

		// Append the created category to the CategoryList.
		categories.Categories = append(categories.Categories, &pb.Category{
			Id:          categoryResult.ID,
			Name:        categoryResult.Name,
			Description: categoryResult.Description,
		})

	}
}

// CreateCategoryStream is a bidirectional streaming RPC method that creates multiple categories.
func (c *CategoryService) CreateCategoryBidirectionalStream(stream pb.CategoryService_CreateCategoryBidirectionalStreamServer) error {

	// Loop until the client finishes sending categories.
	for {
		// Receive a category from the client.
		category, err := stream.Recv()

		// If the client finished sending categories, close the stream.
		if err == io.EOF {
			return nil
		}

		// If there was an error receiving the category, return the error.
		if err != nil {
			return err
		}

		// Create the category in the database.
		categoryResult, err := c.CategoryDB.Create(category.Name, category.Description)

		if err != nil {
			return err
		}

		// Send the created category back to the client.
		err = stream.Send(&pb.Category{
			Id:          categoryResult.ID,
			Name:        categoryResult.Name,
			Description: categoryResult.Description,
		})

		if err != nil {
			return err
		}
	}
}
