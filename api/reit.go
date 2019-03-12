package api

import (
	"../models"
	"../services"
	"../util"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

type ReitController interface {
	GetReitAll(c echo.Context) error
	GetReitBySymbol(c echo.Context) error
	GetFavoriteReitAll(c echo.Context) error
	SaveFavoriteReit(c echo.Context) error
	DeleteFavoriteReit(c echo.Context) error
}

type Reit struct {
	reitServicer services.ReitServicer
	reitItems []*models.ReitItem
	reitItem models.ReitItem
	reitFavorite []*models.FavoriteInfo
	err error
}


func DeleteFavoriteReitProcess(c echo.Context) error {
	var reitController ReitController
	reitController = Reit {}
	return reitController.DeleteFavoriteReit(c)
}

func GetReitAllProcess(c echo.Context) error {
	var reitController ReitController
	reitController = Reit {}
	return reitController.GetReitAll(c)
}

func GetReitBySymbolProcess(c echo.Context) error {
	var reitController ReitController
	reitController = Reit {}
	return reitController.GetReitBySymbol(c)
}

func GetFavoriteReitAllProcess(c echo.Context) error {
	var reitController ReitController
	reitController = Reit {}
	return reitController.GetFavoriteReitAll(c)
}

func SaveFavoriteReitProcess(c echo.Context) error {
	var reitController ReitController
	reitController = Reit {}
	return reitController.SaveFavoriteReit(c)
}

// Handler
func (self Reit) GetReitAll(c echo.Context) error {

	self.reitServicer = services.Reit_Service{}
	self.reitItems ,self.err = services.GetReitAllProcess(self.reitServicer)
	if self.err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "data not found")
	}
	return c.JSON(http.StatusOK, self.reitItems)
}

func (self Reit) GetReitBySymbol(c echo.Context) error {
	symbol := c.Param("symbol")
	self.reitServicer = services.Reit_Service{}
	self.reitItem, self.err = services.GetReitBySymbolProcess(self.reitServicer,symbol)
	if self.reitItem == (models.ReitItem{}) {
		return echo.NewHTTPError(http.StatusNotFound, "data not found")
	}
	if self.err != nil {
		return echo.NewHTTPError(http.StatusOK, self.err)
	}
	return c.JSON(http.StatusOK, self.reitItem)
}

func (self Reit) GetFavoriteReitAll(c echo.Context) error {
	fmt.Println("start : GetFavoriteReitAll")
	userID := c.Param("id")
	self.reitServicer = services.Reit_Service{}
	self.reitFavorite = services.GetReitFavoriteByUserIDJoinProcess(self.reitServicer,userID)
	return c.JSON(http.StatusOK, self.reitFavorite)
}

func (self Reit) SaveFavoriteReit(c echo.Context) error {
	// Get name and email
	fmt.Println("start : SaveFavoriteReit")
	userID := c.FormValue("userId")
	ticker := c.FormValue("Ticker")
	self.reitServicer = services.Reit_Service{}
	self.err = services.SaveReitFavoriteProcess(self.reitServicer, userID, ticker)
	if self.err != nil {
		return c.String(http.StatusBadRequest, "fail")
	}
	return c.String(http.StatusOK, "success")
}

func (self Reit) DeleteFavoriteReit(c echo.Context) error {
	// Get name and email
	fmt.Println("start : DeleteFavoriteReit")
	userID := c.FormValue("userId")
	ticker := c.FormValue("Ticker")
	self.reitServicer = services.Reit_Service{}
	self.err = services.DeleteReitFavoriteProcess(self.reitServicer , userID, ticker)
	if self.err != nil {
		return c.String(http.StatusBadRequest, "fail")
	}
	return c.String(http.StatusOK, "success")
}

func GetUserProfile(c echo.Context) error {
	userID,site := util.GetUserFromToken(c);
	profile := services.GetUserProfileByCriteria(userID, site)
	return c.JSON(http.StatusOK, profile)
}

