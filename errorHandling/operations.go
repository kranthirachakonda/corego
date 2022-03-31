package main

type CategoryError struct {
	reqCategory string
}

func (e *CategoryError) Error() string {
	return "Category" + e.reqCategory + "does not exist"
}

type ChanMessage struct {
	Category string
	Total    float64
	*CategoryError
}

func (slice ProductSlice) TotalPrice(category string) (total float64, err *CategoryError) {
	prodCount := 0
	for _, p := range slice {
		if p.Category == category {
			total += p.Price
			prodCount++
		}
	}
	if prodCount == 0 {
		err = &CategoryError{reqCategory: category}
	}
	return
}

func (slice ProductSlice) TotalPriceAsync(categories []string, priceChannel chan<- ChanMessage) {
	for _, c := range categories {
		total, err := slice.TotalPrice(c)
		priceChannel <- ChanMessage{
			Category:      c,
			Total:         total,
			CategoryError: err,
		}
	}
	close(priceChannel)
}
