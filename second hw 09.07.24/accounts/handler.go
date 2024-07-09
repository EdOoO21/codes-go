package accounts

import (
	"awesomeProject/accounts/dto"
	"awesomeProject/accounts/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
)

func New() *Handler {
	return &Handler{
		accounts: make(map[string]*models.Account),
		guard:    &sync.RWMutex{},
	}
}

type Handler struct {
	accounts map[string]*models.Account
	guard    *sync.RWMutex
}

func (h *Handler) CreateAccount(c echo.Context) error {
	var request dto.CreateAccountRequest // {"name": "alice", "amount": 50}
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)

		return c.String(http.StatusBadRequest, "invalid request")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "empty name")
	}

	h.guard.Lock()

	if _, ok := h.accounts[request.Name]; ok {
		h.guard.Unlock()

		return c.String(http.StatusForbidden, "account already exists")
	}

	h.accounts[request.Name] = &models.Account{
		Name:   request.Name,
		Amount: request.Amount,
	}

	h.guard.Unlock()

	return c.NoContent(http.StatusCreated)
}

func (h *Handler) GetAccount(c echo.Context) error {
	name := c.QueryParams().Get("name")

	h.guard.RLock()

	account, ok := h.accounts[name]

	h.guard.RUnlock()

	if !ok {
		return c.String(http.StatusNotFound, "account not found")
	}

	response := dto.GetAccountResponse{
		Name:   account.Name,
		Amount: account.Amount,
	}

	return c.JSON(http.StatusOK, response)
}

// Удаляет аккаунт
func (h *Handler) DeleteAccount(c echo.Context) error {
	name := c.QueryParams().Get("name")
	if len(name) == 0 {
		return c.String(http.StatusBadRequest, "empty name")
	}
	h.guard.Lock()
	if _, ok := h.accounts[name]; !ok {
		h.guard.Unlock()
		return c.String(http.StatusForbidden, "account does not exists")
	}
	delete(h.accounts, name)
	h.guard.Unlock()
	return c.NoContent(http.StatusCreated)
}

// Меняет баланс
func (h *Handler) PatchAccount(c echo.Context) error {
	var request dto.CreateAccountRequest // {"name": "alice", "amount": 50}
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)

		return c.String(http.StatusBadRequest, "invalid request")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "empty name")
	}

	h.guard.Lock()

	if _, ok := h.accounts[request.Name]; !ok {
		h.guard.Unlock()

		return c.String(http.StatusForbidden, "account doesn't exists")
	}

	h.accounts[request.Name].Amount = request.Amount

	h.guard.Unlock()

	return c.NoContent(http.StatusCreated)
}

// Меняет имя
func (h *Handler) ChangeAccount(c echo.Context) error {
	var request dto.ChangeAccountRequest
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "invalid request")
	}

	if 0 == len(request.New) || 0 == len(request.Name) {
		return c.String(http.StatusBadRequest, "empty name")
	}
	h.guard.Lock()

	if _, ok := h.accounts[request.Name]; !ok {
		h.guard.Unlock()
		return c.String(http.StatusConflict, "account not exists")
	}
	if _, ok := h.accounts[request.Name]; ok {
		h.guard.Unlock()
		return c.String(http.StatusConflict, "account already exists")
	}

	h.accounts[request.Name].Name = request.Name
	delete(h.accounts, request.Name)
	h.accounts[request.Name] = &models.Account{
		Name:   request.Name,
		Amount: h.accounts[request.Name].Amount,
	}
	h.guard.Unlock()
	return c.NoContent(http.StatusCreated)
}

// Написать клиент консольный, который делает запросы
