package twoorderedslicepicker

import (
	// standard libs only above!

	"github.com/litesoft-go/utils/enums"
)

type Pick struct {
	enums.Enum
}

var (
	Done       = Pick{Enum: enums.New("Done")}
	Left       = Pick{Enum: enums.New("Left")}
	Right      = Pick{Enum: enums.New("Right")}
	Equivalent = Pick{Enum: enums.New("Equivalent")}
)

func init() {
	enums.Add(&Done, &Left, &Right, &Equivalent)
}

func (p *Pick) UpdateFrom(found enums.IEnum) {
	_ = found.(*Pick)
}

func (p *Pick) UnmarshalJSON(data []byte) error {
	return enums.UnmarshalJSON(p, data) // p is Dst
}

type Pickable interface {
	AsOrderable() *string
}

func Which(left, right Pickable) Pick {
	lpValue := left.AsOrderable()
	rpValue := right.AsOrderable()
	noLeft := lpValue == nil
	noRight := rpValue == nil
	if noLeft && noRight {
		return Done
	}
	if noLeft {
		return Right
	}
	if noRight {
		return Left
	}
	lValue := *lpValue
	rValue := *rpValue
	if lValue == rValue {
		return Equivalent
	}
	if lValue < rValue {
		return Left
	}
	return Right
}
