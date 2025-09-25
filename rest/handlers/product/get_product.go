package product

import (
	"ecommerce/util"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request){
	productID := r.PathValue("id")

	pId , err := strconv.Atoi(productID)
	if err != nil {
		fmt.Println(err)
		util.SendError(w,http.StatusBadRequest,"Invalid Id")
		return 
	}

	product,err :=  h.productRepo.Get(pId)// database.Get(pId) 
	if err != nil {
		fmt.Println(err)
		util.SendError(w,http.StatusInternalServerError,"Internal Server Error")
		return
	}
	if product == nil {
		util.SendError(w,404,"Product Not Found")
		return 
	}

	util.SendData(w, http.StatusOK, product)	
}