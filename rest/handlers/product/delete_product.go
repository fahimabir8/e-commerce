package product

import (
	"ecommerce/util"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pId, err := strconv.Atoi(productID)
	if err != nil {
		fmt.Println(err)
		util.SendError(w, http.StatusBadRequest, "Enter a valid Id")
		return
	}
	product, _ := h.productRepo.Get(pId)
	if product == nil {
		util.SendError(w, 404, "Product Not Found")
		return
	}
	err = h.productRepo.Delete(pId)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, http.StatusCreated, "Successfully deleted product")

}
