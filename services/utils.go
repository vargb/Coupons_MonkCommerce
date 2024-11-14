package services

import "monkCommerce/storage"

func maxof3(a, b, c int) int {
	if a >= b && a >= c {
		return a
	}
	if b >= c && b >= a {
		return b
	}
	return c
}

func calculateDiscount(coupon storage.Coupon, cart storage.Cart) float64 {
	discount := 0.0

	switch coupon.Type {
	case "cart-wise":
		if cart.Total > coupon.CartWiseRule.MinimumAmount {
			discount = cart.Total * (coupon.CartWiseRule.Discount / 100)
		}

	case "product-wise":
		for _, item := range cart.Items {
			if containsProducts(coupon.ProductWiseRule.Products, item.ProductID) {
				discount += item.Price * float64(item.Quantity) * (coupon.ProductWiseRule.Discount / 100)
			}
		}

	case "bxgy":
		buyCount := 0
		getEligibleItems := []storage.CartItem{}
		for _, item := range cart.Items {
			if containsBuyGetProducts(coupon.BxGyRule.BuyProducts, item.ProductID) {
				buyCount += item.Quantity
			} else if containsBuyGetProducts(coupon.BxGyRule.GetProducts, item.ProductID) {
				getEligibleItems = append(getEligibleItems, item)
			}
		}
		applicableTimes := min(buyCount/len(coupon.BxGyRule.BuyProducts), coupon.BxGyRule.RepetitionLimit)
		freeItemsCount := applicableTimes * len(coupon.BxGyRule.GetProducts)
		for _, item := range getEligibleItems {
			freeItemQty := min(item.Quantity, freeItemsCount)
			discount += item.Price * float64(freeItemQty)
			freeItemsCount -= freeItemQty
		}
	}

	return discount
}

func contains(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

func containsBuyGetProducts(arr []storage.BxGyProduct, str string) bool {
	for _, product := range arr {
		if containsProducts(product.Products, str) {
			return true
		}
	}
	return false
}

func containsProducts(arr []storage.Product, str string) bool {
	for _, product := range arr {
		if product.Code == str {
			return true
		}
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
