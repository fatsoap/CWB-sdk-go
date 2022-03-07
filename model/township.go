package model

/* 鄉鎮天氣預報 未來 2,7 天天氣預報 */
/* F-C0047-001 */
type TownshipWeatherForecast struct {
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
