package handlers

import (
	"webstore/models"

	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type ItemStore interface {
	GetItem(ctx context.Context, itemID int) (*models.Item, error)
	GetItems(ctx context.Context) ([]*models.Item, error)
}

type Handler struct {
	Store ItemStore
}

func NewHandler(store ItemStore) *Handler {
	return &Handler{
		Store: store,
	}
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func (h *Handler) GetItems(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	items, err := h.Store.GetItems(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, items)
}

func (h *Handler) GetItemByID(w http.ResponseWriter, r *http.Request) {
	// 1: Extract the item ID from the URL path
	itemIDStr := r.PathValue("item_id")
	if itemIDStr == "" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	// 2: Convert the item ID to an integer
	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, ErrorMessageResponse{Message: "item_id must be a valid integer"})
		return
	}

	// 3: Fetch the item from the store
	ctx := r.Context()
	item, err := h.Store.GetItem(ctx, itemID)
	if err != nil {
		fmt.Printf("error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if item == nil {
		writeJSON(w, http.StatusNotFound, nil)
		return
	}
	writeJSON(w, http.StatusOK, item)
}
