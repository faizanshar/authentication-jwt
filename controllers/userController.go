package controllers

import (
	"cobacoba1/initializers"
	"cobacoba1/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"os"
	"time"
)

var MyKey = []byte(os.Getenv("SECRET"))

func CreateUser(c *gin.Context) {
	//get data from body

	var body struct {
		Name     string
		Username string
		Password string
		Status   int
	}

	c.Bind(&body)

	if body.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Name is required",
		})
		return
	} else if body.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Username is required",
		})
		return
	} else if body.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Password is required",
		})
		return
	} else if body.Status == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Status is required",
		})
		return
	}

	// hashpassword
	hash, er := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to has password!",
		})
		return
	}

	// create post
	post := models.User{Name: body.Name, Username: body.Username, Password: string(hash), Status: body.Status}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   post,
	})
}

func LoginUser(c *gin.Context) {
	var body struct {
		Username string
		Password string
	}

	c.Bind(&body)
	if body.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Username is required",
		})
		return
	} else if body.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Password is required",
		})
		return
	}

	var user models.User

	initializers.DB.First(&user, "username = ?", body.Username)
	if user.Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Email or Password!",
		})
		return
	}

	//compare password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Email or Password!",
		})
		return
	}

	//generate token
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"sub":    user.Id,
		"status": user.Status,
		"exp":    time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	//sign and get complete encode token
	tokenString, errr := token.SigningString()

	if errr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid to create token!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func GetAllUser(c *gin.Context) {
	var user []models.User
	initializers.DB.Find(&user)
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   user,
	})
}

func GetUserById(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	if err := initializers.DB.First(&user, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "Data tidak ditemukan",
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   user,
	})
}
