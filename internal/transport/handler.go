package transport

import (
	"app/internal/services"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	service *services.CertificateTLSService
}

func NewHandler(s *services.CertificateTLSService) *Handler {
	return &Handler{service: s}
}

func (h *Handler) Register(router *gin.Engine) {
	router.POST("/certificates", h.CreateCertificateTLS)
	router.GET("/certificates", h.GetAllCertificatesTLS)
	router.GET("/certificates/:id/ttl", h.GetCertificateTLSByID)
	router.GET("/certificates/ttl", h.CheckCertificateTLS)
}

// CreateCertificateTLS godoc
// @Summary Добавить URL
// @Description Добавляет URL для мониторинга
// @Tags urls
// @Accept json
// @Produce json
// @Param request body AddURLRequest true "URL"
// @Success 201
// @Failure 400 {object} ErrorResponse
// @Router /certificates [post]
func (h *Handler) CreateCertificateTLS(c *gin.Context) {
	var req AddURLRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.AddURL(c.Request.Context(), req.URL); err != nil {
		mapServiceError(c, err)
		return
	}

	c.Status(http.StatusCreated)
}

// GetAllCertificatesTLS godoc
// @Summary Получить все сертификаты
// @Tags certificates
// @Produce json
// @Param page query int false "Номер страницы" minimum(1)
// @Param limit query int false "Размер страницы" minimum(1) maximum(100)
// @Success 200 {object} CertificatesListResponse
// @Failure 400 {object} ErrorResponse
// @Router /certificates [get]
func (h *Handler) GetAllCertificatesTLS(c *gin.Context) {
	var req PaginationQuery
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Page == 0 {
		req.Page = 1
	}
	if req.Limit == 0 {
		req.Limit = 20
	}

	certs, err := h.service.GetAll(c.Request.Context(), req.Page, req.Limit)
	if err != nil {
		mapServiceError(c, err)
		return
	}

	responses := make([]CertificateResponse, 0, len(certs))
	for _, cert := range certs {
		responses = append(responses, toResponse(cert))
	}

	c.JSON(http.StatusOK, gin.H{
		"data": responses,
		"meta": gin.H{
			"page":  req.Page,
			"limit": req.Limit,
		},
	})
}

// GetCertificateTLSByID godoc
// @Summary Получить сертификат по ID в БД
// @Tags certificate
// @Produce json
// @Param id path string true "ID сертификата"
// @Success 200 {object} CertificateResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /certificates/{id}/ttl [get]
func (h *Handler) GetCertificateTLSByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	cert, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		mapServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, toResponse(cert))
}

// CheckCertificateTLS godoc
// @Summary Получить сертификат по ссылке без записи в БД
// @Tags certificate
// @Produce json
// @Param url query string true "Цель URL"
// @Success 200 {object} CertificateResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /certificates/ttl [get]
func (h *Handler) CheckCertificateTLS(c *gin.Context) {
	url := c.Query("url")
	if url == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "url is required"})
		return
	}

	cert, err := h.service.GetByURL(url)
	if err != nil {
		mapServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, toResponse(cert))
}

func mapServiceError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, services.ErrInvalidURL):
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	case errors.Is(err, services.ErrInvalidArgs):
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	case errors.Is(err, services.ErrTLS):
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	case errors.Is(err, services.ErrAlreadyExists):
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	case errors.Is(err, services.ErrNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
	}
}
