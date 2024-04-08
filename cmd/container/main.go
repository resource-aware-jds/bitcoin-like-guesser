package main

import (
	"bitcoin-like-guesser/pkg/httpclient"
	"bitcoin-like-guesser/pkg/logic"
	"encoding/json"
	containerlib "github.com/resource-aware-jds/container-lib"
	"github.com/resource-aware-jds/container-lib/model"
	"github.com/resource-aware-jds/container-lib/pkg/containerlibcontext"
	"github.com/sirupsen/logrus"
	"strconv"
)

type TaskData struct {
	Endpoint     string `json:"endpoint"`
	ExpectedHash string `json:"expectedHash"`
	RoundID      string `json:"roundID"`
	NodeID       string `json:"nodeID"`
	Start        int    `json:"start"`
	End          int    `json:"end"`
}

func main() {
	containerlib.Run(func(ctx containerlibcontext.Context, task model.Task) error {
		var data TaskData
		err := json.Unmarshal(task.Attributes, &data)
		if err != nil {
			logrus.Error(err)
			return err
		}

		client := httpclient.ProvideHTTPClient(httpclient.HTTPConfig{
			Endpoint: data.Endpoint,
			NodeID:   data.NodeID,
			RoundID:  data.RoundID,
		})

		isWin := false
		for i := data.Start; i < data.End; i++ {
			iStr := strconv.Itoa(i)
			result := logic.GenerateBase64Hash(iStr)

			if result == data.ExpectedHash {
				// Guess the answer
				res, err := client.GuessAnswer(iStr)
				if err != nil {
					logrus.Error("Guess Answer incorrect!: ", err)
					continue
				}
				isWin = res
				break
			}

			if i%1000 != 0 {
				continue
			}
			hasWinner, err := client.CheckRound()
			if err != nil {
				logrus.Error("Check round failed", err)
			}
			if hasWinner {
				logrus.Info("Already got the winner for this round. Breaking!")
				break
			}
		}

		logrus.Info("IsWin!: ", isWin)
		ctx.Success()
		return nil
	})
}
