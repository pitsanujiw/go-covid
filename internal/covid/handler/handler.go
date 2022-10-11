package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/pitsanujiw/go-covid/internal/constant"
	"github.com/pitsanujiw/go-covid/internal/covid/rangeage"
	"github.com/pitsanujiw/go-covid/internal/covid/service"
	"github.com/pitsanujiw/go-covid/internal/entity"
)

type handler struct {
	serv service.Covider
}

// new Handler
func New(r *gin.RouterGroup, serv service.Covider) {
	handle := &handler{
		serv: serv,
	}

	// mapping a handler with method type and path
	r.GET("/summary", handle.summary)
}

func (h *handler) summary(c *gin.Context) {
	// get data from service
	covidInfos, err := h.serv.CovidData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	// mapping data
	provinceCount := make(map[string]int)
	ageCount := make(map[string]int)
	// lopping get data from
	for _, covidInfo := range covidInfos {
		// validate have province or not
		if covidInfo.ProvinceEn != "" {
			provinceCount[covidInfo.ProvinceEn] += 1
		}

		// mapping range of age
		switch rangeage.FindRangeAge(covidInfo.Age) {
		case constant.UNKNOWN:
			ageCount["N/A"] += 1
		case constant.ADULT:
			ageCount["0-30"] += 1
		case constant.OLD:
			ageCount["31-60"] += 1
		case constant.ELDER:
			ageCount["61+"] += 1
		}
	}

	res := entity.CovidSummaryResponse{
		Province: provinceCount,
		AgeGroup: ageCount,
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(http.StatusOK, res)
}
