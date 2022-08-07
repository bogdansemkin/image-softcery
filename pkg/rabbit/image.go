package rabbit

import (
	"github.com/nfnt/resize"
	"github.com/sirupsen/logrus"
	"image/jpeg"
	"io/ioutil"
	"os"
)

func imageResize(path string) string {
	file, err := os.Open(path)
	if err != nil {
		logrus.Errorf("Error on opening file, %s", err)
	}
	img, err := jpeg.Decode(file)
	if err != nil {
		logrus.Fatal(err)
	}
	file.Close()

	m := resize.Resize(1280, 768, img, resize.Lanczos3)

	out, err := ioutil.TempFile("D:\\image-softcery\\templates\\img", "resize75-*.png")
	if err != nil {
		logrus.Fatal(err)
	}
	defer out.Close()

	jpeg.Encode(out, m, nil)

	return out.Name()
}

func imageHalfResize(path string) string {
	file, err := os.Open(path)
	if err != nil {
		logrus.Errorf("Error on opening file, %s", err)
	}
	img, err := jpeg.Decode(file)
	if err != nil {
		logrus.Fatal(err)
	}
	file.Close()

	m := resize.Resize(1024, 768, img, resize.Lanczos3)

	out, err := ioutil.TempFile("D:\\image-softcery\\templates\\img", "resize50-*.png")
	if err != nil {
		logrus.Fatal(err)
	}
	defer out.Close()

	jpeg.Encode(out, m, nil)

	return out.Name()
}

func imageFullResize(path string) string {
	file, err := os.Open(path)
	if err != nil {
		logrus.Errorf("Error on opening file, %s", err)
	}
	img, err := jpeg.Decode(file)
	if err != nil {
		logrus.Fatal(err)
	}
	file.Close()

	m := resize.Resize(800, 600, img, resize.Lanczos3)

	out, err := ioutil.TempFile("D:\\image-softcery\\templates\\img", "resize25-*.png")
	if err != nil {
		logrus.Fatal(err)
	}
	defer out.Close()

	jpeg.Encode(out, m, nil)

	return out.Name()
}
