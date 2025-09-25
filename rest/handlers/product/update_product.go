package product

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type ReqUpdateProduct struct {
	ID          int     `json: "id"`
	Title       string  `json: "title"`
	Description string  `json: "description"`
	Price       float64 `json: "price"`
	ImageURL    string  `json: "imageurl"`
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pId, err := strconv.Atoi(productID)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	var req ReqUpdateProduct
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)
	if err != nil {
		fmt.Println(err)
		util.SendError(w, http.StatusBadRequest, "Invalid Request Body")
		return
	}

	_,err = h.productRepo.Update(repo.Product{
		ID: pId,
		Title: req.Title,
		Description: req.Description,
		Price: req.Price,
		ImageURL: req.ImageURL,
	})

	util.SendData(w, http.StatusOK, "Successfully updated product")

}
