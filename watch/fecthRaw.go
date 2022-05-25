package watch

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func (dartReport *NewReportWatch) OuterReportText() error {
	/*
		DART 주소에서 html 파일을 빼냄.
		진짜 정보는 iframe 안에 담겨있지만, iframe 안에 든 정보를 찾기 위한 것

		iframe을 찾기위한 DART 주소 html 전문을 내보냄.
		-----
		Retrieve html file from the DART url.
		Real information is hidden inside <iframe> block.
		This function helps retrieve information so that we can get an access to that block.

		Returns content string
	*/
	resp, err := http.Get(dartReport.Ctx.Link)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()

	// check request status code. if 200 proceed
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	} else {
		dartReport.OrgHTML = string(data)
		return nil
	}
}

func (dartReport *NewReportWatch) GetDcmNo() (string, string) {
	// dcmNo is crucial variable to retrieve iframe inner information
	i := strings.Index(dartReport.OrgHTML, dcmId)
	dcmNo := dartReport.OrgHTML[i+dcmParseIndSrt : i+dcmParseIndSrt+dcmLength]
	return "dcmNo", dcmNo
}

func (dartReport *NewReportWatch) GetRcpNo() (string, string) {
	// rcpNo is crucial variable to retrieve iframe inner information
	i := strings.Index(dartReport.OrgHTML, rcpId)
	rcpNo := dartReport.OrgHTML[i+rcpParseIndSrt : i+rcpParseIndSrt+rcpLength]
	return "rcpNo", rcpNo
}

func (dartReport *NewReportWatch) InnerReportURL() (string, error) {
	dcmKey, dcmVal := dartReport.GetDcmNo()
	rcpKey, rcpVal := dartReport.GetRcpNo()

	var InnerReportURLBase = "https://dart.fss.or.kr/report/viewer.do?"
	var InnerReportParam = ReportMap{
		dcmKey:   dcmVal,
		rcpKey:   rcpVal,
		"eleId":  "0",
		"offset": "0",
		"length": "0",
		"dtd":    "dart3.xsd",
	}

	// ceate request
	resp, err := http.NewRequest("GET", InnerReportURLBase, nil)
	if err != nil {
		log.Println(err)
		return "", err
	}

	// add parameter
	qry := resp.URL.Query()
	for key, element := range InnerReportParam {
		qry.Add(key, element)
	}

	// return URL result
	resp.URL.RawQuery = qry.Encode()

	dartReport.RawUrl = resp.URL.String()
	return resp.URL.String(), nil
}

func (dartReport *NewReportWatch) Run() interface{} {
	// 1. get outer html
	err := dartReport.OuterReportText()
	if err != nil {
		log.Println(err)
	}

	// 2. extract inner html url
	u, err := dartReport.InnerReportURL()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("URL:", u)

	// 3. insert inner html content
	resp, err := http.Get(dartReport.RawUrl)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	dartReport.RAWHTML = string(data)

	// 4. parse inner html content
	cnt, err := dartReport.parseEventDriven()

	// 5. finish
	if err != nil {
		log.Println("error", err)
	}
	return cnt
}
