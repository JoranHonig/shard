package core

import (
	"strings"
	"github.com/ethereum/go-ethereum/common/compiler"
	"errors"
)

// Compiles the contract at _filename
// _filename can also be of the form filename:ContractName
// In which case only the contract with ContractName will be considered
func Compile(filename string) ([]string, error) {
	parts := strings.Split(filename, ":")

	if len(parts) > 1 {
		contracts, err := compiler.CompileSolidity("", parts[0])
		if err != nil {
			return nil, err
		}

		contract, ok := contracts[filename]
		if !ok {
			return nil, errors.New("Wrong contract name provided")

		}
		return []string{contract.Code}, nil
	}

	contracts, err := compiler.CompileSolidity("", filename)
	if err != nil {
		return nil, err
	}

	result := make([]string, 0)
	for _, contract := range contracts {
		result = append(result, contract.Code)
	}

	return result, nil
}
