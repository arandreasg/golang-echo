package controller

import (
	"fmt"
	"net/http"

	"pegadaianempat/config"
	"pegadaianempat/helpers"
	"pegadaianempat/model"

	"github.com/labstack/echo/v4"
)

func UserLogin(c echo.Context) error {
	db := config.GetDB()

	Emp := model.Employee{}
	password := ""

	if err := c.Bind(&Emp); err != nil {
		return err
	}

	password = Emp.Password

	if err := db.Debug().Where("email =?", Emp.Email).Take(&Emp).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
	}

	fmt.Println(password)
	fmt.Println(Emp.Password)
	ComparePass := helpers.ComparePass([]byte(Emp.Password), []byte(password))

	if !ComparePass {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
	}

	fmt.Println("TEST", uint(Emp.ID))
	fmt.Println("TEST2", Emp.Email)

	token := helpers.GenerateToken(uint(Emp.ID), Emp.Email)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
