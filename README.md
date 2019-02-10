# mix

Go module for blending/mixing two image files.

Fork of [phrozen/blend](https://github.com/phrozen/blend).

Example use:

    mix.Files("a.png", "b.png", "out.png", 0.5)

The last argument is a float that regulates the transition from one image to the other, where 0.0 is only "a.png", while 1.0 is only "b.png".

## Screenshots

0% lemur 100% mountain

![0% lemur](img/lagginhorn.jpg)

20% lemur 80% mountain

![20% lemur](img/out20.png)

50% lemur 50% mountain

![50% lemur](img/out50.png)

80% lemur 20% mountain

![80% lemur](img/out80.png)

100% lemur 0% mountain

![lemur](img/lemur.jpg)

## General info

* License: MIT
* Version: 0.0.0
