package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pitsanujiw/go-covid/internal/constant"
	"github.com/pitsanujiw/go-covid/internal/entity"
	"github.com/stretchr/testify/require"
)

type mockHttpClient struct {
	err  error
	body io.ReadCloser
}

func (h *mockHttpClient) Get(url string) (resp *http.Response, err error) {
	return &http.Response{
		Body: h.body,
	}, h.err
}

const mock = `{"Data":[{"ConfirmDate":"2021-05-04","No":null,"Age":51,"Gender":"หญิง","GenderEn":"Female","Nation":null,"NationEn":"China","Province":"Phrae","ProvinceId":46,"District":null,"ProvinceEn":"Phrae","StatQuarantine":5}]}`

func TestCovidData(t *testing.T) {
	t.Run("Should return covid summary", func(t *testing.T) {
		rr := httptest.NewRecorder()

		rr.WriteHeader(http.StatusOK)

		client := &mockHttpClient{
			err:  nil,
			body: ioutil.NopCloser(bytes.NewReader([]byte(mock))),
		}

		serv := NewCovidServ(client)

		got, err := serv.CovidData()
		if err != nil {
			fmt.Println(err)
		}

		var want entity.CovidResponse
		data := []byte(mock)

		if err = json.Unmarshal(data, &want); err != nil {
			return
		}

		require.Equal(t, got, want.Data)
	})

	t.Run("Should throw error if a service error", func(t *testing.T) {
		rr := httptest.NewRecorder()

		rr.WriteHeader(http.StatusOK)

		client := &mockHttpClient{
			err:  errors.New("error"),
			body: nil,
		}

		serv := NewCovidServ(client)

		_, err := serv.CovidData()

		require.Equal(t, err, constant.ErrGetRequestError)
	})

	t.Run("", func(t *testing.T) {
		rr := httptest.NewRecorder()

		rr.WriteHeader(http.StatusOK)

		client := &mockHttpClient{
			err:  nil,
			body: ioutil.NopCloser(bytes.NewReader([]byte(""))),
		}

		serv := NewCovidServ(client)

		_, err := serv.CovidData()

		require.Equal(t, err, constant.ErrConvertToStructError)
	})

}
