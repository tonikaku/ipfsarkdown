package markdowner


import (
	"github.com/russross/blackfriday"
)

type MarkdownTransformer struct {
	transform func([]byte) []byte
}

var MdTransformer = GetNewMarkdownTransformer()

func GetNewMarkdownTransformer() *MarkdownTransformer {
	return &MarkdownTransformer{blackfriday.MarkdownCommon}
}

func (md *MarkdownTransformer) UseBasic() {
	md.transform = blackfriday.MarkdownBasic
}

func (md *MarkdownTransformer) Transform(origin []byte) []byte {
	return md.transform(origin)
}