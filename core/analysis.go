package core

import (
	"shard/mythril/generic"
	"time"
	"errors"
	"log"
	"github.com/sirupsen/logrus"
)


type Issue struct {

}

type SoliditySource struct {
	Filename string
	SourceCode string
}

type SolidityContract struct {
	BontractName string
	Bytecode string
	RuntimeBytecode string
	SourceMap string
	RuntimeSourceMap string
	Sources []SoliditySource
}

type AnalysisService interface {
	AnalyzeRuntimeBytecode(bytecode string) ([]Issue, error)
	AnalyzeBytecode(bytecode string) ([]Issue, error)
	AnalyzeSourceCode(sourceCode string) ([]Issue, error)
	AnalyzeContract(contract SolidityContract) ([]Issue, error)
}


type BaseAnalysisService struct {
	MythrilService generic.MythrilService
}

func IsClosed(ch <-chan []Issue) bool {
	select {
	case <-ch:
		return true
	default:
	}

	return false
}

func (b *BaseAnalysisService) AnalyzeRuntimeBytecode(bytecode string) ([]Issue, error){
	resultChannel := make(chan []Issue, 1)

	select {
	case <- time.After(10 * time.Second):
		return nil, errors.New("Timeout encountered in the analysis")
	case result := <- resultChannel:
		return result, nil
	}
}

func (b *BaseAnalysisService) AnalyzeBytecode(bytecode string) ([]Issue, error) {

	resultChannel := make(chan []Issue, 1)
	go func() {
		logrus.Info("Submitting job to the mythril service")
		id, err := b.MythrilService.Submit(bytecode)

		if err != nil {
			log.Fatal(err)
		}

		previousStatus := ""
		for !IsClosed(resultChannel) {
			time.Sleep(1 * time.Second)

			logrus.Info("Checking Status")
			s, err := b.MythrilService.CheckStatus(*id)
			if err != nil {
				logrus.Info(err)
				continue
			}
			if s.Status != previousStatus {
				logrus.Info("Analysis status for job changed to: ", s.Status)
				previousStatus = s.Status
			}
			switch s.Status {
			case "Done":
				resultChannel <- nil
			case "Error":
				logrus.Info("Error encountered during analysis")
				resultChannel <- nil
			default:
				resultChannel <- nil
			}
		}
	}()

	select {
	case <- time.After(10 * time.Second):
		close(resultChannel)
		return nil, errors.New("Timeout encountered in the analysis")
	case result := <- resultChannel:
		return result, nil
	}
}

func (b *BaseAnalysisService) AnalyzeSourceCode(sourceCode string) ([]Issue, error){
	resultChannel := make(chan []Issue, 1)

	select {
	//case <- time.After(10 * time.Second):
	//	return nil, errors.New("Timeout encountered in the analysis")
	case result := <- resultChannel:
		return result, nil
	}
}

func (b *BaseAnalysisService) AnalyzeContract(contract SolidityContract) ([]Issue, error){
	resultChannel := make(chan []Issue, 1)

	select {
	case <- time.After(10 * time.Second):
		return nil, errors.New("Timeout encountered in the analysis")
	case result := <- resultChannel:
		return result, nil
	}
}