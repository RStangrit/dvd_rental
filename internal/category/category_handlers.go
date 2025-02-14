package category

import (
	"main/pkg/db"
	"main/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostCategoryHandler(context *gin.Context) {
	var newCategory Category
	var err error

	if err = context.ShouldBindJSON(&newCategory); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = ValidateCategory(&newCategory); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = CreateCategory(db.GORM, &newCategory); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": newCategory})
}

func GetCategoriesHandler(context *gin.Context) {
	var pagination db.Pagination
	var err error

	if err = context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination parameters"})
		return
	}

	categories, totalRecords, err := ReadAllCategories(db.GORM, pagination)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": categories, "page": pagination.Page, "limit": pagination.Limit, "total": totalRecords})
}

func GetCategoryHandler(context *gin.Context) {
	categoryId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID format"})
		return
	}

	category, err := ReadOneCategory(db.GORM, categoryId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": category})
}

func PutCategoryHandler(context *gin.Context) {
	categoryId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID format"})
		return
	}

	category, err := ReadOneCategory(db.GORM, categoryId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updatedCategory Category
	err = context.ShouldBindJSON(&updatedCategory)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category data format"})
		return
	}

	if err = ValidateCategory(&updatedCategory); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedCategory.CategoryID = int(category.CategoryID)

	err = UpdateOneCategory(db.GORM, updatedCategory)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": updatedCategory})
}

func DeleteCategoryHandler(context *gin.Context) {
	categoryId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID format"})
		return
	}

	category, err := ReadOneCategory(db.GORM, categoryId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = DeleteOneCategory(db.GORM, *category)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"deleted": category})
}
