package auth

import (
	"os"
	"time"
	"emad.com/config"
	"emad.com/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)


var body struct {
		 Email string
		 Password string
	}

func Login (c *gin.Context)  {

	

	var user = new(models.Users)

		c.Bind(&body)

		if body.Email == "" || body.Password == "" {
			c.JSON(400, gin.H{
				"messege": "Invalid data",
			})
			return 
		}

		 result := config.DB.Where("email = ?", body.Email).First(&user)

		if result.Error == gorm.ErrRecordNotFound {
				c.JSON(400, gin.H{
					"messege": "user not found",
				})
				return
		    }	
		

       if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		
		    	c.JSON(400, gin.H{
					"messege": "Incorrect email or password",
				})
				return
      }


	   token, err := generateJWT(int(user.ID))
		if err != nil {
			 c.JSON(400, gin.H{
					"messege": err,
				})
				return
		}
	

		checkAuth(int(user.ID),token)

         c.JSON(200, gin.H{
			"message": "loged in seccessfully",
			"user": user,
			"token":token,
		 })

		 return
}



 func generateJWT(id int) (string, error) {

	
     secretKey := os.Getenv("SECRET_KEY")


	// Create JWT token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
		"exp":      time.Now().Add(time.Hour * 720).Unix(),
	})

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, err
}



func checkAuth (id int, token string) error {

	 var auth = new(models.Auths)
	
   if err := config.DB.Where("user_id = ?", id).First(&auth).Error; err != nil {

		if err == gorm.ErrRecordNotFound {

			auth.User_id = id
			auth.Token = token

			if err := config.DB.Create(&auth).Error; err != nil {

		     return err
	        }
		}
	    	
		return err
	}

	 auth.User_id = id
	 auth.Token = token

    if err := config.DB.Save(&auth).Error; err!=nil {
		 return err
	}
	 
	return nil

}