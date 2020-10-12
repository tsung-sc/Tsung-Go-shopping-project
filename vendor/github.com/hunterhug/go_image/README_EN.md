# Simple Golang Picture transformation lib |  Thumbnail | Scale

## What

1. Very hard to find such lib so I make this.
2. Scale the size of Picture such the format of jpg or png.


## New

- 2019/5/12 Support Bytes Stream of Picture to Scale.

## How

### Install

```
go get -v -u github.com/hunterhug/go_image
```

### Main funciton

- Scale by width，input and output is bytes:

```
func ScaleB2B(InRaw []byte, width int) (OutRaw []byte, err error)
```

- Scale by width，input and output is the location(filename) of picture:

```
func ScaleF2F(filename string, savepath string, width int) (err error)
```

- Scale by width and height，input and output is bytes:

```
func ThumbnailB2B(InRaw []byte, width int, height int) (OutRaw []byte, err error)
```

- Scale by width and height，input and output is the location(filename) of picture:

```
func ThumbnailF2F(filename string, savepath string, width int, height int) (err error)
```

- Check the image real file type, such 4.jpg return 4.png(input is the location of file):

```
func RealImageName(filename string) (filerealname string, err error)
```

- Rename the file,if file exist will throw err if force is false, when force is true overwrite:

```
func ChangeImageName(oldname string, newname string, force bool) (err error) 
```

## Example

### example_test.go

```
package go_image

import (
	"fmt"
	"testing"
)

//将某一图片文件进行缩放后存入另外的文件中
func TestImage(t *testing.T) {
	//打印当前文件夹位置
	fmt.Printf("本文件文件夹位置:%s\n", CurDir())

	//图像位置
	filename := "./testdata/gopher.png"

	//宽度,高度
	width := 500
	height := 800

	//保存位置
	save1 := "./testdata/gopher500.jpg"
	save2 := "./testdata/gopher500_800.png"

	//按照宽度进行等比例缩放
	err := ScaleF2F(filename, save1, width)
	if err != nil {
		fmt.Printf("生成按宽度缩放图失败：%s\n", err.Error())
	} else {
		fmt.Printf("生成按宽度缩放图：%s\n", save1)
	}

	//按照宽度和高度进行等比例缩放
	err = ThumbnailF2F(filename, save2, width, height)
	if err != nil {
		fmt.Printf("生成按宽度高度缩放图:%s\n", err.Error())
	} else {
		fmt.Printf("生成按宽度高度缩放图:%s\n", save2)
	}

	//查看图像文件的真正名字
	//如 ./testdata/gopher500.jpg其实是png类型,但是命名错误,需要纠正!
	realfilename, err := RealImageName(save1)
	if err != nil {
		fmt.Printf("真正的文件名: %s->? err:%s\n", save1, err.Error())
	} else {
		fmt.Printf("真正的文件名:%s->%s\n", save1, realfilename)
	}

	//文件改名,强制性
	err = ChangeImageName(save1, realfilename, true)
	if err != nil {
		fmt.Printf("文件改名失败:%s->%s,%s\n", save1, realfilename, err.Error())
	} else {
		fmt.Println("改名成功")
	}

	//文件改名,不强制性
	err = ChangeImageName(save1, realfilename, false)
	if err != nil {
		fmt.Printf("文件改名失败:%s->%s,%s\n", save1, realfilename, err.Error())
	}
}

```

### Result

```
=== RUN   TestImage
本文件文件夹位置:E:\gocode\src\github.com\hunterhug\go_image
生成按宽度缩放图：./testdata/gopher500.jpg
生成按宽度高度缩放图:./testdata/gopher500_800.png
真正的文件名:./testdata/gopher500.jpg->./testdata/gopher500.png
改名成功
文件改名失败:./testdata/gopher500.jpg->./testdata/gopher500.png,File already exist error
--- PASS: TestImage (1.66s)
PASS
```

Origin:

![/testdata/gopher.png](/testdata/gopher.png)


Width 500px Scale:


![testdata/gopher500.png](testdata/gopher500.png)

Width 500px,Height 800px Scale:

![/testdata/gopher500_800.png](/testdata/gopher500_800.png)

## Come from

This is a Graphics library for the Go programming language.

Unless otherwise noted, the graphics-go source files are distributed
under the BSD-style license found in the LICENSE file.

Contributions should follow the same procedure as for the Go project:
http://golang.org/doc/contribute.html

