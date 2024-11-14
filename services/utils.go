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
		if cart.Total > coupon.ThresholdAmount {
			discount = cart.Total * (coupon.DiscountPercent / 100)
		}

	case "product-wise":
		for _, item := range cart.Items {
			if contains(coupon.BuyArray, item.ProductID) {
				discount += item.UnitPrice * float64(item.Quantity) * (coupon.DiscountPercent / 100)
			}
		}

	case "bxgy":
		buyCount := 0
		getEligibleItems := []CartItem{}
		for _, item := range cart.Items {
			if contains(coupon.BuyArray, item.ProductID) {
				buyCount += item.Quantity
			} else if contains(coupon.GetArray, item.ProductID) {
				getEligibleItems = append(getEligibleItems, item)
			}
		}
		applicableTimes := min(buyCount/coupon.X, coupon.RepetitionLimit)
		freeItemsCount := applicableTimes * coupon.Y
		for _, item := range getEligibleItems {
			freeItemQty := min(item.Quantity, freeItemsCount)
			discount += item.UnitPrice * float64(freeItemQty)
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
