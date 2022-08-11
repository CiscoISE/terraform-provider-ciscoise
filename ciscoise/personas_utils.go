package ciscoise

import (
	"archive/zip"
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"unicode/utf8"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"github.com/go-resty/resty/v2"
)

var Error map[string]interface{}

type NodeR struct {
	Response Node   `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}
type Node struct {
	Name     string   `json:"name,omitempty"`     //
	Ip       string   `json:"ip,omitempty"`       //
	HostName string   `json:"hostname,omitempty"` // Consent to import the self-signed certificate of the registering node.
	Fqdn     string   `json:"fqdn,omitempty"`     //
	Password string   `json:"password,omitempty"` //
	Roles    []string `json:"roles,omitempty"`    // Roles can be empty or have many values for a node.
	Services []string `json:"services,omitempty"` // Services can be empty or have many values for a node.
	UserName string   `json:"userName,omitempty"` //
}

// *********************************************Node Methods*******************************************************
func (node Node) IsStandAlone() (bool, error) {
	path := fmt.Sprintf("https://%s/api/v1/deployment/node/%s", node.Ip, node.HostName)

	log.Printf("[DEBUG] My Path %s", path)

	response, resty, err := customGetNode(path, node.UserName, node.Password, true)

	if err != nil || response == nil {
		if resty != nil {
			log.Printf("[DEBUG] Retrieved error response %s", resty.String())
		}
		return false, err
	}

	for _, role := range response.Roles {
		if strings.ToUpper(role) == "STANDALONE" {
			return true, err
		}
	}
	return false, err
}

func (node Node) AppServerIsRunning() (bool, error) {
	path := fmt.Sprintf("https://%s/ers/config/op/systemconfig/iseversion", node.Ip)
	_, resty, err := customGet(path, node.UserName, node.Password, false)

	if err != nil {
		if resty != nil {
			log.Printf("[DEBUG] Retrieved error response %s", resty.String())
		}
		return false, err
	}

	return true, err
}

func (node Node) ReturnIdOfCertificate() (*string, error) {
	path := fmt.Sprintf("https://%s/api/v1/certs/system-certificate/%s", node.Ip, node.HostName)
	log.Printf("[DEBUG] ImportCertificateIntoPrimary 2 %s", path)
	_, resty, err := customGetCerts(path, node.UserName, node.Password, false)

	if err != nil {
		if resty != nil {
			log.Printf("[DEBUG] Retrieved error response %s", resty.String())
		}
		return nil, err
	}
	log.Printf("[DEBUG] ImportCertificateIntoPrimary 3 %s", resty.String())
	response := resty.Result().(*isegosdk.ResponseCertificatesGetSystemCertificates)
	log.Printf("[DEBUG] ImportCertificateIntoPrimary 4 %s", response)
	if response.Response == nil {
		return nil, err
	}

	items := *response.Response
	// This endpoint has pagination, may be included or not?
	for _, item := range items {
		if item.FriendlyName == "Default self-signed server certificate" {
			return &item.ID, err
		}
	}

	return nil, err
}

func (node Node) RegisterToPrimary(primary Node) error {
	path := fmt.Sprintf("https://%s/api/v1/deployment/node", primary.Ip)
	allow := true
	request := isegosdk.RequestNodeDeploymentRegisterNode{
		Fqdn:            node.Fqdn,
		UserName:        node.UserName,
		Password:        node.Password,
		AllowCertImport: &allow,
		Roles:           node.Roles,
		Services:        node.Services,
	}
	_, err := customPost(path, node.UserName, node.Password, request)

	return err
}

func (node Node) UpdateRolesServices() error {
	path := fmt.Sprintf("https://%s/api/v1/deployment/node", node.Ip)
	request := isegosdk.RequestNodeDeploymentRegisterNode{}
	request.Roles = node.Roles
	request.Services = node.Services
	_, err := customPut(path, node.UserName, node.HostName, request)

	return err
}

func (node Node) ImportCertificateIntoPrimary(primary Node) error {
	log.Printf("[DEBUG] ImportCertificateIntoPrimary 1")
	certId, err := node.ReturnIdOfCertificate()
	if err != nil && certId != nil {
		return err
	}
	exportRequest := isegosdk.RequestCertificatesExportSystemCert{
		ID:     *certId,
		Export: "CERTIFICATE",
	}

	path := fmt.Sprintf("https://%s/api/v1/certs/system-certificate/export", node.Ip)

	response, err := customPost(path, node.UserName, node.Password, exportRequest)

	fdownload := isegosdk.FileDownload{}
	if err != nil {
		return err
	}

	fdownload.FileData = response.Body()
	zf, err := zip.NewReader(bytes.NewReader(fdownload.FileData), int64(len(fdownload.FileData)))
	if err != nil {
		return err
	}
	certData, err := zf.Open("Defaultselfsignedservercerti.pem")
	if err != nil {
		return err
	}
	certDataBytes, err := ioutil.ReadAll(certData)
	if err != nil {
		return err
	}
	defer certData.Close()
	certDataBytesDecoded := decodeUtf8(certDataBytes)

	allowBasicConstraintCAFalse := true
	allowOutOfDateCert := false
	allowSHA1Certificates := true
	trustForCertificateBasedAdminAuth := true
	trustForCiscoServicesAuth := true
	trustForClientAuth := true
	trustForIseAuth := true
	validateCertificateExtensions := true

	request := isegosdk.RequestCertificatesImportTrustCert{
		AllowBasicConstraintCaFalse:       &allowBasicConstraintCAFalse,
		AllowOutOfDateCert:                &allowOutOfDateCert,
		AllowSHA1Certificates:             &allowSHA1Certificates,
		TrustForCertificateBasedAdminAuth: &trustForCertificateBasedAdminAuth,
		TrustForCiscoServicesAuth:         &trustForCiscoServicesAuth,
		TrustForClientAuth:                &trustForClientAuth,
		Data:                              certDataBytesDecoded,
		TrustForIseAuth:                   &trustForIseAuth,
		Name:                              node.Name,
		ValidateCertificateExtensions:     &validateCertificateExtensions,
	}
	path = fmt.Sprintf("https://%s/api/v1/certs/trusted-certificate/import", primary.Ip)
	_, err = customPost(path, node.UserName, node.Password, request)
	if err != nil {
		return err
	}

	return err
}

func (node Node) PromoteToPrimary() error {
	path := fmt.Sprintf("https://%s/api/v1/deployment/primary", node.Ip)

	_, err := customPostWithNoBody(path, node.UserName, node.Password)
	if err != nil {
		if err.Error() == "error with operation customPost." {
			return fmt.Errorf("Could not update node to PRIMARY")
		}
	}
	return err
}

// *********************************************Util Funcs*******************************************************
func decodeUtf8(b []byte) string {
	result := ""
	for len(b) > 0 {
		r, size := utf8.DecodeLastRune(b)

		b = b[:len(b)-size]
		result = fmt.Sprintf("%c%s", r, result)
	}
	return result
}

// *********************************************API FUNCS********************************************************
func customGet(path string, username string, password string, castResult bool) (*Node, *resty.Response, error) {
	client := resty.New()
	client.SetDebug(true)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	client.SetBasicAuth(username, password)
	response, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Get(path)
	if err != nil {
		return nil, nil, err

	}
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation customGet. %s", response)
	}

	return nil, response, err
}

func customGetCerts(path string, username string, password string, castResult bool) (*Node, *resty.Response, error) {
	client := resty.New()
	client.SetDebug(true)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	client.SetBasicAuth(username, password)
	response, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&isegosdk.ResponseCertificatesGetSystemCertificates{}).
		SetError(&Error).
		Get(path)
	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation customGet. %s", response)
	}

	return nil, response, err
}

func customGetNode(path string, username string, password string, castResult bool) (*Node, *resty.Response, error) {
	client := resty.New()
	client.SetDebug(true)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	client.SetBasicAuth(username, password)
	response, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&NodeR{}).
		SetError(&Error).
		Get(path)
	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation customGet. %s", response)
	}

	var result *NodeR
	if castResult {
		result = response.Result().(*NodeR)
	}
	return &result.Response, response, err
}

func customPost(path string, username string, password string, requestBody interface{}) (*resty.Response, error) {
	client := resty.New()
	client.SetDebug(true)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	client.SetBasicAuth(username, password)
	response, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestBody).
		SetError(&Error).
		Post(path)
	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation customPost.")
	}
	return response, err
}

func customPostWithNoBody(path string, username string, password string) (*resty.Response, error) {
	client := resty.New()
	client.SetDebug(true)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	client.SetBasicAuth(username, password)
	response, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Post(path)
	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation customPost.")
	}
	return response, err
}

func customPut(path string, username string, password string, requestBody interface{}) (*resty.Response, error) {
	client := resty.New()
	client.SetDebug(true)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	client.SetBasicAuth(username, password)
	response, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestBody).
		SetError(&Error).
		Put(path)
	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation customPut.")
	}
	return response, err
}
