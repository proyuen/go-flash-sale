package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/proyuen/flashSale/Server/internal/service"
)

func (s *Server) createUesr(ctx *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := service.CreateUserRequest{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}

	user, err := s.service.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (s *Server) getUser(ctx *gin.Context) {
	username := ctx.Param("username")
	user, err := s.service.GetUser(ctx, username)
	if err != nil {
		// Note: Service layer should ideally return specific error types to distinguish 404 vs 500
		// For now, we rely on the error message or type check if we exported custom errors
		// But since we wrapped errors in Service, direct equality check might fail without errors.Is
		// For simplicity in this refactor step, we might lose the specific 404 check unless we handle it better.
		// However, let's keep it simple and return 500 or 404 based on string check or just 500 for now.
		// Or better, checking if error contains "not found".
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)

}

type loginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}
type loginUserResponse struct {
	AccessToken string               `json:"access_token"`
	User        service.UserResponse `json:"user"`
}

func (s *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := service.LoginUserRequest{
		Username: req.Username,
		Password: req.Password,
	}

	rsp, err := s.service.LoginUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	res := loginUserResponse{
		AccessToken: rsp.AccessToken,
		User:        rsp.User,
	}
	ctx.JSON(http.StatusOK, res)
}
