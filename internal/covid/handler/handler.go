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

func New(r *gin.RouterGroup, serv service.Covider) {
	handle := &handler{
		serv: serv,
	}

	r.GET("/summary", handle.summary)
}

func (h *handler) summary(c *gin.Context) {
	resp, err := h.serv.CovidData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	provinceCount := make(map[string]int)
	ageCount := make(map[string]int)
	for _, row := range resp {
		if row.ProvinceEn != "" {
			provinceCount[row.ProvinceEn] += 1
		}
		switch rangeage.FindRangeAge(row.Age) {
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
