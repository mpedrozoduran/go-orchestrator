package runner

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mpedrozoduran/go-orchestrator/internal/api"
	"github.com/mpedrozoduran/go-orchestrator/internal/config"
	"github.com/mpedrozoduran/go-orchestrator/internal/controller"
	"github.com/mpedrozoduran/go-orchestrator/internal/middleware"
	"github.com/mpedrozoduran/go-orchestrator/internal/repository/integrations"
	"github.com/mpedrozoduran/go-orchestrator/internal/repository/persistence"
	"github.com/mpedrozoduran/go-orchestrator/internal/repository/persistence/entities"
	"github.com/mpedrozoduran/go-orchestrator/internal/util"
	"log"
)

type ServiceRestAPI struct {
	config.Config
	persistence.DbClient
	api.PaymentsAPI
	api.RefundsAPI
	api.AuditTrailAPI
	router *gin.Engine
}

func NewServiceRestAPI() ServiceRestAPI {
	// initialize main app config
	appConfig := config.LoadConfig()

	// initialize db client
	dbClient := persistence.NewDbClient(appConfig)

	// initialize repositories
	paymentsRepository := persistence.NewPaymentRepository[entities.Payment](dbClient, util.PaymentsTable)
	refundsRepository := persistence.NewPaymentRepository[entities.Refund](dbClient, util.RefundsTable)
	auditTrailsRepository := persistence.NewPaymentRepository[entities.AuditTrail](dbClient, util.AuditTrailTable)

	coreBankClient := integrations.NewCoreBankClient(appConfig)

	// initialize controllers: business logic
	paymentsController := controller.NewPaymentController(paymentsRepository, auditTrailsRepository, coreBankClient)
	refundsController := controller.NewRefundController(refundsRepository, auditTrailsRepository, coreBankClient)
	auditTrailsController := controller.NewAuditTrailController(auditTrailsRepository)

	// initialize api
	authAPI := api.NewAuthAPI(appConfig)
	paymentsAPI := api.NewPaymentsAPI(paymentsController)
	refundsAPI := api.NewRefundsAPI(refundsController)
	auditTrailsAPI := api.NewAuditTrailAPI(auditTrailsController)

	// initialize middlewares
	authMiddleware := middleware.NewMiddleware(appConfig)

	// Initialize Gin router
	router := gin.Default()

	// setup routes
	router.POST("/v1/auth/login", authAPI.Login)
	v1 := router.Group("/v1/payments")
	{
		v1.POST("", authMiddleware.AuthMiddleware(), paymentsAPI.ProcessPayment)
		v1.GET("/:id", authMiddleware.AuthMiddleware(), paymentsAPI.GetPayment)
		v1.GET("/history", authMiddleware.AuthMiddleware(), auditTrailsAPI.GetAuditTrails)
		v1.POST("/refund", authMiddleware.AuthMiddleware(), refundsAPI.ProcessRefund)
	}

	return ServiceRestAPI{
		Config:        appConfig,
		DbClient:      dbClient,
		PaymentsAPI:   paymentsAPI,
		RefundsAPI:    refundsAPI,
		AuditTrailAPI: auditTrailsAPI,
		router:        router,
	}
}

func (s ServiceRestAPI) Run() {
	// run server
	err := s.router.Run(fmt.Sprintf(":%v", s.Config.Server.Port))
	if err != nil {
		log.Fatal(err)
	}
}
