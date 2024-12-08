// Реализовать паттерн «адаптер» на любом примере.

package main

import "fmt"

type USBC struct{}

// existing class
type LighteningCable struct{}

func (ltCable *LighteningCable) ChargeLightening() string {
	return "beep beep charging"
}

// needed interface
type PhoneCharger interface {
	ChargeWithNormalCable() string
}

// target function
func chargePhone(pc PhoneCharger) {
	fmt.Println(pc.ChargeWithNormalCable())
}

// adapter
type ChargerAdapter struct {
	ltCable *LighteningCable
}

func (chAdapter *ChargerAdapter) ChargeWithNormalCable() string {
	return chAdapter.ltCable.ChargeLightening()
}

func main() {
	ltCable := &LighteningCable{}
	adapter := &ChargerAdapter{
		ltCable: ltCable,
	}

	chargePhone(adapter)

}
