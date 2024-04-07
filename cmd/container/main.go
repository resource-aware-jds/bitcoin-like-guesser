package main

import (
	"bitcoin-like-guesser/pkg/logic"
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
		// TODO: Marshal task attribute into TaskData

		// TODO: Provide HTTPClient

		for i := 0; i < 1000000000; i++ {
			iStr := strconv.Itoa(i)
			result := logic.GenerateBase64Hash(iStr)

			// TODO: Use from TaskData
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
		}

		ctx.Success()
		return nil
	})
}
