package blend

import (
	"github.com/xyproto/imagelib"
)

func MixFiles(inFilename1, inFilename2, outFilename string, ratio float64) error {
	img1, err := imagelib.Read(inFilename1)
	if err != nil {
		return err
	}
	img2, err := imagelib.Read(inFilename2)
	if err != nil {
		return err
	}

	// Crossfade
	SMul = ratio * 2.0
	DMul = -ratio*2.0 + 2.0
	outImage := BlendNewImage(img1, img2, Mix)

	err = imagelib.Write(outFilename, outImage)
	if err != nil {
		return err
	}
	return nil
}

func OverlayMixFiles(inFilename1, inFilename2, outFilename string, ratio float64) error {
	img1, err := imagelib.Read(inFilename1)
	if err != nil {
		return err
	}
	img2, err := imagelib.Read(inFilename2)
	if err != nil {
		return err
	}

	// Crossfade
	SMul = ratio * 2.0
	DMul = -ratio*2.0 + 2.0
	outImage := BlendNewImage(img1, img2, OverlayMix)

	err = imagelib.Write(outFilename, outImage)
	if err != nil {
		return err
	}
	return nil
}
