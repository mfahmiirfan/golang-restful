package service

import (
	"context"
	"fmt"
	"mfahmii/golang-restful/app"
	"mfahmii/golang-restful/exception"
	"mfahmii/golang-restful/helper"
	"mfahmii/golang-restful/model/domain"
	"mfahmii/golang-restful/model/web"
	"mfahmii/golang-restful/repository"
	"reflect"

	"gorm.io/gorm"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *gorm.DB
	Validate           *app.Validation
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *gorm.DB, validate *app.Validation) CategoryService {
	fmt.Println(reflect.TypeOf(categoryRepository))
	fmt.Println(categoryRepository)
	fmt.Printf("%p\n", categoryRepository)
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	fmt.Println(reflect.TypeOf(request))
	fmt.Println(request)
	fmt.Printf("%p\n", &request)
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	helper.PanicIfError(tx.Error)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	helper.PanicIfError(tx.Error)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	category.Name = request.Name

	category = service.CategoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx := service.DB.Begin()
	helper.PanicIfError(tx.Error)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.CategoryRepository.Delete(ctx, tx, category)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx := service.DB.Begin()
	helper.PanicIfError(tx.Error)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx := service.DB.Begin()
	helper.PanicIfError(tx.Error)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	return helper.ToCategoryResponses(categories)
}
