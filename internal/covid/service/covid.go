package service

import (
	"encoding/json"
	"io"

	"github.com/pitsanujiw/go-covid/internal/constant"
	"github.com/pitsanujiw/go-covid/internal/entity"
	"github.com/pitsanujiw/go-covid/pkg/httpclient"
)

type covid struct {
	client httpclient.HttpClienter
}

type Covider interface {
	CovidData() (data []entity.CovidRecord, err error)
}

func NewCovidServ(client httpclient.HttpClienter) *covid {
	return &covid{
		client: client,
	}
}

func (c *covid) CovidData() ([]entity.CovidRecord, error) {
	resp, err := c.client.Get(constant.CovidUrl)
	if err != nil {
		return nil, constant.ErrGetRequestError
	}

	defer resp.Body.Close()
	body, bodyErr := io.ReadAll(resp.Body)

	if bodyErr != nil {
		return nil, constant.ErrGetRequestError
	}

	var records entity.CovidResponse

	if convertErr := json.Unmarshal(body, &records); convertErr != nil {
		return nil, constant.ErrConvertToStructError
	}

	return records.Data, nil
}
