package books

const CategoryNovel = 1
const CategoryShortStory = 2

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b Book) Category() int {
	if b.Pages > 300 {
		return CategoryNovel
	}

	return CategoryShortStory
}
