package model

import (
	"io"
	"net/http"
)

type Domitory struct {
	AreaName string
	AreaID   string
}

func GetAreaNameAndID() (string, error) {
	req, err := http.NewRequest("GET", "https://jnb.ccnu.edu.cn/icbs/PurchaseWebService.asmx/getAreaInfo", nil)
	if err != nil {
		return "", err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func GetArchitectureNameAndID(areaID string) (string, error) {
	url := "https://jnb.ccnu.edu.cn/icbs/PurchaseWebService.asmx/getArchitectureInfo?Area_ID=" + areaID
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func GetRoomNameAndID(arcID, floor string) (string, error) {
	url := "https://jnb.ccnu.edu.cn/icbs/PurchaseWebService.asmx/getRoomInfo?Architecture_ID=" + arcID + "&Floor=" + floor
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func GetAmMeterID(roomID string) (string, error) {
	url := "https://jnb.ccnu.edu.cn/icbs/PurchaseWebService.asmx/getMeterInfo?Room_ID=" + roomID
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func GetEnergyList(amID, sT, eT string) (string, error) {
	url := "https://jnb.ccnu.edu.cn/icbs/PurchaseWebService.asmx/getMeterDayValue?AmMeter_ID=" + amID + "&startDate=" + sT + "&endDate=" + eT
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
