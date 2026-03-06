package routes

import (
	"backend/controllers"
	"backend/repositories"
	"backend/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Setup(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api")

	userRepo := repositories.UserRepository{DB: db}
	userServ := services.UserService{Repo: userRepo}
	userCtrl := controllers.UserController{Service: userServ}
	UserRoutes(api, &userCtrl)

	catRepo := repositories.ProductCategoryRepository{DB: db}
	catServ := services.ProductCategoryService{Repo: catRepo}
	catCtrl := controllers.ProductCategoryController{Service: catServ}
	ProductCategoryRoutes(api, &catCtrl)

	socialRepo := repositories.SocialMediaRepository{DB: db}
	socialServ := services.SocialMediaService{Repo: socialRepo}
	socialCtrl := controllers.SocialMediaController{Service: socialServ}
	SocialMediaRoutes(api, &socialCtrl)

	productRepo := repositories.ProductRepository{DB: db}
	productServ := services.ProductService{Repo: productRepo}
	productCtrl := controllers.ProductController{Service: productServ}
	ProductRoutes(api, &productCtrl)

	storeRepo := repositories.StoreRepository{DB: db}
	storeServ := services.StoreService{Repo: storeRepo}
	storeCtrl := controllers.StoreController{Service: storeServ}
	StoreRoutes(api, &storeCtrl)

	storeCategoryRepo := repositories.StoreCategoryRepository{DB: db}
	storeCategoryServ := services.StoreCategoryService{Repo: storeCategoryRepo}
	storeCategoryCtrl := controllers.StoreCategoryController{Service: storeCategoryServ}
	StoreCategoryRoutes(api, &storeCategoryCtrl)

}
