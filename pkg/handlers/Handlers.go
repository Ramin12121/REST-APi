package handlers

import (
	"Subscription/pkg/Server"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service Server.ServiceInterface
}

func NewHandler(s Server.ServiceInterface) *Handler {
	return &Handler{service: s}
}

// @Summary Get all subscriptions
// @Description Retrieve a list of all subscriptions
// @Tags subscriptions
// @ID get-all-subscriptions
// @Produce json
// @Success 200 {array} Server.Subscription
// @Router /subscriptions [get]
func (h *Handler) GetAll(c echo.Context) error {
	subscriptions, err := h.service.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get subscriptions"})
	}
	return c.JSON(http.StatusOK, subscriptions)

}

// @Summary Create a new subscription
// @Description Add a new subscription to the database
// @Tags subscriptions
// @ID create-subscription
// @Accept json
// @Produce json
// @Param subscription body Server.SubscriptionRequest true "Subscription data"
// @Success 201 {object} Server.Subscription
// @Failure 400 {object} map[string]string
// @Router /subscriptions [post]
func (h *Handler) Post(c echo.Context) error {
	var req Server.SubscriptionRequest
	if err := c.Bind(&req); err != nil {
		return newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	subsc, err := h.service.Create(&req)
	if err != nil {
		return newErrorResponse(c, http.StatusBadRequest, "Could not create subscription")
	}

	return c.JSON(http.StatusCreated, subsc)
}

// SubscriptionRequest структура для примера
// @Description Subscription data
// @Param subscription body Server.SubscriptionRequest true "Subscription data"
// @Schema
// type SubscriptionRequest struct {
//     Price       int    `json:"price"`
//     ServiceName string `json:"service_name"`
//     StartDate   string `json:"start_date"` // Убедитесь, что формат соответствует ISO 8601
//     UserID      string `json:"user_id"`
// }

// @Summary Update an existing subscription
// @Description Update subscription details by ID
// @Tags subscriptions
// @ID update-subscription
// @Accept json
// @Produce json
// @Param id path int true "Subscription ID"
// @Param subscription body Server.SubscriptionRequest true "Updated subscription data"
// @Success 200 {object} Server.Subscription
// @Failure 400 {object} map[string]string
// @Router /subscriptions/{id} [patch]
func (h *Handler) Patch(c echo.Context) error {
	id := c.Param("id")
	var req Server.SubscriptionRequest
	if err := c.Bind(&req); err != nil {
		return newErrorResponse(c, http.StatusBadRequest, "Invalid request")
	}
	updatedSubsc, err := h.service.Update(id, &req)
	if err != nil {
		return newErrorResponse(c, http.StatusBadRequest, "Could not update calculation")
	}
	return c.JSON(http.StatusOK, updatedSubsc)
}

// @Summary Delete a subscription
// @Description Remove a subscription by ID
// @Tags subscriptions
// @ID delete-subscription
// @Accept json
// @Produce json
// @Param id path int true "Subscription ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Router /subscriptions/{id} [delete]
func (h *Handler) Delete(c echo.Context) error {
	id := c.Param("id")
	err := h.service.Delete(id)
	if err != nil {
		return newErrorResponse(c, http.StatusBadRequest, "Could not delete calculation")
	}
	return c.NoContent(http.StatusNoContent)
}

// @Summary Get subscriptions by filter
// @Description Retrieve subscriptions based on certain filters
// @Tags subscriptions
// @ID get-subscriptions-by-filter
// @Accept json
// @Produce json
// @Param filter body Server.ToFilter true "Filter criteria"
// @Success 200 {array} Server.Subscription
// @Failure 400 {object} map[string]string
// @Router /subscriptionsByFilter [get]
func (h *Handler) GetByFilter(c echo.Context) error {
	var req Server.ToFilter
	if err := c.Bind(&req); err != nil {
		return newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	result, err := h.service.GetByFilter(&req)
	if err != nil {
		return newErrorResponse(c, http.StatusBadRequest, "Could not get by filter")
	}
	return c.JSON(http.StatusCreated, result)
}
