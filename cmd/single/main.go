package main

import (
	"bitcoin-like-guesser/config"
	"bitcoin-like-guesser/pkg/httpclient"
	"bitcoin-like-guesser/pkg/logic"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

func main() {
	// Try to find the hash from 0 to 1,000,000,000

	cfg := config.Load()

	client := httpclient.ProvideHTTPClient(httpclient.HTTPConfig{
		Endpoint: cfg.Endpoint,
		NodeID:   cfg.NodeID,
		RoundID:  cfg.RoundID,
	})

	isWin := false
	logrus.Info("Start to work on : ", time.Now().String())

	for i := 0; i < 1000000000; i++ {
		iStr := strconv.Itoa(i)
		result := logic.GenerateBase64Hash(iStr)

		if result == cfg.ExpectedHash {
			// Guess the answer
			res, err := client.GuessAnswer(iStr)
			if err != nil {
				logrus.Error("Guess Answer incorrect!: ", err)
				return
			}
			isWin = res
			break
		}

		if i%1000000 == 0 {
			logrus.Info("Finding in section: ", i, " - ", i+1000000)
		}
	}

	logrus.Info("IsWin!: ", isWin)

}
