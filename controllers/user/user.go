package usercontroller

import (
	"net/http"
	"os"
	"time"

	"example.com/gin-backend-api/configs"
	"example.com/gin-backend-api/models"
	"example.com/gin-backend-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/matthewhartstonge/argon2"
)

func GetAll(c *gin.Context) {
	var users []models.User
	// configs.DB.Order("id DESC").Find(&users)

	//SQL
	// configs.DB.Raw("select * from users order by id desc").Scan(&users)

	configs.DB.Preload("Blogs").Find(&users)

	c.JSON(200, gin.H{
		"data": users,
	})
}

func Register(c *gin.Context) {
	var input InputRegister
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Fullname: input.Fullname,
		Email:    input.Email,
		Password: input.Password,
	}

	//เช็คอีเมล์ซ้ำ
	userExist := configs.DB.Where("email = ?", input.Email).First(&user)
	if userExist.RowsAffected == 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	result := configs.DB.Debug().Create(&user)

	//db error
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Register successfully",
	})
}

func Login(c *gin.Context) {

	var input InputLogin
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Email:    input.Email,
		Password: input.Password,
	}

	//เช็คว่ามีผู้ใช้นี้ในระบบเราหรือไม่
	userAccount := configs.DB.Where("email = ?", input.Email).First(&user)
	if userAccount.RowsAffected < 1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	//เปรียบเทียบรหัสผ่านว่าที่ส่งมา กับในตาราง (เข้ารหัส) ตรงกันหรือไม่
	ok, _ := argon2.VerifyEncoded([]byte(input.Password), []byte(user.Password))
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
		return
	}

	//สร้าง token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24 * 2).Unix(),
	})

	jwtSecret := os.Getenv("JWT_SECRET")
	token, _ := claims.SignedString([]byte(jwtSecret))

	c.JSON(http.StatusCreated, gin.H{
		"message":      "Login success",
		"access_token": token,
	})

}

func GetById(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	result := configs.DB.First(&user, id)

	if result.RowsAffected < 1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, gin.H{
		"data": user,
	})
}

func SearchByFullname(c *gin.Context) {
	fullanme := c.Query("fullname") //?fullname=John

	var users []models.User
	result := configs.DB.Where("fullname LIKE ?", "%"+fullanme+"%").Scopes(utils.Paginate(c)).Find(&users)
	if result.RowsAffected < 1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, gin.H{
		"data": users,
	})
}

func GetProfile(c *gin.Context) {
	user := c.MustGet("user")
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func Update(c *gin.Context) {
	id := c.Param("id")

	var input InputUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	result := configs.DB.First(&user, id)

	if result.RowsAffected < 1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	user.Fullname = input.Fullname
	resultUpdate := configs.DB.Save(&user)

	//db error
	if resultUpdate.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": resultUpdate.Error})
		return
	}

	c.JSON(200, gin.H{
		"message": "Successfully Updated",
	})

}

func DeleteById(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	result := configs.DB.First(&user, id)

	if result.RowsAffected < 1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	resultDelete := configs.DB.Delete(&user)

	//db error
	if resultDelete.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": resultDelete.Error}) // id นี้ลบไม่ได้ อาจใช้กับตาราง blogs อยู่
		return
	}

	c.JSON(200, gin.H{
		"message": "Successfully deleted",
	})
}
