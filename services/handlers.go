package services

import (
	"monkCommerce/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (repo *Repository) CreateProduct(c *gin.Context) {
	var product storage.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := repo.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failure", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": product})
}

func (repo *Repository) GetAllProducts(c *gin.Context) {
	var products []storage.Product
	if err := repo.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": products})
}

func (repo *Repository) CreateCoupon(c *gin.Context) {
	var coupon storage.Coupon
	if err := c.ShouldBindJSON(&coupon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := repo.DB.Create(&coupon).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": coupon})
}

func (repo *Repository) GetAllCoupons(c *gin.Context) {
	var coupons []storage.Coupon
	if err := repo.DB.Find(&coupons).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": coupons})
}

func (repo *Repository) UpdateCoupon(c *gin.Context) {
	id := c.Param("id")
	var coupon storage.Coupon
	if err := repo.DB.First(&coupon, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Coupon not found"})
		return
	}

	if err := c.ShouldBindJSON(&coupon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	repo.DB.Save(&coupon)
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": coupon})
}

func (repo *Repository) GetCoupon(c *gin.Context) {
	id := c.Param("id")
	var coupon storage.Coupon
	if err := repo.DB.First(&coupon, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Coupon not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": coupon})
}

func (repo *Repository) GetApplicableCoupons(c *gin.Context) {
	// loggy := config.GetLogger()
	// jsonData, err := io.ReadAll(c.Request.Body)
	// if err != nil {
	// 	loggy.Error("error in getting coupon applicable body")
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "error in getting coupon applicable body"})
	// 	return
	// }

}

func (repo *Repository) ApplyCoupon(c *gin.Context) {
	return
}
