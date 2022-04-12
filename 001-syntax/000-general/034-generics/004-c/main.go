package main

import(
	"fmt"

	"github.com/FuaAlfu/golang-dev/tree/master/001-syntax/000-general/034-generics/004-c/blog"
	"github.com/FuaAlfu/golang-dev/tree/master/001-syntax/000-general/034-generics/004-c/cache"
)

func main() {
	fmt.Println("Starting..")
	
    // create a new category
    category := blog.Category{
        ID: 1,
        Name: "Go Generics",
        Slug: "go-generics",
    }
    // create cache for blog.Category struct
    cc := cache.New[blog.Category]()
    // add category to cache
    cc.Set(category.Slug, category)

    // create a new post
    post := blog.Post{
        ID: 1,
        Categories: []blog.Category{
            {ID: 1, Name: "Go Generics", Slug: "go-generics"},
        },
        Title: "Generics in Golang structs",
        Text: "Here go's the text",
        Slug: "generics-in-golang-structs",
    }
    // create cache for blog.Post struct
    cp := cache.New[blog.Post]()
    // add post to cache
    cp.Set(post.Slug, post)
}