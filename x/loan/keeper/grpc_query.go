package keeper

import (
	"github.com/anpan2val/loan/x/loan/types"
)

var _ types.QueryServer = Keeper{}
