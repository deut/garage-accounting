package main

import (
	"os"

	"fyne.io/fyne/v2/container"
	"go.uber.org/zap"

	"github.com/deut/garage-accounting/db"
	"github.com/deut/garage-accounting/internal/app/ui"
	"github.com/deut/garage-accounting/internal/models"
)

const (
	DBName = "garage.db"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	err := db.Connect("garage.db")
	if err != nil {
		sugar.Errorf("db connection error: %v", err)
		os.Exit(1)
	}

	sugar.Info("initializing DB schema")
	a := models.Account{}
	err = a.InitSchema()
	if err != nil {
		sugar.Errorf("intet account schema error: %v", err)
		os.Exit(1)
	}

	appLayout := ui.NewLayout("app.name")
	acc := models.Account{}
	accForm := ui.NewCreateAccountForm(appLayout.MainWindow, &acc)
	listAccs := ui.NewAccountsList(appLayout.MainWindow)

	tabs := container.NewAppTabs(
		container.NewTabItem("create.account", accForm.Build()),
		container.NewTabItem("list.account", listAccs.Build()),
	)

	appLayout.SetContent(tabs)
	appLayout.ShowMainWindow()
}
