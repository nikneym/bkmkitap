# bkmkitap
Unofficial Bkmkitap API

Usage Example
===========
```go
package main

import (
	"fmt"
	"log"

	"github.com/nikneym/bkmkitap"
)

func main() {
	user, err := bkmkitap.Login("john@doe.com", "johndoe123")
	if err != nil {
		log.Fatal(err)
	}

	basket, err := user.GetBasket()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(basket.Products)

	favorites, err := user.GetFavorites()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(favorites)
}

```
