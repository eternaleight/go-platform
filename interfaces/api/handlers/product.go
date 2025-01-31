package handlers

import (
	"net/http"
	"strconv"

	"github.com/eternaleight/go-backend/domain/models"
	"github.com/gin-gonic/gin"
)

type ProductUsecasesInterface interface {
	CreateProduct(input models.ProductInput) (*models.Product, error)
	ListProducts() ([]models.Product, error)
	GetProductByID(id uint) (*models.Product, error)
	UpdateProduct(id uint, input models.ProductInput) (*models.Product, error)
	DeleteProduct(id uint) error
}

type ProductHandler struct {
	ProductUsecases ProductUsecasesInterface
}

// 新しいProductHandlerのインスタンスを初期化
func NewProductHandler(productUsecases ProductUsecasesInterface) *ProductHandler {
	return &ProductHandler{
		ProductUsecases: productUsecases,
	}
}

// 新しい商品を作成
func (ph *ProductHandler) CreateProduct(c *gin.Context) {
	var product models.ProductInput
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "商品のデータの形式が正しくありません。"})
		return
	}

	createdProduct, err := ph.ProductUsecases.CreateProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "データベースに商品を保存できませんでした。"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": createdProduct})
}

// 全商品をリストとして取得
func (ph *ProductHandler) ListProducts(c *gin.Context) {
	products, err := ph.ProductUsecases.ListProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "商品のリストの取得に失敗しました。"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": products})
}

// 指定されたIDの商品を取得
func (ph *ProductHandler) GetProductByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "商品IDが無効です。"})
		return
	}

	product, err := ph.ProductUsecases.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "商品の情報を取得できませんでした。"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": product})
}

// 指定されたIDの商品情報を更新
func (ph *ProductHandler) UpdateProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "商品IDが無効です。"})
		return
	}

	var product models.ProductInput
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "商品のデータの形式が正しくありません。"})
		return
	}

	updatedProduct, err := ph.ProductUsecases.UpdateProduct(uint(id), product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "商品の更新に失敗しました。"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": updatedProduct})
}

// 指定されたIDの商品を削除
func (ph *ProductHandler) DeleteProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "商品IDが無効です。"})
		return
	}

	err = ph.ProductUsecases.DeleteProduct(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "商品の削除に失敗しました。"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "商品は正常に削除されました。"})
}
