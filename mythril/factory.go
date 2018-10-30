package mythril

import (
	"errors"
	"github.com/JoranHonig/shard/mythril/alpha"
	"github.com/JoranHonig/shard/mythril/generic"
)

type MythrilServiceType int

const (
	ALPHA MythrilServiceType = 1 << iota
	V1
	V2
)

func BuildMythrilService(version MythrilServiceType, apiKey string) (generic.MythrilService, error) {
	switch version {
	case ALPHA:
		return alpha.BuildMythrilServiceALPHA(apiKey), nil
	default:
		return nil, errors.New("Invalid MythrilServiceType")
	}
}
