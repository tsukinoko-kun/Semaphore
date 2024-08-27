package main

import (
	"embed"
	"runtime"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:             "Semaphore",
		WindowStartState:  options.Maximised,
		HideWindowOnClose: runtime.GOOS == "darwin",
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour:         &options.RGBA{R: 0, G: 0, B: 0, A: 1},
		OnStartup:                app.startup,
		EnableDefaultContextMenu: false,
		DragAndDrop: &options.DragAndDrop{
			DisableWebViewDrop: true,
		},
		Bind: []interface{}{
			app,
		},
		Windows: &windows.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			DisableWindowIcon:    true,
			IsZoomControlEnabled: false,
			DisablePinchZoom:     true,
			Theme:                windows.Dark,
			CustomTheme: &windows.ThemeSettings{
				DarkModeTitleText:         0x00FFFFFF,
				DarkModeTitleTextInactive: 0x00A1A1A1,
			},
			WebviewGpuIsDisabled: false,
			EnableSwipeGestures:  false,
		},
		Mac: &mac.Options{
			TitleBar:             mac.TitleBarDefault(),
			Appearance:           mac.NSAppearanceNameAccessibilityHighContrastVibrantDark,
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			DisableZoom:          true,
			About: &mac.AboutInfo{
				Title:   "Semaphore",
				Message: "not just an email client\n\nyour messenger for the professional world",
			},
		},
		Linux: &linux.Options{
			WindowIsTranslucent: false,
			ProgramName:         "Semaphore",
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
