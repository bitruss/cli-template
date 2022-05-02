package captcha

import (
	"bytes"
	"encoding/base64"
	"math"
	"math/rand"
	"strconv"
	"time"

	svg "github.com/ajstarks/svgo"
)

func Gen_svg_base64_prefix(width int, height int, color string) (string, string) {

	var result int

	a := rand.Intn(100)
	b := rand.Intn(10)
	c := rand.Intn(10)

	buf := new(bytes.Buffer)
	canvas := svg.New(buf)
	canvas.Start(width, height)
	font_size := int(math.Ceil(float64(height) * 0.6))
	font_style := "dominant-baseline:middle;text-anchor:middle;font-size:" + strconv.Itoa(font_size) + "px;fill:" + color

	switch rand.Intn(3) {
	case 1:
		result = a + b*c
		canvas.Text(width/2, height/2, strconv.Itoa(a)+" + "+strconv.Itoa(b)+"*"+strconv.Itoa(c), font_style)
	case 2:
		result = a + b + c
		canvas.Text(width/2, height/2, strconv.Itoa(a)+" + "+strconv.Itoa(b)+" + "+strconv.Itoa(c), font_style)
	default:
		result = b*c + a
		canvas.Text(width/2, height/2, strconv.Itoa(b)+"*"+strconv.Itoa(c)+" + "+strconv.Itoa(a), font_style)
	}
	canvas.End()
	return strconv.Itoa(result), "data:image/svg+xml;base64," + base64.StdEncoding.EncodeToString(buf.Bytes())

}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}
