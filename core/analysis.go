package core

import (
	"shard/mythril/generic"
	"time"
	"errors"
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

	select {
	case <- time.After(10 * time.Second):
		return nil, errors.New("Timeout encountered in the analysis")
	case result := <- resultChannel:
		return result, nil
	}
}

func (b *BaseAnalysisService) AnalyzeSourceCode(sourceCode string) ([]Issue, error){
	resultChannel := make(chan []Issue, 1)

	select {
	case <- time.After(10 * time.Second):
		return nil, errors.New("Timeout encountered in the analysis")
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