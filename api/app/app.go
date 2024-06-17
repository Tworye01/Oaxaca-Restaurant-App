package app

import (
	"log"
	"os"
	"team-project/database"
	"team-project/handlers"
	"team-project/middleware"
	"team-project/models"
	"team-project/server"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

// Configure the App
type Config struct {
	Port string
	Log  bool
}

// Stores its own config, the database store and the fiber app
type App struct {
	Config Config
	Store  *database.Store
	Fiber  *fiber.App
}

// Creates a new App, optionally can give your own config otherwise it will use the default config
func New(args ...Config) App {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	c := Config{
		Port: ":8080",
		Log:  false,
	}
	s := database.New()
	f := fiber.New()

	if len(args) > 0 {
		c = args[0]
	}

	return App{
		Config: c,
		Store:  s,
		Fiber:  f,
	}
}

// app godoc
//
//	@title			Oaxaca
//	@version		0.1
//	@description	Oaxaca backend

//	@contact.name	Team 25

//	@license.name	MIT
//	@license.url	https://opensource.org/license/mit/

// @host		localhost:8080
// @BasePath	/
func (a *App) Run() {
	log.Println("app running")

	// Migrate tables
	a.Store.AutoMigrate(models.Menu{}, models.MenuItem{}, models.Drink{}, models.Dish{}, models.Credentials{}, models.Order{}, models.Bill{})

	// Enable logger
	a.Fiber.Use(logger.New())

	// Add doc endpoints
	a.Fiber.Use(swagger.New(swagger.Config{
		BasePath: "/api/v1/",
		FilePath: "./docs/swagger.json",
		Path:     "docs",
		Title:    "oaxaca docs",
	}))

	// Enable CORS when in dev mode
	a.Fiber.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {
			return os.Getenv("ENVIRONMENT") == "development"
		},
	}))

	var routes = []Route{
		// Health Handlers
		{
			Path:   "/api/v1/health",
			Method: Get,
			Middleware: []fiber.Handler{
				middleware.BasicMiddleware(),
			},
			Handler: handlers.GetHealth(),
		},
		// Menu Handlers
		{
			Path:   "/api/v1/menu",
			Method: Get,
			Middleware: []fiber.Handler{
				middleware.RequiresRole(models.Customer, a.Store),
			},
			Handler: handlers.GetMenu(a.Store),
		},
		// Dish Handlers

		{
			Path:   "/api/v1/menu/dishes",
			Method: Get,
			Middleware: []fiber.Handler{
				middleware.RequiresRole(models.Customer, a.Store),
			},
			Handler: handlers.GetFilter(a.Store),
		},
		{
			Path:   "/api/v1/menu/dishes/:id",
			Method: Get,
			Middleware: []fiber.Handler{
				middleware.RequiresRole(models.Customer, a.Store),
			},
			Handler: handlers.GetDish(a.Store),
		},
		{
			Path:   "/api/v1/menu/dishes",
			Method: Post,
			Middleware: []fiber.Handler{
				middleware.RequiresRole(models.Waiter, a.Store),
			},
			Handler: handlers.PostDish(a.Store),
		},
		{
			Path:   "/api/v1/menu/dishes",
			Method: Patch,
			Middleware: []fiber.Handler{
				middleware.RequiresRole(models.Waiter, a.Store),
			},
			Handler: handlers.PatchDish(a.Store),
		},
		{
			Path:   "/api/v1/menu/dishes/:id",
			Method: Delete,
			Middleware: []fiber.Handler{
				middleware.RequiresRole(models.Waiter, a.Store),
			},
			Handler: handlers.DeleteDish(a.Store),
		},
		// Drink Handlers
		{
			Path:   "/api/v1/menu/drinks",
			Method: Get,
			Middleware: []fiber.Handler{
				middleware.RequiresRole(models.Customer, a.Store),
			},
			Handler: handlers.ListDrinks(a.Store),
		},
		{
			Path:   "/api/v1/menu/drinks/:id",
			Method: Get,
			Middleware: []fiber.Handler{
				middleware.RequiresRole(models.Customer, a.Store),
			},
			Handler: handlers.GetDrink(a.Store),
		},
		{
			Path:   "/api/v1/menu/drinks",
			Method: Post,
			Middleware: []fiber.Handler{
				middleware.RequiresRole(models.Waiter, a.Store),
			},
			Handler: handlers.PostDrink(a.Store),
		},
		{
			Path:   "/api/v1/menu/drinks",
			Method: Patch,
			Middleware: []fiber.Handler{
				middleware.RequiresRole(models.Waiter, a.Store),
			},
			Handler: handlers.PatchDrink(a.Store),
		},
		{
			Path:   "/api/v1/menu/drinks/:id",
			Method: Delete,
			Middleware: []fiber.Handler{
				middleware.RequiresRole(models.Waiter, a.Store),
			},
			Handler: handlers.DeleteDrink(a.Store),
		},
		// Image Handlers
		{
			Path:    "/images/:dir/:file",
			Method:  Get,
			Handler: handlers.GetImage(),
		},
		{
			Path:    "/images/:dir/:file",
			Method:  Post,
			Handler: handlers.PostImage(),
		},
		// User Handlers
		{
			Path:   "/api/v1/login",
			Method: Post,
			Middleware: []fiber.Handler{
				middleware.RequiresRole(models.Customer, a.Store),
			},
			Handler: handlers.LoginHandler(a.Store),
		},
		{
			Path:   "/api/v1/login/credentials",
			Method: Get,
			Middleware: []fiber.Handler{
				middleware.RequiresRole(models.Customer, a.Store),
			},
			Handler: handlers.LoginHandler(a.Store),
		},
		{
			Path:   "/api/v1/login/credentials/:username",
			Method: Get,
			Middleware: []fiber.Handler{
				middleware.RequiresRole(models.Admin, a.Store),
			},
			Handler: handlers.GetUser(a.Store),
		},
		{
			Path:   "/api/v1/login/credentials",
			Method: Patch,
			Middleware: []fiber.Handler{
				middleware.RequiresRole(models.Admin, a.Store),
			},
			Handler: handlers.UpdateUser(a.Store),
		},
		{
			Path:   "/api/v1/login/credentials/:user",
			Method: Delete,
			Middleware: []fiber.Handler{
				middleware.RequiresRole(models.Admin, a.Store),
			},
			Handler: handlers.DeleteUser(a.Store),
		},
		{
			Path:   "/api/v1/login/credentials",
			Method: Post,
			Middleware: []fiber.Handler{
				middleware.RequiresRole(models.Customer, a.Store),
			},
			Handler: handlers.PostUser(a.Store),
		},
		// Order Handlers
		{
			Path:   "/api/v1/orders",
			Method: Get,
			Middleware: []fiber.Handler{
				middleware.RequiresRole(models.Customer, a.Store),
			},
			Handler: handlers.ListOrder(a.Store),
		},
		{
			Path:   "/api/v1/orders/:id",
			Method: Get,
			Middleware: []fiber.Handler{
				middleware.RequiresRole(models.Customer, a.Store),
			},
			Handler: handlers.GetOrder(a.Store),
		},
		{
			Path:   "/api/v1/orders/:id/status",
			Method: Get,
			Middleware: []fiber.Handler{
				middleware.RequiresRole(models.Customer, a.Store),
			},
			Handler: handlers.GetOrderStatus(a.Store),
		},
		{
			Path:   "/api/v1/orders",
			Method: Post,
			Middleware: []fiber.Handler{
				middleware.RequiresRole(models.Customer, a.Store),
			},
			Handler: handlers.PostOrder(a.Store),
		},
		{
			Path:   "/api/v1/orders",
			Method: Patch,
			Middleware: []fiber.Handler{
				middleware.RequiresRole(models.Waiter, a.Store),
			},
			Handler: handlers.PatchOrder(a.Store),
		},
		{
			Path:   "/api/v1/orders/:id/:status",
			Method: Patch,
			Middleware: []fiber.Handler{
				middleware.RequiresRole(models.Customer, a.Store),
			},
			Handler: handlers.PatchOrderStatus(a.Store),
		},
		{
			Path:   "/api/v1/orders/:id",
			Method: Delete,
			Middleware: []fiber.Handler{
				middleware.RequiresRole(models.Waiter, a.Store),
			},
			Handler: handlers.DeleteOrder(a.Store),
		},
		{
			Path:   "/api/v1/table/orders/:table",
			Method: Get,
			Middleware: []fiber.Handler{
				middleware.RequiresRole(models.Customer, a.Store),
			},
			Handler: handlers.GetTableOrders(a.Store),
		},
		{
			Path:   "/api/v1/table/bills/:table",
			Method: Get,
			Middleware: []fiber.Handler{
				middleware.RequiresRole(models.Customer, a.Store),
			},
			Handler: handlers.GetTableBills(a.Store),
		},
		// Websocket
		{
			Path:   "/api/v1/staff/manager",
			Method: Get,
			Middleware: []fiber.Handler{
				middleware.RequiresRole(models.Customer, a.Store),
			},
			Handler: server.HandleCreationManager(a.Store),
		},
		{
			Path:   "/api/v1/staff/:table",
			Method: Get,
			Middleware: []fiber.Handler{
				middleware.RequiresRole(models.Customer, a.Store),
			},
			Handler: server.HandleCreationWaiter(a.Store),
		},
		{
			Path:   "/api/v1/order/:table",
			Method: Get,
			Middleware: []fiber.Handler{
				middleware.RequiresRole(models.Customer, a.Store),
			},
			Handler: server.HandleCreationCustomer(a.Store),
		},
		// Bill Handlers
		{
			Path:   "/api/v1/bill",
			Method: Patch,
			Middleware: []fiber.Handler{
				middleware.RequiresRole(models.Customer, a.Store),
			},
			Handler: handlers.PatchBill(a.Store),
		},
		{
			Path:   "/api/v1/bill/:id",
			Method: Patch,
			Middleware: []fiber.Handler{
				middleware.RequiresRole(models.Customer, a.Store),
			},
			Handler: handlers.PatchBills(a.Store),
		},
		{
			Path:   "/api/v1/bill/:id",
			Method: Get,
			Middleware: []fiber.Handler{
				middleware.RequiresRole(models.Customer, a.Store),
			},
			Handler: handlers.GetBill(a.Store),
		},
		{
			Path:   "/api/v1/bill/orders/:id",
			Method: Get,
			Middleware: []fiber.Handler{
				middleware.RequiresRole(models.Customer, a.Store),
			},
			Handler: handlers.GetBillOrders(a.Store),
		},
		// Basket Handlers
		{
			Path:   "/api/v1/baskets/items",
			Method: Post,
			Middleware: []fiber.Handler{
				middleware.RequiresRole(models.Customer, a.Store),
			},
			Handler: handlers.PostBasketItems(a.Store),
		},
		{
			Path:   "/api/v1/baskets/items/:table",
			Method: Post,
			Middleware: []fiber.Handler{
				middleware.RequiresRole(models.Customer, a.Store),
			},
			Handler: handlers.PostBasket(a.Store),
		},
		{
			Path:   "/api/v1/baskets/pay/:table",
			Method: Post,
			Middleware: []fiber.Handler{
				middleware.RequiresRole(models.Customer, a.Store),
			},
			Handler: handlers.PostBasketPayment(a.Store),
		},
	}

	// Assign routes to router
	a.Routes(routes)
	// Start the app and listen at the configured port
	log.Fatalln(a.Fiber.Listen(a.Config.Port))
}

