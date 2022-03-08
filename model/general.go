package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/* 一般天氣預報-今明 36 小時天氣預報 F-C0032-001*/
func (c *Client) GeneralWeatherForecast36() (*GeneralResponse, error) {
	res, err := http.Get(fmt.Sprintf(c.BASE_URL, "F-C0032-001", c.API_KEY))
	if err != nil {
		log.Fatal("Error When Fetching Data")
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error When Reading Body")
	}

	var gf = new(GeneralForecast)
	if err := json.Unmarshal(body, &gf); err != nil || gf.Success == "false" {
		return &GeneralResponse{}, errors.New("bad response json")
	}
	var gs = new(GeneralResponse)
	{
		gs.ResourceId = gf.Result.ResourceId
		gs.Description = gf.Records.DatasetDescription
		gs.Region = "台灣"
		for _, data := range gf.Records.Location {
			var loc GeneralLocation
			loc.Name = data.LocationName
			for _, ele := range data.WeatherElements {
				switch ele.ElementName {
				case "Wx":
					for _, t := range ele.Times {
						loc.Wx = append(loc.Wx, TimeSection{
							StartTime: t.StartTime,
							EndTime:   t.EndTime,
							Key:       t.Parameter.ParameterName,
							Value:     t.Parameter.ParameterName,
						})
					}
				case "PoP":
					for _, t := range ele.Times {
						loc.PoP = append(loc.PoP, TimeSection{
							StartTime: t.StartTime,
							EndTime:   t.EndTime,
							Key:       t.Parameter.ParameterUnit, // %
							Value:     t.Parameter.ParameterName,
						})
					}
				case "CI":
					for _, t := range ele.Times {
						loc.CI = append(loc.CI, TimeSection{
							StartTime: t.StartTime,
							EndTime:   t.EndTime,
							Key:       t.Parameter.ParameterName,
							Value:     t.Parameter.ParameterName,
						})
					}
				case "MinT":
					for _, t := range ele.Times {
						loc.MinT = append(loc.MinT, TimeSection{
							StartTime: t.StartTime,
							EndTime:   t.EndTime,
							Key:       t.Parameter.ParameterUnit, // C
							Value:     t.Parameter.ParameterName,
						})
					}
				case "MaxT":
					for _, t := range ele.Times {
						loc.MaxT = append(loc.MaxT, TimeSection{
							StartTime: t.StartTime,
							EndTime:   t.EndTime,
							Key:       t.Parameter.ParameterUnit, // C
							Value:     t.Parameter.ParameterName,
						})
					}
				}
			}
			gs.Locations = append(gs.Locations, loc)
		}
	}
	return gs, nil
}

type GeneralResponse struct {
	ResourceId  string
	Description string
	Region      string
	Locations   []GeneralLocation
}

type GeneralLocation struct {
	Name string
	Wx   []TimeSection
	PoP  []TimeSection
	CI   []TimeSection
	MinT []TimeSection
	MaxT []TimeSection
}

type TimeSection struct {
	StartTime string
	EndTime   string
	Key       string
	Value     string
}

type GeneralForecast struct {
	Success string `json:"success"`
	Result  struct {
		ResourceId string `json:"resource_id"`
		Fields     []struct {
			Id   string `json:"id"`
			Type string `json:"type"`
		} `json:"fields"`
	} `json:"result"`
	Records struct {
		DatasetDescription string `json:"datasetDescription"`
		/* 某個縣市的所有的天氣資料 */
		Location []struct {
			/* 縣市名稱 */
			LocationName string `json:"locationName"`
			/* 各種類天氣資料 */
			WeatherElements []struct {
				/* 資料種類名稱 */
				ElementName string `json:"elementName"`
				/* 該種類所有時段的天氣資料 */
				Times []struct {
					StartTime string `json:"startTime"`
					EndTime   string `json:"endTime"`
					/* 資料值 */
					Parameter struct {
						ParameterName string `json:"parameterName"`
						/* May Not Exist */
						ParameterValue string `json:"parameterValue"`
						/* May Not Exist */
						ParameterUnit string `json:"parameterUnit"`
					} `json:"parameter"`
				} `json:"time"`
			} `json:"weatherElement"`
		} `json:"location"`
	} `json:"records"`
}

type GeneralConfig struct {
	/* 限制最多回傳的資料筆數(1~22)，預設為全部筆數 */
	Limit int
	/* 指定從第幾筆後開始回傳(0~21)，預設為0 */
	Offset int
}

var Element = struct {
	Wx   string
	PoP  string
	CI   string
	MinT string
	MaxT string
}{
	Wx:   "Wx",
	PoP:  "PoP",
	CI:   "CI",
	MinT: "MinT",
	MaxT: "MaxT",
}
