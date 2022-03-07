package model

/* 一般天氣預報-今明 36 小時天氣預報 */
/* F-C0032-001 */
type TaiwanForecast36 struct {
	Success string `json:"success"`
	Result  struct {
		Resource_id string `json:"resource_id"`
		Fields      []struct {
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