// Shutdown the App
func (a *App) Shutdown() {
	log.Println("app shutting down")
}

// Add a static method to the App
func (a *App) Static(r Route) {
	a.Fiber.Static(r.Path, r.StaticHandler)
}

// Add a Get method to the App
func (a *App) Get(r Route) {
	h := append(r.Middleware, r.Handler)
	a.Fiber.Get(r.Path, h...)
}

// Add a Post method to the App
func (a *App) Post(r Route) {
	h := append(r.Middleware, r.Handler)
	a.Fiber.Post(r.Path, h...)
}

// Add a Patch method to the App
func (a *App) Patch(r Route) {
	h := append(r.Middleware, r.Handler)
	a.Fiber.Patch(r.Path, h...)
}

// Add a Put method to the App
func (a *App) Put(r Route) {
	h := append(r.Middleware, r.Handler)
	a.Fiber.Put(r.Path, h...)
}

// Add a Delete method to the App
func (a *App) Delete(r Route) {
	h := append(r.Middleware, r.Handler)
	a.Fiber.Delete(r.Path, h...)
}

// Add an Options method to the App
func (a *App) Options(r Route) {
	h := append(r.Middleware, r.Handler)
	a.Fiber.Options(r.Path, h...)
}

// Add a Head method to the App
func (a *App) Head(r Route) {
	h := append(r.Middleware, r.Handler)
	a.Fiber.Head(r.Path, h...)
}

// Add a Route to the App
func (a *App) Route(r Route) {
	switch r.Method {
	case Get:
		a.Get(r)
		return
	case Post:
		a.Post(r)
		return
	case Patch:
		a.Patch(r)
		return
	case Put:
		a.Put(r)
		return
	case Delete:
		a.Delete(r)
		return
	case Options:
		a.Options(r)
		return
	case Static:
		a.Static(r)
		return
	case Head:
		a.Head(r)
		return
	}
}

// Add a slice of routes to the Router
func (a *App) Routes(rs []Route) {
	for _, r := range rs {
		a.Route(r)
	}
}
