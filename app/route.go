package app

import (
	"Test-Rizky/config"
	"Test-Rizky/controller"
	"Test-Rizky/repository"
	"Test-Rizky/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.Connection()
	// repository
	orderRepository    repository.OrderRepository    = repository.NewOrderRepository(db)
	customerRepository repository.CustomerRepository = repository.NewCustomerRepository(db)

	// service
	orderService    service.OrderService    = service.NewOrderService(orderRepository)
	customerService service.CustomerService = service.NewCustomerService(customerRepository)
	authService     service.AuthService     = service.NewAuthService(customerRepository)

	// controller
	orderController    controller.OrderController    = controller.NewOrderController(orderService, authService)
	customerController controller.CustomerController = controller.NewCustomerController(customerService)
	authController     controller.AuthController     = controller.NewAuthController(authService)
)

func Route() {

	r := gin.Default()
	order := r.Group("order")
	{
		order.POST("/add", orderController.Add)
		order.PUT("/update", orderController.Update)
		order.GET("detail/:id", orderController.GetDetail)
		order.POST("/list", orderController.GetAllData)
		order.DELETE("/delete/:id", orderController.Delete)
		order.GET("/tes")
	}

	customer := r.Group("customer")
	{
		customer.POST("/add", customerController.Add)
		customer.PUT("/update", customerController.Update)
		customer.GET("detail/:id", customerController.GetDetail)
		customer.POST("/list", customerController.GetAllData)
		customer.DELETE("/delete/:id", customerController.Delete)
	}

	login := r.Group("auth")
	{
		login.POST("/login", authController.GenerateToken)
	}
	err := r.Run()
	if err != nil {
		return
	}
}
