package database  

type Product struct {
	ID          int     `json: "id"`
	Title       string  `json: "title"`
	Description string  `json: "description"`
	Price       float64 `json: "price"`
	ImageURL    string  `json: "imageurl"`
}

var productList []Product 

func Store (p Product) Product {
	p.ID = len(productList) + 1
	productList = append(productList, p)
	return p 
}

func List () []Product {
	return productList
}

func Get (productID int) *Product {
	for _, product := range productList{
		if product.ID == productID{
			return &product
		}
	}
	return nil 
}

func Update (product Product) {
	for idx, p := range productList {
		if p.ID == product.ID {
			productList[idx] = product
		}
	}
}

func Delete (productID int) {
	var tempList []Product

	for _, p := range productList {
		if p.ID != productID {
			tempList = append(tempList, p)
		}
	}

	productList = tempList 

}



func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Wall Hanging Vase",
		Description: "beautiful & elegant vases for your room ",
		Price:       999.99,
		ImageURL:    "https://imgs.search.brave.com/zZGPPoZoLxs617ldN_-Q3FANERcvSJRwznEPtN0K4W0/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly9pNS53/YWxtYXJ0aW1hZ2Vz/LmNvbS9zZW8vdHVt/Z2F0dGUtQ3JlYXRp/dmUtV2FsbC1IYW5n/aW5nLVZhc2VzLUJl/YXV0aWZ1bC1IaWdo/LWVuZC1XYWxsLUhh/bmdpbmctUmVzdGF1/cmFudC1XYWxsLURl/Y29yYXRpb24tV2Fs/bC1BY2Nlc3Nvcmll/cy1Ib21lLUxpdmlu/Zy1Sb29tLUJhY2tn/cm91bmQtV2FsbC1E/ZWNfYWVmNWU1ZDIt/ZjdlZC00NjE5LWJj/YjUtOTllMGI4NDRk/Njg3LmM5ZTU1YmM1/ZjA1ZmU4NDJmZjlj/ZjU0MDhlNjJjZjA4/LmpwZWc_b2RuSGVp/Z2h0PTU4MCZvZG5X/aWR0aD01ODAmb2Ru/Qmc9RkZGRkZG",
	}

	prd2 := Product{
		ID:          2,
		Title:       "Candle Holder",
		Description: "Aesthetic pieces of candle holders",
		Price:       799.99,
		ImageURL:    "https://imgs.search.brave.com/4zDbZV5F2Djlb0YJBplPLLAScD4DHMIGDex2M0UJlJ0/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly9jYi5z/Y2VuZTcuY29tL2lz/L2ltYWdlL0NyYXRl/L2NiX2RTQ18yMDI0/MTIyNl9EZWNvcl9D/YW5kbGVob2xkZXJz/P2JmYz1vbiZ3aWQ9/MTEzMCZxbHQ9ODA",
	}

	prd3 := Product{
		ID:          3,
		Title:       "Mirror",
		Description: "See the reflection of your eyes.",
		Price:       1299.50,
		ImageURL:    "https://imgs.search.brave.com/kIX67igGlil6t8DW93tQPdbU79LnoMABlGm4NondLKg/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly93d3cu/aWtlYS5jb20vZXh0/L2luZ2thZGFtL20v/NWYxN2M3YWNiMDIy/NjQwNy9vcmlnaW5h/bC9QRTg3MTU4Mi5q/cGc_Zj1z",
	}

	prd4 := Product{
		ID: 4,
		Title: "Leaf Lighting ",
		Description: "a touch of greenery ",
		Price: 499.99,
		ImageURL: "https://imgs.search.brave.com/Q6U42GmBPeK6NHF_wowu33GKfqcMvWP0kpoA8guLLxY/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly9tLm1l/ZGlhLWFtYXpvbi5j/b20vaW1hZ2VzL0kv/ODFhM3lqRlE2ekwu/anBn",
	}

	prd5 := Product{
		ID: 5,
		Title: "Galaxy tour",
		Description: "On & Go in the void of space",
		Price: 1100.99,
		ImageURL: "https://imgs.search.brave.com/ijMQNKQXVUe4lFcC1Y-m79ckNQWZ2DsxNFbIuT37BK0/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly93d3cu/YWVzdGhldGljcm9v/bWNvcmUuY29tL2Nk/bi9zaG9wL3Byb2R1/Y3RzLzVfOWNiYzZm/MjEtNWRjYS00NzJj/LThiNzYtNzA4MDc1/YTljZGM2XzQwMHgu/anBnP3Y9MTYyOTQ2/NjI3Mg",
	}

	prd6 := Product{
		ID: 6,
		Title: "Porcelain Jars",
		Description: "beautiful & elegant vases for your room ",
		Price: 899.99,
		ImageURL: "https://imgs.search.brave.com/7Gg4vbP10XA57BsXuMWtjHRYRjR9oY2UklMpt8zMxkc/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly9uZXN0/c2V0LmNvbS9jZG4v/c2hvcC9wcm9kdWN0/cy9zZXRzLXRhbmpp/ZXItNS1wYy1zZXQt/YnItcGV0aXRlLXBv/cmNlbGFpbi1qYXJz/LWRlY29yYXRpdmUt/Ym9va3MtMS5qcGc_/Y3JvcD1jZW50ZXIm/aGVpZ2h0PTE0MDAm/dj0xNTI1MzYzODAy/JndpZHRoPTE0MDA",
	}

	productList = append(productList, prd1)
	productList = append(productList, prd2)
	productList = append(productList, prd3)
	productList = append(productList, prd4)
	productList = append(productList, prd5)
	productList = append(productList, prd6)

}

