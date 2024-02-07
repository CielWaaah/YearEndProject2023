package service

import (
	"fmt"
	"io"
	"net/http"
)

type DomitoryService struct {
	AreaName   string
	AreaID     string
	AmMeter_ID string
}

func (service *DomitoryService) GetAreaNameAndID() (string, error) {
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

func (service *DomitoryService) GetArchitectureNameAndID(areaID string) (string, error) {
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

func (service *DomitoryService) GetRoomNameAndID(arcID, floor string) (string, error) {
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

func (service *DomitoryService) GetAmMeterID(roomID string) (string, error) {
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

func (service *DomitoryService) GetEnergyList(amID, sT, eT string) (string, error) {
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

func (service *DomitoryService) GetMoney(amID ,sM, eM string) (string, error) {
	url := fmt.Sprintf("https://jnb.ccnu.edu.cn/icbs/PurchaseWebService.asmx/getMeterMonthValue?AmMeter_ID=" + amID + "&startMonth=" + sM + "&endMonth=" + eM)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(body), nil
}
