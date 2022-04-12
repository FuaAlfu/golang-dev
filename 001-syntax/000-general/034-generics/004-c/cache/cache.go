package cache

type cacheable interface {
    blog.Category | blog.Post
}


type cache[T cacheable] struct {
    data[string]T
}

func (c *cache[T]) Set(key string, value T) {
    c.data[key] = value
}

func (c *cache[T]) Get(key string) (v T) {
    if v, ok := c.data[key]; ok {
        return v
    }

    return
}

func New[T cacheable]() {
    c := cache[T]{}
    c.data = make(map[string]T)

    return &c
}