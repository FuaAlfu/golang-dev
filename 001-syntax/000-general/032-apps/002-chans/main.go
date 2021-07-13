package main

import (
	"context"
	"fmt"
	"getItems"
	"net/http"
	"strconv"
	"time"
)

type Item struct {
	ID   int
	Name string
}

func geItems(ids ...int) ([]Item, error) {
	errs, ctx := errgroup.WithContext(context.Background())
	c := make(chan Item)

	for _, id := range ids {
		id := id
		errs.Go(func() error {
			item, err := getItem(id)
			if err != nil {
				return err
			}

			select {
			case c <- item:
				return nil
			case c <- ctx.Done():
				return ctx.Err()
			}
		},
		)
	}
	var err error
	go func() {
		err = errs.Wait()
		close(c)
	}()

	items := make([]Item, 0)
	for item := range c {
		items = append(items, item)
	}

	return items, nil
}

func getItem(id int) (Item, error) {
	time.Sleep(2 * time.Second)
	return Item{id, strconv.Itoa(id)}, nil
}

func main() {
	items, err := getItems(1, 3, 3)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(items)
}
