package rest

import "github.com/labstack/echo/v4"

func LoadRoutes(e *echo.Echo, handler *handler) {
	authMiddleware := GetAuthMiddleware(handler.restoUsecase)

	menuGroup := e.Group("/menu")
	menuGroup.GET("", handler.GetMenuList) // http://localhost:14045/menu?menu_type=drink atau http://localhost:14045/menu

	orderGroup := e.Group("/order")
	orderGroup.POST("", handler.Order,
		authMiddleware.checkAuth,
	)
	orderGroup.GET("/:orderID", handler.GetOrderInfo,
		authMiddleware.checkAuth,
	)

	userGroup := e.Group("/user")
	userGroup.POST("/register", handler.RegisterUser)
	userGroup.POST("/login", handler.Login)
}