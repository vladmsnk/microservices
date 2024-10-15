package user

import "github.com/gin-gonic/gin"

type Handler struct {
}

func NewHandler(r *gin.RouterGroup) *Handler {
	h := &Handler{}

	r.POST("/register", h.Register)
	r.POST("/login", h.Login)
	return h
}

// Register godoc
// @Summary Register endpoint
// @Schemes
// @Description Register user
// @Tags example
// @Accept json
// @Produce json
// @Param requestBody body RegisterRequest true "Register user"
// @Success 200 {string} ok
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /user/register [post]
func (h *Handler) Register(ctx *gin.Context) {

}

// Login godoc
// @Summary Login endpoint
// @Schemes
// @Description Login user
// @Tags example
// @Accept json
// @Produce json
// @Param requestBody body LoginRequest true "Login user"
// @Success 200 {string} token
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /user/login [post]
func (h *Handler) Login(ctx *gin.Context) {

}
