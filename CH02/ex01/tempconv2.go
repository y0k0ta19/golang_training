// Author: "Shun Yokota"
// Copyright © 2016 RICOH Co, Ltd. All rights reserved

// Package tempconv2 performs Celsius, Fahrenheit and Kelvin conversions
package tempconv2

import "fmt"

//Celsius is a type of celsius
type Celsius float64

//Fahrenheit is a type of fahrenheit
type Fahrenheit float64

//Kelvin is is a type of kelvin
type Kelvin float64

//AbsoluteZeroC is absolute zero in celsius
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) }
