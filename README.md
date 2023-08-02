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
add

![blend color](img/blend_color.jpg)
color

![blend color_burn](img/blend_color_burn.jpg)
color_burn

![blend color_dodge](img/blend_color_dodge.jpg)
color_dodge

![blend darken](img/blend_darken.jpg)
darken

![blend darker_color](img/blend_darker_color.jpg)
darker_color

![blend difference](img/blend_difference.jpg)
difference

![blend divide](img/blend_divide.jpg)
divide

![blend exclusion](img/blend_exclusion.jpg)
exclusion

![blend hard_light](img/blend_hard_light.jpg)
hard_light

![blend hard_mix](img/blend_hard_mix.jpg)
hard_mix

![blend hue](img/blend_hue.jpg)
hue

![blend lighten](img/blend_lighten.jpg)
lighten

![blend lighter_color](img/blend_lighter_color.jpg)
lighter_color

![blend linear_burn](img/blend_linear_burn.jpg)
linear_burn

![blend linear_dodge](img/blend_linear_dodge.jpg)
linear_dodge

![blend linear_light](img/blend_linear_light.jpg)
linear_light

![blend luminosity](img/blend_luminosity.jpg)
luminosity

![blend multiply](img/blend_multiply.jpg)
multiply

![blend overlay](img/blend_overlay.jpg)
overlay

![blend phoenix](img/blend_phoenix.jpg)
phoenix

![blend pin_light](img/blend_pin_light.jpg)
pin_light

![blend reflex](img/blend_reflex.jpg)
reflex

![blend saturation](img/blend_saturation.jpg)
saturation

![blend screen](img/blend_screen.jpg)
screen

![blend soft_light](img/blend_soft_light.jpg)
soft_light

![blend substract](img/blend_substract.jpg)
substract

![blend vivid_light](img/blend_vivid_light.jpg)
vivid_light

## General info

* License: BSD-3
* Version: 1.1.1
