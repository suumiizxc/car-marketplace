package customer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	// models "github.com/suumiizxc/gin-bookstore/models/core/customer"
)

type CreateCustomerInput struct {
	CustSegCode       string `json:"custSegCode"`
	SexCode           uint   `json:"sexCode"`
	TaxExemption      uint   `json:"taxExemption"`
	Status            uint   `json:"status"`
	IsCompanyCustomer uint   `json:"isCompanyCustomer"`
	IndustryID        uint   `json:"industryId"`
	BirthPlaceID      uint   `json:"birthPlaceId"`
	FamilyName        string `json:"familyName"`
	FamilyName2       string `json:"familyName2"`
	LastName          string `json:"lastName"`
	LastName2         string `json:"lastName2"`
	FirstName         string `json:"firstName"`
	FirstName2        string `json:"firstName2"`
	ShortName         string `json:"shortName"`
	ShortName2        string `json:"shortName2"`
	RegisterMaskCode  string `json:"registerMaskCode"`
	RegisterCode      string `json:"registerCode"`
	BirthDate         string `json:"birthDate"`
	Mobile            string `json:"mobile"`
	CountryCode       string `json:"countryCode"`
	EmploymentID      uint   `json:"employmentId"`
	Email             string `json:"email"`
	IndustryName      string `json:"industryName"`
	CatID             uint   `json:"catId"`
	TitleID           uint   `json:"titleId"`
	NationalityID     uint   `json:"nationalityId"`
	EthnicGroupID     uint   `json:"ethnicGroupId"`
	LangCode          string `json:"langCode"`
	MaritalStatus     uint   `json:"maritalStatus"`
	BirthPlaceName    string `json:"birthPlaceName"`
	BirthPlaceDetail  string `json:"birthPlaceDetail"`
	EducationID       uint   `json:"educationId"`
	Phone             string `json:"phone"`
	Fax               string `json:"fax"`
	CreatedBy         uint   `json:"createdBy"`
	ApprovedBy        uint   `json:"approvedBy"`
	CompanyCode       string `json:"companyCode"`
	IsVatPayer        uint   `json:"isVatPayer"`
}

func CreateCustomer(c *gin.Context) {
	input := CreateCustomerInput{}
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var inputs []CreateCustomerInput
	inputs = append(inputs, input)
	fmt.Println("inputs : ", inputs)
	json_data, err := json.Marshal(inputs)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("bodyData : ", json_data)
	req, err := http.NewRequest("POST", "http://202.131.242.158:4020/nes.s.Web/NesFront", bytes.NewBuffer(json_data))

	if err != nil {
		log.Printf("Request failed : %s", err.Error())
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", "NESSESSION=22QivFUT3jGC681187SzyfmEzJn7DL")
	req.Header.Add("op", "13610313")
	req.Header.Add("company", "11")
	req.Header.Add("lang", "1")
	req.Header.Add("role", "53")

	client := &http.Client{Timeout: time.Second * 10}

	fmt.Println("req : ", req.Body)
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal("Error reading response : ", err)
	}
	defer resp.Body.Close()
	fmt.Println("status code :", resp.StatusCode)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Request failed : %s", err)
	}
	bodyString := string(body)
	log.Print(bodyString)

}