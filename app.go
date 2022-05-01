package main

import (
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"log"
	"time"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GreetAsyncViaChannel() <-chan string {
	messages := make(chan string)
	go func() {
		defer close(messages)
		for _, m := range []string{"Hello A (from server)", "Hello B (from server)", "Hello C (from server)", "DONE"} {
			log.Println(m)
			messages <- m
			time.Sleep(1 * time.Second)
		}
	}()

	return messages // FAT | json: unsupported type: <-chan string
}

func (a *App) GreetAsyncViaEvent() {
	go func() {
		for _, m := range []string{"Hello A (from server)", "Hello B (from server)", "Hello C (from server)", "DONE"} {
			runtime.EventsEmit(a.ctx, "rcv:greet", m)
			log.Println(m)
			time.Sleep(1 * time.Second)
		}
	}()
}
