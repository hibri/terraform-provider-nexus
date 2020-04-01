package nexus

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type NexusClient struct {
	UserName string
	Password string
	BaseUrl  string
	Insecure bool
}
var apiPath =  "/service/rest"
func  (n *NexusClient)  Get(m interface{}, endPoint string) error  {
	client, req, err := n.newRequest(http.MethodGet,endPoint, nil)
	if client != nil {
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		return  Parse(m, body)

	}
	return err
}

func (n *NexusClient) Update(m interface{}, endPoint string) (interface{}, error) {

	data, err := json.Marshal(m)
	fmt.Println(string(data))
	client, req, err := n.newRequest(http.MethodPut,endPoint, bytes.NewReader(data))

	if client != nil {
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		err = handleHttpError(resp)
		if err == nil {
			defer resp.Body.Close()
		}
		return m, err
	}
	return m, err
}

func handleHttpError(resp *http.Response) error {
	if resp.StatusCode >= 400 {
		body, _ := ioutil.ReadAll(resp.Body)
		return errors.New(fmt.Sprintf("HTTP Error \n Status : %d, Body : \n %s", resp.StatusCode, string(body)))
		defer resp.Body.Close()
	}
	return nil
}

func (n *NexusClient) newRequest(httpMethod string,endPoint string, body io.Reader) (*http.Client, *http.Request, error) {

	requestUrl := fmt.Sprintf("%s%s%s", n.BaseUrl, apiPath, endPoint)
	client := &http.Client{}
	req, err := http.NewRequest(httpMethod, requestUrl, body)
	req.Header.Set("Content-Type","application/json")
	n.setAuth(req)
	return client, req, err
}

func (n *NexusClient) setAuth(req *http.Request) {
	req.SetBasicAuth(n.UserName, n.Password)
}



func NewNexusClient(url string, userName string, password string) *NexusClient {
	return &NexusClient{
		BaseUrl:  url,
		UserName: userName,
		Password: password}
}
func Parse(m interface{}, jsonBody []byte) error {
	return  json.Unmarshal(jsonBody, m)

}