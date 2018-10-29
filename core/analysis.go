package core

import (
	"shard/mythril/generic"
	"time"
	"errors"
	"log"
	"github.com/sirupsen/logrus"
)

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
	AnalyzeRuntimeBytecode(bytecode string) ([]generic.Issue, error)
	AnalyzeBytecode(bytecode string) ([]generic.Issue, error)
	AnalyzeSourceCode(sourceCode string) ([]generic.Issue, error)
	AnalyzeContract(contract SolidityContract) ([]generic.Issue, error)
}


type BaseAnalysisService struct {
	MythrilService generic.MythrilService
}

func IsClosed(ch <-chan []generic.Issue) bool {
	select {
	case <-ch:
		return true
	default:
	}

	return false
}

func (b *BaseAnalysisService) AnalyzeRuntimeBytecode(bytecode string) ([]generic.Issue, error){
	resultChannel := make(chan []generic.Issue, 1)

	select {
	case <- time.After(10 * time.Second):
		return nil, errors.New("Timeout encountered in the analysis")
	case result := <- resultChannel:
		return result, nil
	}
}

func (b *BaseAnalysisService) AnalyzeBytecode(bytecode string) ([]generic.Issue, error) {

	resultChannel := make(chan []generic.Issue, 1)
	go func() {
		logrus.Debug("Submitting job to the mythril service")
		id, err := b.MythrilService.Submit(bytecode)

		if err != nil {
			log.Fatal(err)
		}

		previousStatus := ""
		for !IsClosed(resultChannel) {
			time.Sleep(1 * time.Second)

			logrus.Debug("Checking Status")
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
			case "Finished":
				res, err := b.MythrilService.GetIssueResult(*id)
				if err != nil {
					log.Fatal(err)
				}
				resultChannel <- res
				return
			case "Error":
				logrus.Info("Error encountered during analysis")
				resultChannel <- nil
			default:
				continue
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

func (b *BaseAnalysisService) AnalyzeSourceCode(sourceCode string) ([]generic.Issue, error){
	resultChannel := make(chan []generic.Issue, 1)

	select {
	//case <- time.After(10 * time.Second):
	//	return nil, errors.New("Timeout encountered in the analysis")
	case result := <- resultChannel:
		return result, nil
	}
}

func (b *BaseAnalysisService) AnalyzeContract(contract SolidityContract) ([]generic.Issue, error){
	resultChannel := make(chan []generic.Issue, 1)

	select {
	case <- time.After(10 * time.Second):
		return nil, errors.New("Timeout encountered in the analysis")
	case result := <- resultChannel:
		return result, nil
	}
}


