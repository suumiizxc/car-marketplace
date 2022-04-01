package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type core_helper struct {
	polaris_company string
	polaris_lang    string
	polaris_role    string
	polaris_url     string
}

func (corehelper *core_helper) Init() {
	if err := godotenv.Load(".env"); err != nil {
		panic(fmt.Sprintf("Failed env : %v", err))
	}
	corehelper.polaris_company = os.Getenv("POLARIS_COMPANY")
	corehelper.polaris_lang = os.Getenv("POLARIS_LANG")
	corehelper.polaris_role = os.Getenv("POLARIS_ROLE")
	corehelper.polaris_url = os.Getenv("POLARIS_URL")
}

// func (corehelper core_helper) setNewName() {
// 	if err := godotenv.Load(".env"); err != nil {
// 		panic(fmt.Sprintf("Failed env : %v", err))
// 	}
// 	corehelper.polaris_company = os.Getenv("POLARIS_COMPANY")
// 	corehelper.polaris_lang = os.Getenv("POLARIS_LANG")
// 	corehelper.polaris_role = os.Getenv("POLARIS_ROLE")
// 	corehelper.polaris_url = os.Getenv("POLARIS_URL")
// }

func (corehelper core_helper) Request(opcode string, token string, field []byte) ([]interface{}, error) {
	req, err := http.NewRequest("POST", corehelper.polaris_url, bytes.NewBuffer(field))

	if err != nil {
		log.Printf("Request failed : %s", err.Error())
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", "NESSESSION="+token)
	req.Header.Add("op", opcode)
	req.Header.Add("company", corehelper.polaris_company)
	req.Header.Add("lang", corehelper.polaris_lang)
	req.Header.Add("role", corehelper.polaris_role)

	client := &http.Client{Timeout: time.Second * 3}
	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("error : %v", err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Request failed : %s", err)
	}
	bodyString := string(body)

	var tmp []interface{}
	var data = []byte(bodyString)
	if err := json.Unmarshal(data, &tmp); err != nil {
		return nil, err
	}
	return tmp, nil
}

var CH = new(core_helper)
