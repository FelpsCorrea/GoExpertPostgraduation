package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/FelpsCorrea/GoExpertPostgraduation/SQLC/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

type CourseDB struct {
	dbConn *sql.DB
	*db.Queries
}

func NewCourseDB(dbConn *sql.DB) *CourseDB {
	return &CourseDB{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

type CourseParams struct {
	ID          string
	Name        string
	Description sql.NullString
	Price       float64
}

type CategoryParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

func (c *CourseDB) callTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := c.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := db.New(tx)

	err = fn(q)

	if err != nil {

		if errRb := tx.Rollback(); errRb != nil {
			return fmt.Errorf("tx error: %v, rb error: %v", err, errRb)
		}

		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (c *CourseDB) CreateCourseAndCategory(ctx context.Context, argsCategory CategoryParams, argsCourse CourseParams) error {

	err := c.callTx(ctx, func(q *db.Queries) error {
		err := q.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          argsCategory.ID,
			Name:        argsCategory.Name,
			Description: argsCategory.Description,
		})

		if err != nil {
			return err
		}

		err = q.CreateCourse(ctx, db.CreateCourseParams{
			ID:          argsCourse.ID,
			Name:        argsCourse.Name,
			Description: argsCourse.Description,
			CategoryID:  argsCategory.ID,
			Price:       argsCourse.Price,
		})

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3307)/courses")

	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	courses, err := queries.ListCourses(ctx)
	if err != nil {
		panic(err)
	}

	for _, course := range courses {
		fmt.Printf("Category: %s, Course ID: %s, Course Name: %s, Course Description: %s,Course Price: %f\n",
			course.CategoryName, course.ID, course.Name, course.Description.String, course.Price)
	}

	// courseArgs := CourseParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "Go Programming",
	// 	Description: sql.NullString{String: "Go Programming Description", Valid: true},
	// 	Price:       100.00,
	// }

	// categoryArgs := CategoryParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "Programming Category",
	// 	Description: sql.NullString{String: "Programming Category Description", Valid: true},
	// }

	// courseDB := NewCourseDB(dbConn)

	// err = courseDB.CreateCourseAndCategory(ctx, categoryArgs, courseArgs)

	// if err != nil {
	// 	panic(err)
	// }

	// err = queries.CreateCategory(ctx, db.CreateCategoryParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "Programming Category",
	// 	Description: sql.NullString{String: "Programming Category Description", Valid: true},
	// })

	// if err != nil {
	// 	panic(err)
	// }

	// categories, err := queries.ListCategories(ctx)

	// if err != nil {
	// 	panic(err)
	// }

	// for _, category := range categories {
	// 	println(category.Name)
	// 	println(category.ID)
	// }

	// err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
	// 	ID:          "6348b845-1caf-4e2d-a00c-6435d4bc0b79",
	// 	Name:        "Programming Category Updated",
	// 	Description: sql.NullString{String: "Programming Category Description Updated", Valid: true},
	// })

}
