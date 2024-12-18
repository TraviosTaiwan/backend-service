package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/iwachan14736/travios-backend-service/pkg/domain"
	"github.com/iwachan14736/travios-backend-service/pkg/models"
	"github.com/iwachan14736/travios-backend-service/pkg/types"
	"github.com/labstack/echo/v4"
)

var CategoryService domain.ICategoryService

func SetCategoryService(cService domain.ICategoryService) {
	CategoryService = cService
}

// @Summary Create a new category
// @Description Create a new category with the provided details
// @Tags categories
// @Accept json
// @Produce json
// @Security Bearer
// @Param category body types.Category true "Category"
// @Success 201 {object} types.Category
// @Failure 400 {string} string "Invalid input"
// @Failure 401 {string} string "Unauthorized"
// @Router /category [post]
func CreateCategory(e echo.Context) error {
	reqCategory := &types.Category{}
	if err := e.Bind(reqCategory); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	if err := reqCategory.ValidateCategory(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	category := &models.Category{
		Name:        reqCategory.Name,
		Description: reqCategory.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if err := CategoryService.CreateCategory(category); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusCreated, category)
}

// @Summary Get categories
// @Description Get a list of categories or a specific category by ID
// @Tags categories
// @Produce json
// @Security Bearer
// @Param categoryID query int false "Category ID"
// @Success 200 {array} types.Category
// @Failure 401 {string} string "Unauthorized"
// @Router /category [get]
func GetCategories(e echo.Context) error {
	categoryID, _ := strconv.Atoi(e.QueryParam("categoryID"))
	categories, err := CategoryService.GetCategories(uint(categoryID))
	if err != nil {
		return e.JSON(http.StatusNotFound, err.Error())
	}
	return e.JSON(http.StatusOK, categories)
}

// @Summary Update a category
// @Description Update a category with the provided details
// @Tags categories
// @Accept json
// @Produce json
// @Security Bearer
// @Param categoryID path int true "Category ID"
// @Param category body types.Category true "Category"
// @Success 200 {object} types.Category
// @Failure 400 {string} string "Invalid input"
// @Failure 401 {string} string "Unauthorized"
// @Router /category/{categoryID} [put]
func UpdateCategory(e echo.Context) error {
	categoryID, _ := strconv.Atoi(e.Param("categoryID"))
	reqCategory := &types.Category{}
	if err := e.Bind(reqCategory); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	if err := reqCategory.ValidateCategory(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	category := &models.Category{
		ID:          uint(categoryID),
		Name:        reqCategory.Name,
		Description: reqCategory.Description,
	}
	if err := CategoryService.UpdateCategory(category); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, category)
}

// @Summary Delete a category
// @Description Delete a category by ID
// @Tags categories
// @Produce json
// @Security Bearer
// @Param categoryID path int true "Category ID"
// @Success 200 {string} string "Category deleted successfully"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Category not found"
// @Router /category/{categoryID} [delete]
func DeleteCategory(e echo.Context) error {
	categoryID, _ := strconv.Atoi(e.Param("categoryID"))
	if err := CategoryService.DeleteCategory(uint(categoryID)); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, "Category deleted successfully")
}
