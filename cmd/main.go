package main

import (
	"bmi-calculator/helpers"
	"bmi-calculator/validators"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"math"
	"strconv"
)

func main() {

	application := app.New()

	// Создание окна программы
	window := application.NewWindow("Калькулятор ИМТ")
	window.Resize(helpers.WindowSize)
	window.SetFixedSize(true)

	application.Settings().SetTheme(&helpers.CustomTheme{})

	// Обьявление заголовка в программе
	titleLabel := canvas.NewText("IWEIGHT", helpers.TextColor)
	titleLabel.TextStyle.Bold = true
	titleLabel.TextSize = helpers.TextSize

	// Обьявление заголовка для ввода веса
	weightLabel := canvas.NewText("Вес в кг", helpers.TextColor)
	weightLabel.TextSize = helpers.TextSize

	// Обьявление виджета для ввода данных
	weightEntry := widget.NewEntry()
	weightEntry.SetPlaceHolder("47.21")

	// Обьявление заголовка для ввода роста
	heightLabel := canvas.NewText("Рост в см", helpers.TextColor)
	heightLabel.TextSize = helpers.TextSize

	// Обьявление виджера для ввода роста
	heightEntry := widget.NewEntry()
	heightEntry.SetPlaceHolder("47.21")

	// Обьявление заголовка для выбора пола
	sexLabel := canvas.NewText("Выберите пол", helpers.TextColor)
	sexLabel.TextSize = helpers.TextSize

	// Обьявление виджета для выбора пола
	sexSelect := widget.NewSelect([]string{"Мужчина", "Женщина"}, func(s string) {})
	sexSelect.PlaceHolder = "Выберите один"

	// Обьявление кнопки и ее функционал для подсчета ИМТ по введенным данным
	calculateButton := widget.NewButton("Рассчитать ИМТ", func() {
		height := heightEntry.Text
		err := validators.NumValidator(height)
		if err != nil {
			dialog.ShowError(err, window)
			return
		}

		weight := weightEntry.Text
		err = validators.NumValidator(weight)
		if err != nil {
			dialog.ShowError(err, window)
			return
		}

		validHeight, _ := strconv.ParseFloat(height, 64)
		validWeight, _ := strconv.ParseFloat(weight, 64)

		bmi := validWeight / math.Pow(validHeight/100, 2)

		// Создание второго окна
		ResultWindow(application, bmi)

	})

	titleLayout := container.New(
		layout.NewCenterLayout(),
		titleLabel,
	)

	buttonLayout := container.New(
		layout.NewHBoxLayout(),
		calculateButton,
	)
	buttonWrapper := container.New(
		layout.NewCenterLayout(),
		buttonLayout,
	)

	// Вывод всех созданных виджетов в окно программы
	window.SetContent(container.NewVBox(
		titleLayout,
		layout.NewSpacer(),
		weightLabel,
		weightEntry,
		heightLabel,
		heightEntry,
		sexLabel,
		sexSelect,
		layout.NewSpacer(),
		buttonWrapper,
		layout.NewSpacer(),
	))

	window.ShowAndRun()
}

// Создание второго окна с результатами расчетов
func ResultWindow(app fyne.App, result float64) {
	window := app.NewWindow("Результат")
	window.Resize(helpers.WindowSize)
	window.SetFixedSize(true)

	resultLabel := canvas.NewText(fmt.Sprintf("Ваш ИМТ %.2f", result), helpers.TextColor)
	resultLabel.TextSize = helpers.TextSize

	resultText := canvas.NewText(helpers.BmiToText(result), helpers.TextColor)
	resultText.TextSize = helpers.TextSize

	closeButton := widget.NewButton("Закрыть", func() {
		window.Close()
	})

	resultLabelWrapper := container.NewCenter()
	resultLabelWrapper.Add(resultLabel)

	resultTextWrapper := container.NewCenter()
	resultTextWrapper.Add(resultText)

	window.SetContent(container.New(
		layout.NewVBoxLayout(),
		resultLabelWrapper,
		resultTextWrapper,
		layout.NewSpacer(),
		closeButton,
	))

	window.Show()
}
