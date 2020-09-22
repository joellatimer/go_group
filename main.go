package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	// "log"
	"net/http"
	"os"
	"time"
)
// Group struct
type Group struct {
	ID uint64 `json:"id"`
	Groupname string `json:"groupname"`
	Password string `json:"password"`
}
var group = Group{
	ID: 1,
	Groupname:"East Macon",
	Password: "123456",
}
//Login func
func Login(c *gin.Context){
	var g Group
	if err := c.ShouldBindJSON(&g); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid json provided")
		return
	}
	//compare the group from the request, with the one we defined
	if group.Groupname != g.Groupname || group.Password != g.Password {
		c.JSON(http.StatusUnauthorized, "Please provide a vlid login")
		return
	}
	token, err := createToken(group.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, token)
}
func createToken(groupid uint64)(string, error){
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET","mybigsecret")
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["group_id"] = groupid
	atClaims["exp"]= time.Now().Add(time.Minute * 60).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

var (
	router=gin.Default()
)

func main() {
	router.POST("/login", Login)
	router.Run(":8080")
}