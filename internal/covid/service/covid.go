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

func (c *covid) CovidData() (data []entity.CovidRecord, err error) {
	resp, err := c.client.Get(constant.CovidUrl)
	if err != nil {
		return nil, constant.ErrGetRequestError
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	var records entity.CovidResponse
	if err = json.Unmarshal(body, &records); err != nil {
		return nil, constant.ErrConvertToStructError
	}

	return records.Data, nil
}
