package imagex_test

import (
	"bytes"
	"github.com/gophero/goal/errorx"
	"github.com/gophero/goal/imagex"
	"image/jpeg"
	"os"
	"testing"
)

var pic = "/Users/sam/Pictures/壁纸/唯美/53733594913dc.jpg"
var dpic = "/Users/sam/Pictures/壁纸/唯美/53733594913dc_1.jpg"
var dpic2 = "/Users/sam/Pictures/壁纸/唯美/53733594913dc_2.jpg"

func TestResize(t *testing.T) {
	f, err := os.Open(pic)
	errorx.Throw(err)
	img, err := jpeg.Decode(f)
	errorx.Throw(err)
	dimg := imagex.Resize(200, 200, img)

	errorx.Throw(err)
	var buf bytes.Buffer
	err = jpeg.Encode(&buf, dimg, &jpeg.Options{Quality: 100})
	err = os.WriteFile(dpic, buf.Bytes(), os.ModePerm)
	errorx.Throw(err)
}

func TestThumbnail(t *testing.T) {
	f, err := os.Open(pic)
	errorx.Throw(err)
	img, err := jpeg.Decode(f)
	errorx.Throw(err)
	dimg := imagex.Thumbnail(400, 400, img)

	errorx.Throw(err)
	var buf bytes.Buffer
	err = jpeg.Encode(&buf, dimg, &jpeg.Options{Quality: 100})
	err = os.WriteFile(dpic2, buf.Bytes(), os.ModePerm)
	errorx.Throw(err)
}
