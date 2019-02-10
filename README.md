# mix

Go module for blending/mixing two image files.

Fork of [phrozen/blend](https://github.com/phrozen/blend).

Example use:

    mix.Files("a.png", "b.png", "out.png", 0.5)

The last argument is a float that regulates the transition from one image to the other, where 0.0 is only "a.png", while 1.0 is only "b.png".

## General info

* License: MIT
* Version: 0.0.0
