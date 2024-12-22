# docx2md

<div style="text-align: center;">
中文
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
<a href="README_CN.md">English</a>
</div>

docx2md: 将docx文件转为markdown

本项目实现了将docx转化为Markdown

> Tips: 现有版本已经支持标题样式识别，但不会根据样式“级别”，而是根据样式的字体大小来识别级别。
> 字体大小与Heading级别对应请见下方详细介绍
> 因此即使样式名称为Heading 1，也不会将其处理为一级标题。

## 快速开始

### 安装
```shell
go get -u github.com/zakahan/docx2md
```

### 使用

```go
package main

import (
    "fmt"
    "github.com/zakahan/docx2md"
)

func main() {
    // the docx file path : examples/example_1.docx
    // the save dir: examples
    path, mdString, err := docx2md.DocxConvert("examples/example_1.docx", "examples")
    if err != nil {
        fmt.Println("error")
        return
    }
    fmt.Println(mdString)
    fmt.Println(path)
}

```

### 例子

当你运行上面的代码，会得到一个markdown文件路径
里面有一些图片

- 输出路径:
> examples\a61a6a30-99a5-4638-b3a3-dd93ee6228ec\example_1.md

- 文件结构

![1.png](images%2F1.jpg)

- 结果展示
    - `.docx` left
    - `.md` right

![2.jpg](images%2F2.jpg)

## Word2Heading

注：换算关系 word的字号*2 = 标记中w:fz的val = pt值

根据实际情况，大部分文档并未严格遵循目录（TOC）或大纲设置，因此本转换过程不依赖于TOC，而是完全依据段落的字体大小来确定标题级别。

由于Markdown的#、##等标题级别与HTML中的\<h1\>、\<h2\>标签相对应，而这些HTML标题标签的默认字体大小与Word文档中的FontSize存在对应关系。

在大多数浏览器和默认CSS设置下，HTML标题标签的默认字体大小如下：

- `<h1>`: 32px (约24pt)
- `<h2>`: 24px (约18pt)
- `<h3>`: 18.72px (约14pt)
- `<h4>`: 16px (约12pt)
- `<h5>`: 13.28px (约10pt)
- `<h6>`: 12px (约9pt)

考虑到正文通常使用的小四号字体（12pt），于是\<h4\>、\<h5\>和\<h6\>，因为三个标题级别已经足够区分文档结构。

### 区分正文与标题

为了更准确地区分正文与标题，特别是对于可能全部使用四号字体书写的内容，引入了长度限制：

- 如果段落长度小于或等于5个汉字（即len(x) <= 15），则该段落可被识别为\<h3\>标题。
- Word中`w:sz`属性的值是字体大小（pt）的两倍，所以我们将它乘以2后进行比较：
    - h1: 当`w:sz` >= 48
    - h2: 当36 <= `w:sz` < 48
    - h3: 当28 <= `w:sz` < 36 并且 len(x) <= 15

### 特殊情况处理

此外，如果段落以编号或特定标记（如“一、二、三”或“第xxxx”或带有Word自身的编号标记`w:numPr`）开头，那么我们可以放宽对\<h3\>的长度限制至15个汉字（= 45字符），并且允许出现\<h4\>级别的标题。


### 最新更新

- [x] 支持了样式设置标题识别，对于通过word样式设置的标题，目前已经可以实现识别功能。