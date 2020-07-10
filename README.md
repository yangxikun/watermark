# watermark
This is a simple program for batch add watermark.

## 编译

go build

## 使用

watermark -h

```text
Usage of watermark:
  -image string
    	图片
  -scale int
    	控制水印之间的间距 (default 2)
  -size float
    	水印文本大小 (default 12)
  -text string
    	水印文本
```

## 示例

watermark -image 身份证正面.jpeg -text ziroom -scale 4 -size 32

输出添加水印的图片：watermark_身份证正面.jpeg