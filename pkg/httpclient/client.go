package httpclient

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type HTTPClient struct {
	config HTTPConfig
	client *http.Client
}

type HTTPConfig struct {
	Endpoint string
	NodeID   string
	RoundID  string
}

func ProvideHTTPClient(config HTTPConfig) HTTPClient {
	return HTTPClient{
		config: config,
		client: &http.Client{},
	}
}

func (h *HTTPClient) GuessAnswer(answer string) (bool, error) {
	endpoint := fmt.Sprintf("%s/submit-answer/%s", h.config.Endpoint, answer)
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return false, err
	}
	req.Header.Add("X-NODE-ID", h.config.NodeID)
	req.Header.Add("X-ROUND-ID", h.config.RoundID)

	res, err := h.client.Do(req)
	if err != nil {
		return false, err
	}

	if res.StatusCode == http.StatusOK {
		return true, nil
	}

	resultBody, err := io.ReadAll(res.Body)
	logrus.Error("Read Error ", err)
	if err != nil {
		return false, err
	}
	err = res.Body.Close()
	if err != nil {
		return false, err
	}

	logrus.Error("Server Response ", string(resultBody))
	return false, errors.New("response incorrect")
}

func (h *HTTPClient) CheckRound() (bool, error) {
	endpoint := fmt.Sprintf("%s/round-has-winner", h.config.Endpoint)
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		logrus.Error("Create Request Failed: ", err)
		return false, err
	}
	req.Header.Add("X-NODE-ID", h.config.NodeID)
	req.Header.Add("X-ROUND-ID", h.config.RoundID)

	res, err := h.client.Do(req)
	if err != nil {
		logrus.Error("Request Failed: ", err)
		return false, err
	}

	return res.StatusCode == http.StatusOK, nil
}
