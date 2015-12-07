package user_controllers

import (
	"fmt"
	"iot-go-api/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	UserController struct {
		session *mgo.Session
	}
)

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc UserController) GetUser(c *gin.Context) {
	name := c.Request.Header.Get("name")
	id := c.Request.Header.Get("id")
	gender := c.Request.Header.Get("gender")

	if !bson.IsObjectIdHex(id) {
		c.Writer.WriteHeader(404)
		return
	}
	oid := bson.ObjectIdHex(id)
	fmt.Println(oid)
	u := models.User{}

	fmt.Println(id)
	fmt.Println(name)
	//	if err := uc.session.DB("go_rest_tutorial").C("users").FindId(oid).One(&u); err != nil {
	if err := uc.session.DB("go_rest_tutorial").C("users").Find(bson.M{"gender": gender}).One(&u); err != nil {
		c.Writer.WriteHeader(404)
		return
	}
	c.JSON(200, u)
}



func (uc UserController) CreateUser(c *gin.Context) {
	u := models.User{}
	c.BindJSON(&u)

	token := jwt.New(jwt.SigningMethodHS256)
	const myToken = "test"
	// Set a header and a claim
	fmt.Println(u.Name)
	token.Header["none"] = u.Id
	token.Claims["Id"] = u.Id
	token.Claims["Name"] = u.Name
	token.Claims["Gender"] = u.Gender
	token.Claims["Age"] = u.Age
	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, _ := token.SignedString([]byte(myToken))

	dtoken, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		//		return myLookupKey(token.Header["kid"])
		return []byte(myToken), nil
	})

	if err == nil && dtoken.Valid {
		fmt.Println("ok")
		fmt.Println(dtoken)
	} else {
		fmt.Println("miss")
	}

	u.Id = bson.NewObjectId()
	u.Jwt = t
	uc.session.DB("go_rest_tutorial").C("users").Insert(u)
	c.JSON(201, u)

}

func (uc UserController) RemoveUser(c *gin.Context) {
	id := c.Param("id")
	if !bson.IsObjectIdHex(id) {
		c.Writer.WriteHeader(404)
		return
	}
	oid := bson.ObjectIdHex(id)
	if err := uc.session.DB("go_rest_tutorial").C("users").RemoveId(oid); err != nil {
		c.Writer.WriteHeader(404)
		return
	}
	c.Writer.WriteHeader(200)
}


func (uc UserController) GetMessageold(c *gin.Context) {
	u := models.Message{}
	if err := uc.session.DB("go_rest_tutorial").C("messages").Find(bson.M{"theme": "test"}).One(&u); err != nil {
		c.Writer.WriteHeader(404)
		return
	}
	fmt.Println(u)
	c.JSON(200, u)
}

func (uc UserController) CreateMessage(c *gin.Context) {
	u := models.Message{}
	c.BindJSON(&u)
	u.Id = bson.NewObjectId()
	uc.session.DB("go_rest_tutorial").C("messages").Insert(u)
	c.JSON(201, u)
}

// Binding from JSON
type Jwt struct {
	Jwt string `form:"jwt" json:"jwt" binding:"required"`
}

func (uc UserController) GetMessage(c *gin.Context) {
	const myToken = "test"

	var json Jwt
	c.BindJSON(&json)
	t := json.Jwt

	dtoken, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		//		return myLookupKey(token.Header["kid"])
		return []byte(myToken), nil
	})

	if err == nil && dtoken.Valid {
		fmt.Println("ok")
		fmt.Println(dtoken.Claims["room"])
	} else {
		fmt.Println("miss")
	}

	u := []models.Message{}
//	if err := uc.session.DB("chat_app").C("messages").Find(bson.M{"room": dtoken.Claims["room"]}).One(&u); err != nil {
//		c.Writer.WriteHeader(404)
//		return
//	}

	if err := uc.session.DB("chat_app").C("messages").Find(bson.M{"user": dtoken.Claims["user"]}).All(&u); err != nil {
		c.Writer.WriteHeader(404)
		return
	}

	c.JSON(200, u)
}


func (uc UserController) JwtCreateOwner(c *gin.Context) {
	u := models.User{}
	c.BindJSON(&u)

	token := jwt.New(jwt.SigningMethodHS256)
	const myToken = "test"
	// Set a header and a claim
	fmt.Println(u.Name)
	token.Header["none"] = u.Id
	token.Claims["Id"] = u.Id
	token.Claims["Name"] = u.Name
	token.Claims["Gender"] = u.Gender
	token.Claims["Age"] = u.Age
	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, _ := token.SignedString([]byte(myToken))
	c.JSON(201, t)
}




func (uc UserController) JwtCreateClient(c *gin.Context) {
	u := models.User{}
	c.BindJSON(&u)

	token := jwt.New(jwt.SigningMethodHS256)
	const myToken = "test"
	// Set a header and a claim
	fmt.Println(u.Name)
	token.Header["none"] = u.Id
	token.Claims["Id"] = u.Id
	token.Claims["Name"] = u.Name
	token.Claims["Gender"] = u.Gender
	token.Claims["Age"] = u.Age
	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, _ := token.SignedString([]byte(myToken))
	c.JSON(201, t)
}