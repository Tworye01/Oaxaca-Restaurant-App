package app

import (
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func MockHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}

func MockMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}

func TestMain(m *testing.M) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	v := m.Run()
	os.Exit(v)
}

func TestNew(t *testing.T) {
	a := New()
	assert.NotNil(t, a)
}

func TestRun(t *testing.T) {
	// a := New()
	// assert.NotPanics(t, a.Run)
}

func TestShutdown(t *testing.T) {
	a := New()
	assert.NotPanics(t, a.Shutdown)
}

func TestGet(t *testing.T) {
	a := New()
	a.Get(Route{Path: "/", Handler: MockHandler()})
}

func TestPost(t *testing.T) {
	a := New()
	a.Post(Route{Path: "/", Handler: MockHandler()})
}

func TestPatch(t *testing.T) {
	a := New()
	a.Patch(Route{Path: "/", Handler: MockHandler()})
}

func TestPut(t *testing.T) {
	a := New()
	a.Get(Route{Path: "/", Handler: MockHandler()})
}

func TestDelete(t *testing.T) {
	a := New()
	a.Delete(Route{Path: "/", Handler: MockHandler()})
}

func TestStatic(t *testing.T) {
	a := New()
	a.Static(Route{Path: "/", StaticHandler: "/web"})
}

func TestOptions(t *testing.T) {
	a := New()
	a.Options(Route{Path: "/", Handler: MockHandler()})
}

func TestHead(t *testing.T) {
	a := New()
	a.Head(Route{Path: "/", Handler: MockHandler()})
}

func TestRoute(t *testing.T) {
	a := New()
	a.Route(Route{Path: "/", Method: Get, Handler: MockHandler()})
	a.Route(Route{Path: "/", Method: Post, Handler: MockHandler()})
	a.Route(Route{Path: "/", Method: Patch, Handler: MockHandler()})
	a.Route(Route{Path: "/", Method: Put, Handler: MockHandler()})
	a.Route(Route{Path: "/", Method: Delete, Handler: MockHandler()})
	a.Route(Route{Path: "/", Method: Options, Handler: MockHandler()})
	a.Route(Route{Path: "/", Method: Head, Handler: MockHandler()})
}

func TestRoutes(t *testing.T) {
	a := New()
	r := []Route{
		{
			Path:    "/",
			Handler: MockHandler(),
		},
	}
	a.Routes(r)
}
