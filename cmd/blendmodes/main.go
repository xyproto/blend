package main

import (
	"fmt"

	"github.com/xyproto/crossfade"
	"github.com/xyproto/imagelib"
)

var modes = map[string]crossfade.BlendFunc{
	"add":           crossfade.Add,
	"color":         crossfade.Color,
	"color_burn":    crossfade.ColorBurn,
	"color_dodge":   crossfade.ColorDodge,
	"darken":        crossfade.Darken,
	"darker_color":  crossfade.DarkerColor,
	"difference":    crossfade.Difference,
	"divide":        crossfade.Divide,
	"exclusion":     crossfade.Exclusion,
	"hard_light":    crossfade.HardLight,
	"hard_mix":      crossfade.HardMix,
	"hue":           crossfade.Hue,
	"lighten":       crossfade.Lighten,
	"lighter_color": crossfade.LighterColor,
	"linear_burn":   crossfade.LinearBurn,
	"linear_dodge":  crossfade.LinearDodge,
	"linear_light":  crossfade.LinearLight,
	"luminosity":    crossfade.Luminosity,
	"multiply":      crossfade.Multiply,
	"overlay":       crossfade.Overlay,
	"phoenix":       crossfade.Phoenix,
	"pin_light":     crossfade.PinLight,
	"reflex":        crossfade.Reflex,
	"saturation":    crossfade.Saturation,
	"screen":        crossfade.Screen,
	"soft_light":    crossfade.SoftLight,
	"substract":     crossfade.Substract,
	"vivid_light":   crossfade.VividLight,
}

func main() {
	dst, err := imagelib.Read("destination.jpg")
	if err != nil {
		panic(err)
	}

	src, err := imagelib.Read("source.jpg")
	if err != nil {
		panic(err)
	}

	fmt.Println("This program tests all the color blending modes in the library.")
	for name, mode := range modes {
		fmt.Println("Blending Mode: ", name)
		img := crossfade.BlendNewImage(dst, src, mode)
		err = imagelib.Write("blend_"+name+".jpg", img)
		if err != nil {
			panic(err)
		}
	}
}
