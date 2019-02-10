package main

import (
	"fmt"
	"github.com/xyproto/crossfade"
	"github.com/xyproto/imagelib"
)

var modes = map[string]blend.BlendFunc{
	"add":           blend.Add,
	"color":         blend.Color,
	"color_burn":    blend.ColorBurn,
	"color_dodge":   blend.ColorDodge,
	"darken":        blend.Darken,
	"darker_color":  blend.DarkerColor,
	"difference":    blend.Difference,
	"divide":        blend.Divide,
	"exclusion":     blend.Exclusion,
	"hard_light":    blend.HardLight,
	"hard_mix":      blend.HardMix,
	"hue":           blend.Hue,
	"lighten":       blend.Lighten,
	"lighter_color": blend.LighterColor,
	"linear_burn":   blend.LinearBurn,
	"linear_dodge":  blend.LinearDodge,
	"linear_light":  blend.LinearLight,
	"luminosity":    blend.Luminosity,
	"multiply":      blend.Multiply,
	"overlay":       blend.Overlay,
	"phoenix":       blend.Phoenix,
	"pin_light":     blend.PinLight,
	"reflex":        blend.Reflex,
	"saturation":    blend.Saturation,
	"screen":        blend.Screen,
	"soft_light":    blend.SoftLight,
	"substract":     blend.Substract,
	"vivid_light":   blend.VividLight,
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
		img := mix.BlendNewImage(dst, src, mode)
		err = imagelib.Write("blend_"+name, img)
		if err != nil {
			panic(err)
		}
	}
}
