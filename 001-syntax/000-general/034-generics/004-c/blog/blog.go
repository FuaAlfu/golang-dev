package blog

type Category struct {
    ID int32
    Name string
    Slug string
}

type Post struct {
    ID int32
    Categories []Category
    Title string
    Description
    Slug string
}