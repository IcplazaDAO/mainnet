package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// farm module sentinel errors
var (
	ErrExpiredHeight      = sdkerrors.Register(ModuleName, 2, "expired block height")
	ErrInvalidLPToken     = sdkerrors.Register(ModuleName, 3, "invalid lp token denom")
	ErrNotMatch           = sdkerrors.Register(ModuleName, 4, "the data does not match")
	ErrPoolExpired        = sdkerrors.Register(ModuleName, 5, "the farm pool has expired")
	ErrPoolNotStart       = sdkerrors.Register(ModuleName, 6, "the farm pool don't start")
	ErrPoolExist          = sdkerrors.Register(ModuleName, 7, "the farm pool already exist")
	ErrPoolNotFound       = sdkerrors.Register(ModuleName, 8, "the farm pool does not exist")
	ErrInvalidOperate     = sdkerrors.Register(ModuleName, 9, "invalid operate")
	ErrFarmerNotFound     = sdkerrors.Register(ModuleName, 10, "the farmer does not exist")
	ErrInvalidPoolName    = sdkerrors.Register(ModuleName, 11, "invalid pool name, must contain alphanumeric characters, _ and - only length greater than 0 and less than or equal to 70 and start with an english letter")
	ErrInvalidDescription = sdkerrors.Register(ModuleName, 12, "invalid pool description, length must be less than or equal to 280")
	ErrInvalidAppend      = sdkerrors.Register(ModuleName, 13, "cannot add new token as a reward")
	ErrInvalidRewardRule  = sdkerrors.Register(ModuleName, 14, "invalid reward rule")
	ErrAllEmpty           = sdkerrors.Register(ModuleName, 15, "shouldn't all be empty")
	ErrMinStaking         = sdkerrors.Register(ModuleName, 16, "staking amount should greater than minStaking")
	ErrInvalidParam       = sdkerrors.Register(ModuleName, 17, "invalid param")
)
