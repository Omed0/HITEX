package controllers

import (
	
	"emad.com/config"
	"emad.com/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)


func Create (c *gin.Context)  {

	var body struct{

		Name  string
		Email string
		Role  string
		Password  string
	}


	c.Bind(&body)

	if body.Name == "" || body.Email == "" || body.Role == "" || body.Password == "" {

			 c.JSON(500, gin.H{
				"messege": "all data required",
			} )
		return
		}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

		if err != nil {
			
			 c.JSON(400, gin.H{
				"error": err,
			 })
			 return
		}
	

	post := models.Users{Name: body.Name, Email: body.Email, Role: body.Role, Password: string(hash)}

    result := config.DB.Create(&post)

	if result.Error !=nil {
		c.JSON(400, gin.H{
			"messege": result.Error,
		})
		 return
		 
	}

	c.JSON(200, gin.H{
		"messege": "user created seccessfully",
		"user": result,
	})

	return

}

func List (c *gin.Context) {
	    
		var user = new([]models.Users)

		result := config.DB.Find(&user) 
		
		if result.Error != nil {

		 c.JSON(400,gin.H{
			"message": result.Error,
		})

		return
	    }


		 c.JSON(200,gin.H{
			"message": "list of users",
			"users": user,
		})
		
	}


func One (c *gin.Context)  {
	    
		var user = new(models.Users)
	
		 id := c.Param("id")

		 result := config.DB.First(&user, id) 
		 if result.Error !=nil {
			if result.Error == gorm.ErrRecordNotFound {
				 c.JSON(400, gin.H{
				"message": "user not found",
		   })
			}
			return

		} 

		
		 c.JSON(200, gin.H{
			"message": "one user",
			"user": user,
			
		})
		return
		
	}


func Update (c *gin.Context)  {
	    
		var body struct{

		Name  string
		Email string
		Role  string
		Password  string
	}

	     user := new(models.Users)
	
		 id :=c.Param("id")

		  result1 := config.DB.First(&user, id)
			
		 if result1.Error == gorm.ErrRecordNotFound {
				 c.JSON(400, gin.H{
				"message": "user not found",
		     })
			 return
			}

			c.Bind(&body)

		if body.Name == "" ||  body.Email == "" ||  body.Role == "" ||  body.Password == "" {
			 c.JSON(400, gin.H{
				"message": "Invalid data",
		     })
			 return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

		if err != nil {
			 c.JSON(400, gin.H{
				"message": err,
		     })
			 return
		}
        
		user.Name = body.Name
		user.Email= body.Email
		user.Role = body.Role
		user.Password = string(hash)
		

	if    config.DB.Save(&user).RowsAffected == 0 {
		c.JSON(400, gin.H{
				"message": "update failed",
		     })
			 return
	}

			
		
		c.JSON(400, gin.H{
				"message": "user updated seccessfully",
				"user": user,
		     })
			 return
		
}


func Delete (c *gin.Context)  {
	    
		var user = new(models.Users)
	
		 id :=c.Param("id")

		 result := config.DB.First(&user, id)
			
		 if result.Error == gorm.ErrRecordNotFound {

				 c.JSON(400, gin.H{
				"message": "user not found",
		     })
			 return
			}


	 if  config.DB.Delete(&user).RowsAffected == 0 {
		 c.JSON(400, gin.H{
				"message": "delete failed",
		     })
			 return
	}

			
		
		 c.JSON(200, gin.H{
				"message": "user deleted successfully",
		     })
			 return
		
	}