package routes



import (
	"expendit-server/controller"
	"github.com/gin-gonic/gin"
)

func ExpenseRoutes(incomingRoutes *gin.Engine){
	
	incomingRoutes.GET("/expense", controller.GetExpenses())
	incomingRoutes.GET("/expense/:id", controller.GetExpenseById())
	incomingRoutes.POST("/expense/create",controller.CreateExpense())
	incomingRoutes.PUT("/expense/:id",controller.UpdateExpense())
	incomingRoutes.DELETE("/expense/:id",controller.DeleteExpense())
	incomingRoutes.GET("/expense/search", controller.SearchExpense())
}