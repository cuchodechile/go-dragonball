// handler para las entradas tipo REST htpp (get/post)) seria transporte
package ginhandler

import (
	"net/http"

	"cuchodechile.cl/reto-amaris/internal/character"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *character.CharacterService
}

func New(svc *character.CharacterService) *Handler {
	return &Handler{svc: svc}
}

// registrar las rutas en el server 
func (h *Handler) Register(r *gin.Engine) {
	grp := r.Group("/characters")
	grp.GET("/:name", h.getByName)
	grp.POST("/", h.postByName)
}


type createReq struct {
	Name string `json:"name" binding:"required"`
}

func (h *Handler) postByName(c *gin.Context) {
	var req createReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name es requerido"})
		return
	}

	ch, err := h.svc.FindOrCreate(c, req.Name)
	switch err {
	case nil:
		c.JSON(http.StatusOK, ch)
	case character.ErrNotFound:
		c.JSON(http.StatusNotFound, gin.H{"error": "character not found"})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

// GET /characters/:name
func (h *Handler) getByName(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}

	ch, err := h.svc.FindOrCreate(c.Request.Context(), name)
	switch err {
	case nil:
		c.JSON(http.StatusOK, ch)
	case character.ErrNotFound:
		c.JSON(http.StatusNotFound, gin.H{"error": "character not found"})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
