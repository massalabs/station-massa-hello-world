package plugin

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

const ThyraRegisterEndpoint = "http://my.massa/plugin-manager/register"

type registerBody struct {
	ID          int64
	Name        string
	Author      string
	Description string
	Logo        string
	URL         string
	APISpec     string
}

func Register(
	pluginID int64,
	name string, author string,
	shortDescription string,
	socket net.Addr, spec string,
	logo []byte) error {

	reg := registerBody{
		ID:          pluginID,
		Name:        name,
		Author:      author,
		Description: shortDescription,
		URL:         "http://" + socket.String(),
		APISpec:     spec,
		Logo:        base64.StdEncoding.EncodeToString(logo),
	}

	body, err := json.Marshal(reg)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", ThyraRegisterEndpoint, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusNoContent {
		body, _ := ioutil.ReadAll(response.Body)
		defer response.Body.Close()

		return fmt.Errorf("plugin registry failed: HTTP status: %d, HTTP body: %v", response.StatusCode, body)
	}

	return nil
}
