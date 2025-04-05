package route

import (
	"net/http"
	"server/models"
	"server/utils"

	"github.com/gin-gonic/gin"
)

func signup(cxt *gin.Context) {
	var user models.User
	err := cxt.ShouldBindJSON(&user)
	if err != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse input data"})
		return
	}

	err = user.Save()
	if err != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user"})
		return
	}
	cxt.JSON(http.StatusCreated, gin.H{"message": "User created!", "user": user})
}

func login(cxt *gin.Context) {
	var user models.User
	err := cxt.ShouldBindJSON(&user)
	if err != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse input data"})
		return
	}
	err = user.ValidateCredentials()
	if err != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user"})
		return
	}

	cxt.JSON(http.StatusCreated, gin.H{"message": "User login success!", "token": token})
}
