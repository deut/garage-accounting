package main

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"go.uber.org/zap"

	"github.com/deut/garage-accounting/config"
	"github.com/deut/garage-accounting/config/translate"
	"github.com/deut/garage-accounting/db"
	"github.com/deut/garage-accounting/internal/app/ui"
	"github.com/deut/garage-accounting/internal/models"
)

const (
	DBName = "garage.db"
)

func main() {
	translate.SetLang(translate.UA)
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	sugar.Debug("Loading app")
	config.Conf = &config.C{}
	config.Conf.Defaults()

	err := db.Connect(config.Conf.DBFileLocation)
	if err != nil {
		sugar.Errorf("db connection error: %v", err)
		os.Exit(1)
	}

	if os.ReadFile(config.Conf.DBFileLocation); err != nil {
		sugar.Info("initializing DB schema")
		db.DB.AutoMigrate(&models.Account{}, &models.Payment{}, &models.Rate{})
	}

	if err != nil {
		sugar.Errorf("account schema error: %v", err)
		os.Exit(1)
	}

	appLayout := ui.NewUI("app.name", 500, 480)
	acc := models.Account{}
	accForm := ui.NewCreateAccountForm(appLayout.MainWindow, &acc)
	listAccs := ui.NewAccountsList(appLayout.MainWindow)

	accFormCanvasObj := accForm.Build()
	accListObj := listAccs.Build()
	accListObj.Resize(fyne.NewSize(accFormCanvasObj.Size().Width, 1000))
	accListObj.Refresh()

	formTab := container.NewTabItem("form", accFormCanvasObj)
	listTab := container.NewTabItem("main", accListObj)

	cont := container.NewAppTabs(listTab, formTab)

	appLayout.SetContent(cont)
	appLayout.ShowMainWindow()

}
