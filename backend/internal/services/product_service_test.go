package services

import (
	"testing"

	"ecommerce-api/internal/models"
)

func setupProducts() {
	products = []models.Product{
		{ID: "p1", Name: "Red Shirt", Price: 10.0, Stock: 5},
		{ID: "p2", Name: "Blue Shirt", Price: 15.0, Stock: 3},
		{ID: "p3", Name: "Green Hat", Price: 8.0, Stock: 10},
		{ID: "p4", Name: "Yellow Shirt", Price: 20.0, Stock: 2},
		{ID: "p5", Name: "Black Shoes", Price: 50.0, Stock: 1},
	}
}

func TestSearchProducts_ByName(t *testing.T) {
	setupProducts()
	items, total := SearchProducts("shirt", 1, 10, 0, 0)
	if total != 3 {
		t.Fatalf("expected total 3, got %d", total)
	}
	if len(items) != 3 {
		t.Fatalf("expected 3 items, got %d", len(items))
	}
}

func TestSearchProducts_Pagination(t *testing.T) {
	setupProducts()
	items, total := SearchProducts("", 2, 2, 0, 0)
	if total != 5 {
		t.Fatalf("expected total 5, got %d", total)
	}
	if len(items) != 2 {
		t.Fatalf("expected 2 items on page 2, got %d", len(items))
	}
	if items[0].ID != "p3" {
		t.Fatalf("expected p3 at start of page 2, got %s", items[0].ID)
	}
}

func TestSearchProducts_PriceFilter(t *testing.T) {
	setupProducts()
	items, total := SearchProducts("", 1, 10, 10.0, 20.0)
	if total != 3 {
		t.Fatalf("expected 3 items in price range, got %d", total)
	}
	if len(items) != 3 {
		t.Fatalf("expected 3 items returned, got %d", len(items))
	}
}
