# docx2md

<div style="text-align: center;">
<a href="README_CN.md">中文</a>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 
English
</div>

docx2md: convert docx to markdown.

This project implements the conversion from DOCX to Markdown, achieved through XML matching.

>Tips: The current version already supports recognition of heading styles, but it identifies the level based on the font size of the style rather than the "level" name. 
> For a detailed correspondence between font sizes and Heading levels, please see the introduction below.
> Therefore, even if a style is named Heading 1, it will not be processed as a first-level heading unless it matches the font size designated for Heading 1.


## Quick Start

### Installation
```shell
go get -u github.com/zakahan/docx2md
```

### Usage

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


### Example

when you run this code, you will get a path of the markdown file
and some images file.
- The output path:
> examples\a61a6a30-99a5-4638-b3a3-dd93ee6228ec\example_1.md

- The struct of the result:

![1.png](images%2F1.jpg)

- The Result File
    - `.docx` left
    - `.md` right

![2.jpg](images%2F2.jpg)


## Word2Heading

Note: The conversion relationship is as follows: Word's font size * 2 = the val of w:fz in the markup = pt value.

According to practical conditions, most documents do not strictly follow the Table of Contents (TOC) or outline settings;
therefore, we will not rely on the TOC for this process. Instead, it will completely depend on the font size.


Since the header levels in Markdown (# ##) correspond to HTML's \<h1\>, \<h2\> tags, and \<h1\>... have a corresponding relationship with FontSize,
this method can be used for comparison.

In most browsers and with default CSS settings, the default font sizes for HTML heading tags are as follows:

- \<h1\>: 32px (approximately 24pt) 
- \<h2\>: 24px (approximately 18pt)
- \<h3\>: 18.72px (approximately 14pt)
- \<h4\>: 16px (approximately 12pt)
- \<h5\>: 13.28px (approximately 10pt)
- \<h6\>: 12px (approximately 9pt)

However, there is another issue: regular body text in small four-point size corresponds to 12pt,
so h4, h5, and h6 are disregarded as unnecessary. Three levels should be sufficient.

Next is how to distinguish between h3 and body text? What if someone just writes a bunch of text in four-point size? We rely on length, setting that only strings with a length less than or equal to 5 Chinese characters can be considered as h3, i.e., len(x) <= 15.

In Word, the sz value is twice the pt value, so it should be multiplied by 2, then check if it is greater than or equal to:

Thus, it is as follows:

- h1: 48<=x<inf
- h2: 36<=x<48
- h3: 28<=x<36 & len(x) <= 15

Additionally,
check if there is a numbered list at the beginning, such as 12345, or one two three, or "Chapter xxxx",
or if there is a w:numPr tag (Word's own numbering tag), then the limit can be relaxed to 15 Chinese characters = 15*3=45,

and allow the appearance of h4 level headers.


### Lastest Update

- [x] Support for recognizing titles set via style has been added. For titles set through Word styles, the system can now recognize these titles based on their style settings.