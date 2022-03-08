package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (c *Client) TownshipWeatherForecast() (*TownshipResponse, error) {
	res, err := http.Get(fmt.Sprintf(c.BASE_URL, "F-D0047-001", c.API_KEY))
	if err != nil {
		log.Fatal("Error When Fetching Data")
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error When Reading Body")
	}
	var tf = new(TownshipForecast)
	if err := json.Unmarshal(body, &tf); err != nil || tf.Success == "false" {
		return &TownshipResponse{}, errors.New("bad response json")
	}
	var ts = new(TownshipResponse)
	{
		ts.ResourceId = tf.Result.ResourceId
		ts.Description = tf.Records.Locations[0].DatasetDescription
		ts.Region = tf.Records.Locations[0].LocationName
		for _, data := range tf.Records.Locations[0].Location {
			var loc TownshipLocation
			loc.Name = data.LocationName
			loc.Geocode = data.Geocode
			loc.Lontitude = data.Lontitude
			loc.Latitude = data.Latitude
			for _, ele := range data.WeatherElements {
				switch ele.ElementName {
				case "PoP12h":
					for _, t := range ele.Times {
						loc.PoP12h = append(loc.PoP12h, TimeSection{
							StartTime: t.StartTime,
							EndTime:   t.EndTime,
							Key:       t.ElementValue[0].Measures, // %
							Value:     t.ElementValue[0].Value,    // int
						})
					}
				case "Wx":
					for _, t := range ele.Times {
						loc.Wx = append(loc.Wx, TimeSection{
							StartTime: t.StartTime,
							EndTime:   t.EndTime,
							Key:       t.ElementValue[0].Value, // desciption
							Value:     t.ElementValue[1].Value, // int
						})
					}
				case "AT":
					for _, t := range ele.Times {
						loc.AT = append(loc.AT, TimeSection{
							StartTime: t.StartTime,
							EndTime:   t.EndTime,
							Key:       t.ElementValue[0].Measures, // C
							Value:     t.ElementValue[0].Value,    // int
						})
					}
				case "T":
					for _, t := range ele.Times {
						loc.T = append(loc.T, TimeSection{
							StartTime: t.StartTime,
							EndTime:   t.EndTime,
							Key:       t.ElementValue[0].Measures, // C
							Value:     t.ElementValue[0].Value,    // int
						})
					}
				case "RH":
					for _, t := range ele.Times {
						loc.RH = append(loc.RH, TimeSection{
							StartTime: t.StartTime,
							EndTime:   t.EndTime,
							Key:       t.ElementValue[0].Measures, // %
							Value:     t.ElementValue[0].Value,    // int
						})
					}

				case "CI":
					for _, t := range ele.Times {
						loc.CI = append(loc.CI, TimeSection{
							StartTime: t.StartTime,
							EndTime:   t.EndTime,
							Key:       t.ElementValue[1].Value, // decsription
							Value:     t.ElementValue[0].Value, // int
						})
					}

				case "WeatherDescription":
					for _, t := range ele.Times {
						loc.WeatherDescription = append(loc.WeatherDescription, TimeSection{
							StartTime: t.StartTime,
							EndTime:   t.EndTime,
							Key:       t.ElementValue[0].Measures, // X
							Value:     t.ElementValue[0].Value,    // decription
						})
					}

				case "PoP6h":
					for _, t := range ele.Times {
						loc.PoP6h = append(loc.PoP6h, TimeSection{
							StartTime: t.StartTime,
							EndTime:   t.EndTime,
							Key:       t.ElementValue[0].Measures, // %
							Value:     t.ElementValue[0].Value,    // int
						})
					}
				case "WS":
					for _, t := range ele.Times {
						loc.WS = append(loc.WS, TimeSection{
							StartTime: t.StartTime,
							EndTime:   t.EndTime,
							Key:       t.ElementValue[0].Measures, // meter/sec
							Value:     t.ElementValue[0].Value,    // int
						})
					}
				case "WD":
					for _, t := range ele.Times {
						loc.WD = append(loc.WD, TimeSection{
							StartTime: t.StartTime,
							EndTime:   t.EndTime,
							Key:       t.ElementValue[0].Measures, // 8方位
							Value:     t.ElementValue[0].Value,    // description
						})
					}
				case "Td":
					for _, t := range ele.Times {
						loc.Td = append(loc.Td, TimeSection{
							StartTime: t.StartTime,
							EndTime:   t.EndTime,
							Key:       t.ElementValue[0].Measures, // c
							Value:     t.ElementValue[0].Value,    // int
						})
					}
				}
			}
			ts.Locations = append(ts.Locations, loc)
		}
	}
	return ts, nil
}

type TownshipLocation struct {
	Name      string
	Geocode   string
	Latitude  string
	Lontitude string

	PoP12h             []TimeSection
	Wx                 []TimeSection
	AT                 []TimeSection
	T                  []TimeSection
	RH                 []TimeSection
	CI                 []TimeSection
	WeatherDescription []TimeSection
	PoP6h              []TimeSection
	WS                 []TimeSection
	WD                 []TimeSection
	Td                 []TimeSection
}

/* TODO : add 蒲福風級 field in WS section */

type TownshipResponse struct {
	ResourceId  string
	Description string
	Region      string
	Locations   []TownshipLocation
}

/* 鄉鎮天氣預報 未來 2,7 天天氣預報 */
/* F-C0047-001 */
type TownshipForecast struct {
	Success string `json:"success"`
	Result  struct {
		ResourceId string `json:"resource_id"`
		Fields     []struct {
			Id   string `json:"id"`
			Type string `json:"type"`
		} `json:"fields"`
	} `json:"result"`
	Records struct {
		/* 某個縣市的所有的天氣資料 */
		Locations []struct {
			DatasetDescription string `json:"datasetDescription"`
			/* 縣市名稱 */
			LocationName string `json:"locationsName"`
			Dataid       string `json:"dataid"`
			/* 鄉鎮市區的天氣資料 */
			Location []struct {
				/* 鄉鎮市區的名稱 */
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
						ElementValue []struct {
							Value    string `json:"value"`
							Measures string `json:"measures"`
						} `json:"elementValue"`
					} `json:"time"`
					/* May Not Exist */
					Description string `json:"description"`
				} `json:"weatherElement"`
				/* 地址編碼 */
				Geocode string `json:"geocode"`
				/* 緯度 */
				Latitude string `json:"lat"`
				/* 經度 */
				Lontitude string `json:"lon"`
			} `json:"location"`
		} `json:"locations"`
	} `json:"records"`
}
