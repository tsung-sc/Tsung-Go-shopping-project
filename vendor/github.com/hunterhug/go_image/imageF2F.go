package go_image

import (
	"bytes"
	//"errors"
	"github.com/hunterhug/go_image/graphics"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

//按宽度和高度进行比例缩放，输入和输出都是图片字节数组
func ThumbnailB2B(InRaw []byte, width int, height int) (OutRaw []byte, err error) {
	src, filetype, err := LoadImage(bytes.NewReader(InRaw))
	if err != nil {
		return
	}
	dst := image.NewRGBA(image.Rect(0, 0, width, height))
	err = graphics.Thumbnail(dst, src)
	if err != nil {
		return
	}

	w := new(bytes.Buffer)
	err = SaveImage(dst, filetype, w)
	if err != nil {
		return
	}

	OutRaw = w.Bytes()
	return
}

// 按宽度和高度进行比例缩放，输入和输出都是文件
func ThumbnailF2F(filename string, savepath string, width int, height int) (err error) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	out, err := ThumbnailB2B(raw, width, height)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(savepath, out, 0777)
	if err != nil {
		return err
	}

	return
}

//按宽度进行比例缩放，输入和输出都是图片字节数组
func ScaleB2B(InRaw []byte, width int) (OutRaw []byte, err error) {
	img, filetype, err := Scale(InRaw, width)
	if err != nil {
		return
	}

	buffer := new(bytes.Buffer)
	err = SaveImage(img, filetype, buffer)
	if err != nil {
		return
	}

	OutRaw = buffer.Bytes()
	return
}

//按宽度进行比例缩放，输入输出都是文件
func ScaleF2F(filename string, savepath string, width int) (err error) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	out, err := ScaleB2B(raw, width)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(savepath, out, 0777)
	if err != nil {
		return err
	}

	return
}

//图像文件的真正名字
func RealImageName(filename string) (filerealname string, err error) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	_, ext, err := LoadImage(bytes.NewReader(raw))
	if err != nil {
		return
	}
	temp := strings.Split(filename, ".")
	if len(temp) < 2 {
		err = FileNameError
	}
	temp[len(temp)-1] = ext
	filerealname = strings.Join(temp, ".")
	return
}

//文件改名,如果force为假,且新的文件名已经存在,那么抛出错误
func ChangeImageName(oldname string, newname string, force bool) (err error) {
	if !force {
		_, err = os.Open(newname)
		if err == nil {
			err = FileExistError
			return
		}
	}
	err = os.Rename(oldname, newname)
	return

}

// 获取调用者的当前文件DIR
func CurDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return filepath.Dir(filename)
}

// 根据文件名打开图片,并编码,返回编码对象和文件类型
func LoadImage(r io.Reader) (img image.Image, filetype string, err error) {
	img, filetype, err = image.Decode(r)
	if err != nil {
	}
	return
}

// 将编码对象进行处理后返回字节数组
func SaveImage(img *image.RGBA, filetype string, w io.Writer) (err error) {
	if filetype == "png" {
		err = png.Encode(w, img)
	} else if filetype == "jpeg" {
		err = jpeg.Encode(w, img, nil)
	} else {
		err = ExtNotSupportError
	}
	return
}

//对图片字节数组等比例变化,宽度为newdx,返回图像编码和文件类型
func Scale(raw []byte, newdx int) (dst *image.RGBA, filetype string, err error) {
	src, filetype, err := LoadImage(bytes.NewReader(raw))
	if err != nil {
		return
	}
	bound := src.Bounds()
	dx := bound.Dx()
	dy := bound.Dy()
	dst = image.NewRGBA(image.Rect(0, 0, newdx, newdx*dy/dx))
	// 产生缩略图,等比例缩放
	err = graphics.Scale(dst, src)
	return
}
