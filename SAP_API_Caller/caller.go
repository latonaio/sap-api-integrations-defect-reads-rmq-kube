package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-defect-reads-rmq-kube/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

type RMQOutputter interface {
	Send(sendQueue string, payload map[string]interface{}) error
}

type SAPAPICaller struct {
	baseURL      string
	apiKey       string
	outputQueues []string
	outputter    RMQOutputter
	log          *logger.Logger
}

func NewSAPAPICaller(baseUrl string, outputQueueTo []string, outputter RMQOutputter, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:      baseUrl,
		apiKey:       GetApiKey(),
		outputQueues: outputQueueTo,
		outputter:    outputter,
		log:          l,
	}
}

func (c *SAPAPICaller) AsyncGetDefect(defect, defectText string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(defect)
				wg.Done()
			}()
		case "DefectText":
			func() {
				c.DefectText(defectText)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) Header(defect string) {
	data, err := c.callDefectSrvAPIRequirementHeader("A_Defect", defect)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": data, "function": "DefectHeader"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callDefectSrvAPIRequirementHeader(api, defect string) ([]sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "API_DEFECT_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithHeader(req, defect)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) DefectText(defectText string) {
	data, err := c.callDefectSrvAPIRequirementDefectText("A_Defect", defectText)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": data, "function": "DefectHeader"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callDefectSrvAPIRequirementDefectText(api, defectText string) ([]sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "API_DEFECT_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithDefectText(req, defectText)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithHeader(req *http.Request, defect string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("Defect eq '%s'", defect))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithDefectText(req *http.Request, defectText string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("substringof('%s', DefectText)", defectText))
	req.URL.RawQuery = params.Encode()
}
