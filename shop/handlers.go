package shop

import (
	"net/http"

	"github.com/bencbradshaw/framework"
)

// Product represents a shop product with all its details
type Product struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       string   `json:"price"`
	Features    []string `json:"features"`
}

// getProducts returns the list of available products
func getProducts() []Product {
	return []Product{
		{
			Name:        "Pro Starter Kit",
			Description: "Everything you need to get started with professional development",
			Price:       "$49.99",
			Features:    []string{"Premium templates", "24/7 support", "Advanced analytics"},
		},
		{
			Name:        "Enterprise Solution",
			Description: "Complete enterprise-grade solution for large teams",
			Price:       "$199.99",
			Features:    []string{"Unlimited projects", "Priority support", "Custom integrations"},
		},
		{
			Name:        "Developer Toolkit",
			Description: "Professional tools and resources for developers",
			Price:       "$99.99",
			Features:    []string{"Advanced debugging", "Performance monitoring", "Code optimization"},
		},
	}
}

// Handler returns the shop page handler
func Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		framework.RenderWithHtmlResponse(w, "shop.custom.html", map[string]any{
			"Title": "Shop",
			"Body":  "Welcome to our premium product collection!",
			"Items": getProducts(),
		})
	}
}
