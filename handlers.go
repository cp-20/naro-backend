package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type City struct {
	ID          int            `json:"id,omitempty"  db:"ID"`
	Name        sql.NullString `json:"name,omitempty"  db:"Name"`
	CountryCode sql.NullString `json:"countryCode,omitempty"  db:"CountryCode"`
	District    sql.NullString `json:"district,omitempty"  db:"District"`
	Population  sql.NullInt64  `json:"population,omitempty"  db:"Population"`
}

func getCityInfoHandler(c echo.Context) error {
	cityId, err := strconv.Atoi(c.Param("cityId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid city id")
	}

	var city City
	db.Get(&city, "SELECT * FROM city WHERE ID = ?", cityId)
	if !city.Name.Valid {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, city)
}

func postCityHandler(c echo.Context) error {
	var city City
	err := c.Bind(&city)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad request body")
	}

	result, err := db.Exec("INSERT INTO city (Name, CountryCode, District, Population) VALUES (?, ?, ?, ?)", city.Name, city.CountryCode, city.District, city.Population)
	if err != nil {
		log.Printf("failed to insert city data: %s\n", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("failed to get last insert id: %s\n", err)
		return c.NoContent(http.StatusInternalServerError)
	}
	city.ID = int(id)

	return c.JSON(http.StatusCreated, city)
}

type LoginRequestBody struct {
	Username string `json:"username,omitempty" form:"username"`
	Password string `json:"password,omitempty" form:"password"`
}

func signUpHandler(c echo.Context) error {
	var req LoginRequestBody
	c.Bind(&req)

	if req.Username == "" || req.Password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Username or Password is empty")
	}

	var count int
	err := db.Get(&count, "SELECT COUNT(*) FROM users WHERE Username=?", req.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.NoContent(http.StatusUnauthorized)
		} else {
			log.Println(err)
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	if count > 0 {
		return echo.NewHTTPError(http.StatusConflict, "username is already used")
	}

	pw := req.Password + salt

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	_, err = db.Exec("INSERT INTO users (Username, HashedPass) VALUES (?, ?)", req.Username, hashedPass)
	if err != nil {
		log.Println(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusCreated)
}

type User struct {
	Username   string `json:"username,omitempty"  db:"Username"`
	HashedPass string `json:"-"  db:"HashedPass"`
}

func loginHandler(c echo.Context) error {
	var req LoginRequestBody
	c.Bind(&req)

	if req.Password == "" || req.Username == "" {
		return c.String(http.StatusBadRequest, "Username or Password is empty")
	}

	user := User{}
	err := db.Get(&user, "SELECT * FROM users WHERE username=?", req.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.NoContent(http.StatusUnauthorized)
		}

		log.Println(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPass), []byte(req.Password+salt))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return c.NoContent(http.StatusUnauthorized)
		}

		return c.NoContent(http.StatusInternalServerError)
	}

	sess, err := session.Get("sessions", c)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "something wrong in getting session")
	}
	sess.Values["userName"] = req.Username
	sess.Save(c.Request(), c.Response())

	return c.NoContent(http.StatusOK)
}

type Me struct {
	Username string `json:"username,omitempty"  db:"username"`
}

func getWhoAmIHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, Me{
		Username: c.Get("userName").(string),
	})
}

type Country struct {
	Code           string          `json:"code,omitempty"  db:"Code"`
	Name           sql.NullString  `json:"name,omitempty"  db:"Name"`
	Continent      sql.NullString  `json:"continent,omitempty"  db:"Continent"`
	Region         sql.NullString  `json:"region,omitempty"  db:"Region"`
	SurfaceArea    sql.NullFloat64 `json:"surfaceArea,omitempty"  db:"SurfaceArea"`
	IndepYear      sql.NullInt64   `json:"indepYear,omitempty"  db:"IndepYear"`
	Population     sql.NullInt64   `json:"population,omitempty"  db:"Population"`
	LifeExpectancy sql.NullFloat64 `json:"lifeExpectancy,omitempty"  db:"LifeExpectancy"`
	GNP            sql.NullFloat64 `json:"gnp,omitempty"  db:"GNP"`
	GNPOld         sql.NullFloat64 `json:"gnpOld,omitempty"  db:"GNPOld"`
	LocalName      sql.NullString  `json:"localName,omitempty"  db:"LocalName"`
	GovernmentForm sql.NullString  `json:"governmentForm,omitempty"  db:"GovernmentForm"`
	HeadOfState    sql.NullString  `json:"headOfState,omitempty"  db:"HeadOfState"`
	Capital        sql.NullInt64   `json:"capital,omitempty"  db:"Capital"`
	Code2          sql.NullString  `json:"code2,omitempty"  db:"Code2"`
}

func getCountriesHandler(c echo.Context) error {
	countries := []Country{}
	err := db.Select(&countries, "SELECT * FROM country")
	if err != nil {
		log.Println(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, countries)
}

func getCitiesHandler(c echo.Context) error {
	cities := []City{}
	err := db.Select(&cities, "SELECT * FROM city WHERE CountryCode=?", c.Param("countryCode"))
	if err != nil {
		log.Println(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, cities)
}
