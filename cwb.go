package cwb

import (
	"github.com/fatsoap/cwb-sdk-go/model"
)

func New(API_KEY string) (*model.Client, error) {
	return &model.Client{
		API_KEY:  API_KEY,
		BASE_URL: "https://opendata.cwb.gov.tw/api/v1/rest/datastore/%s?Authorization=%s"}, nil
}
