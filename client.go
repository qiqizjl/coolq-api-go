package coolq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type coolQResult struct {
	Status  string      `json:"status"`
	Retcode int         `json:"retcode"`
	Data    interface{} `json:"data"`
}

func (coolq *CoolQ) httpPOST(route string, data map[string]interface{}) (error, interface{}) {
	url := coolq.URL + route
	reqBody, err := json.Marshal(data)
	if err != nil {
		return err, nil
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if coolq.Token != "" {
		req.Header.Set("Authorization", "token "+coolq.Token)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err, nil
	}
	defer resp.Body.Close()
	//先判断HTTPCode异常
	switch resp.StatusCode {
	case 404:
		return fmt.Errorf("routFoundRouter"), nil
	case 401:
		return fmt.Errorf("TokenError"), nil
	}
	//先读取Body
	body, _ := ioutil.ReadAll(resp.Body)
	//解析JSON
	result := coolQResult{}
	json.Unmarshal([]byte(body), &result)
	fmt.Println(string(body))
	if result.Status != "ok" {
		//处理retcode
		if result.Retcode == 100 {
			return fmt.Errorf("inputError"), nil
		} else if result.Retcode == 102 {
			return fmt.Errorf("HaveError"), nil
		} else {
			return fmt.Errorf("CoolQ Error" + fmt.Sprintf("%d", result.Retcode)), nil
		}
	}
	return nil, result.Data
}
