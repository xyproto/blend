# crossfade

* Go module for crossfading between two image files.
* Started out as a fork of [phrozen/blend](https://github.com/phrozen/blend).
* Includes a `blend` command line utility, for mixing two images 50%/50%.

[github.com/anthonysimon/bild](https://github.com/anthonynsimon/bild) is a more popular choice than this package.

## Example use of the Go package

    crossfade.Files("a.png", "b.png", "out.png", 0.5)

The last argument is a float that regulates the transition from one image to the other, where 0.0 is only `"a.png"`, while 1.0 is only `"b.png"`.

## Screenshots

0% lemur 100% mountain

![0% lemur](img/lagginhorn.jpg)

20% lemur 80% mountain

![20% lemur](img/out80.png)

50% lemur 50% mountain

![50% lemur](img/out50.png)

80% lemur 20% mountain

![80% lemur](img/out20.png)

100% lemur 0% mountain

![lemur](img/lemur.jpg)

The images are from wikipedia: <a href="https://en.wikipedia.org/wiki/File:Eulemur_mongoz_(male_-_face).jpg">lemur</a> | [mountain](https://nn.wikipedia.org/wiki/Fil:Lagginhorn_west_face.jpg)

## Blendmode examples

| A                                    | B                                              |
|--------------------------------------|------------------------------------------------|
| ![source](cmd/blendmodes/source.jpg) | ![destination](cmd/blendmodes/destination.jpg) |

![blend add](img/blend_add.jpg)

Add

![blend color](img/blend_color.jpg)

Color

![blend color_burn](img/blend_color_burn.jpg)

Color Burn

![blend color_dodge](img/blend_color_dodge.jpg)

Color Dodge

![blend darken](img/blend_darken.jpg)

Darken

![blend darker_color](img/blend_darker_color.jpg)

Darker Color

![blend difference](img/blend_difference.jpg)

Difference

![blend divide](img/blend_divide.jpg)

Divide

![blend exclusion](img/blend_exclusion.jpg)

Exclusion

![blend hard_light](img/blend_hard_light.jpg)

Hard Light

![blend hard_mix](img/blend_hard_mix.jpg)

Hard Mix

![blend hue](img/blend_hue.jpg)

Hue

![blend lighten](img/blend_lighten.jpg)

Lighten

![blend lighter_color](img/blend_lighter_color.jpg)

Lighter Color

![blend linear_burn](img/blend_linear_burn.jpg)

Linear Burn

![blend linear_dodge](img/blend_linear_dodge.jpg)

Linear Dodge

![blend linear_light](img/blend_linear_light.jpg)

Linear Light

![blend luminosity](img/blend_luminosity.jpg)

Luminosity

![blend multiply](img/blend_multiply.jpg)

Multiply

![blend overlay](img/blend_overlay.jpg)

Overlay

![blend phoenix](img/blend_phoenix.jpg)

Phoenix

![blend pin_light](img/blend_pin_light.jpg)

Pin Light

![blend reflex](img/blend_reflex.jpg)

Reflex

![blend saturation](img/blend_saturation.jpg)

Saturation

![blend screen](img/blend_screen.jpg)

Screen

![blend soft_light](img/blend_soft_light.jpg)

Soft Light

![blend substract](img/blend_substract.jpg)

Substract

![blend vivid_light](img/blend_vivid_light.jpg)

Vivid Light

Note that these images and blend modes are from the [phrozen/blend](github.com/phrozen/blend) repo that this repo was forked from.

## General info

* License: MIT
* Version: 1.1.2
