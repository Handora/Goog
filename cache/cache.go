package cache

type TagCache struct {
	Id int
	Name string
}

type ArticleCache struct {
	Id 		int
	Title   string
	Content string
	Tag     []string
}

