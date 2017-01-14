package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

type Geolocation struct {
	Altitude  float64
	Latitude  float64
	Longitude float64
}

func GetStreamResponse(c echo.Context) error {
	var locations = []Geolocation{
		{-97, 37.819929, -122.478255},
		{1899, 39.096849, -120.032351},
		{2619, 37.865101, -119.538329},
		{42, 33.812092, -117.918974},
		{15, 37.77493, -122.419416},
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().WriteHeader(http.StatusOK)
	for _, l := range locations {
		if err := json.NewEncoder(c.Response()).Encode(l); err != nil {
			return err
		}
		c.Response().Flush()
		time.Sleep(1 * time.Second)
	}
	return nil
}