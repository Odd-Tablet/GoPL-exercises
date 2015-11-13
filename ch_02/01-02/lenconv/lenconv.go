package lenconv

import (
	"fmt"
)

type Foot float64
type Meter float64
type Mile float64
type Kilometer float64

func (k Kilometer) String() string {
	return fmt.Sprintf("%g Kilometers", k)
}
func (m Meter) String() string {
	return fmt.Sprintf("%g Meters", m)
}
func (m Mile) String() string {
	return fmt.Sprintf("%g Miles", m)
}
func (f Foot) String() string {
	return fmt.Sprintf("%g Feet", f)
}
func KToMi(k Kilometer) Mile {
	return Mile(k * .621)
}

func MeToF(m Meter) Foot {
	return Foot(m * 3.281)
}
func MiToK(m Mile) Kilometer {
	return Kilometer(m * 1.609)
}
func MiToF(m Mile) Foot {
	return Foot(m * 5280)
}
func FToMi(f Foot) Mile {
	return Mile(f * .000189)
}
func FToMe(f Foot) Meter {
	return Meter(f * .305)
}
func MeToK(m Meter) Kilometer {
	return Kilometer(m / 1000)
}
func KToMe(k Kilometer) Meter {
	return Meter(k * 1000)
}
