package controllers

import (
	"revelSampleApp/app/models"

	"github.com/revel/revel"
)

type Cars struct {
	*revel.Controller
}

var cars = []models.Car{
	models.Car{1, "Bugatti Chiron", "Automatic", "8.0 L (488 cu in) Quad-Turbocharged W16"},
	models.Car{2, "Bentley Mulsanne", "Automatic", "Twin Turbocharged V8 Engine"},
	models.Car{3, "Audi A3 Cabriolet", "Automatic", "TFSI Petrol Engine"},
}

func (c Cars) List() revel.Result {
	return c.RenderJSON(cars)
}

func (c Cars) Show(carID int) revel.Result {
	var res models.Car

	for _, car := range cars {
		if car.ID == carID {
			res = car
		}
	}

	if res.ID == 0 {
		return c.NotFound("Could not find Car")
	}

	return c.RenderJSON(res)
}
