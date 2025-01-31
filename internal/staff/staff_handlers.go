package staff

import (
	"main/pkg/db"
	"main/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostStaffHandler(context *gin.Context) {
	var newStaff Staff
	var err error

	if err = context.ShouldBindJSON(&newStaff); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = ValidateStaff(&newStaff); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = CreateStaff(&newStaff); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": newStaff})
}

func GetStaffsHandler(context *gin.Context) {
	var pagination db.Pagination
	var err error

	if err = context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination parameters"})
		return
	}

	staffs, totalRecords, err := ReadAllStaff(pagination)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": staffs, "page": pagination.Page, "limit": pagination.Limit, "total": totalRecords})
}

func GetStaffHandler(context *gin.Context) {
	staffId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid staff ID format"})
		return
	}

	staff, err := ReadOneStaff(staffId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": staff})
}

func PutStaffHandler(context *gin.Context) {
	staffId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid staff ID format"})
		return
	}

	staff, err := ReadOneStaff(staffId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updatedStaff Staff
	err = context.ShouldBindJSON(&updatedStaff)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid staff data format"})
		return
	}

	updatedStaff.StaffID = int(staff.StaffID)

	err = UpdateOneStaff(updatedStaff)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update staff"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": updatedStaff})
}

func DeleteStaffHandler(context *gin.Context) {
	staffId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid staff ID format"})
		return
	}

	staff, err := ReadOneStaff(staffId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Staff not found"})
		return
	}

	err = DeleteOneStaff(*staff)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete staff"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"deleted": staff})
}
