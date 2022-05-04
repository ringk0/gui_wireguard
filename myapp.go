package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	//создаем приложение
	myapp := app.New()
	myForm := myapp.NewWindow("GUI для VPN-Wireguard (обход блокировок) v.1.0")
	myForm.Resize(fyne.NewSize(600, 300))
	myForm.SetFixedSize(true)

	//установка иконки окна
	myIcon, _ := fyne.LoadResourceFromPath("anonim.png")
	myForm.SetIcon(myIcon)
	myForm.CenterOnScreen()

	//вывод текста "Сервер"
	textServer := canvas.NewText("Выбор сервера", color.White)
	textServer.Move(fyne.NewPos(18, 50))
	textServer.TextSize = 12
	textServer.TextStyle = fyne.TextStyle{Italic: true}

	//1 строка
	label1 := widget.NewLabel("")
	label1.Move(fyne.NewPos(160, 80))
	label1.TextStyle = fyne.TextStyle{Italic: true}
	//2 строка
	label2 := widget.NewLabel("")
	label2.Move(fyne.NewPos(160, 116))
	label2.TextStyle = fyne.TextStyle{Italic: true}
	//3 строка
	label3 := widget.NewLabel("")
	label3.Move(fyne.NewPos(160, 152))
	label3.TextStyle = fyne.TextStyle{Italic: true}

	//вывод текста о приложении
	textAboutApp := canvas.NewText("VPN для Обход блокировок с помощью протокола Wireguard", color.White)
	//вывод инфы в конкретном месте
	textAboutApp.Move(fyne.NewPos(66, 12))
	//создаем группу из радиокнопок для выбора сервера
	chk_servers := widget.NewRadioGroup([]string{"Estonia", "Latvia", "Luxemburg"}, func(string) {

	})
	//отображаем в нужном месте формы
	chk_servers.Move(fyne.NewPos(10, 80))
	//выбор по умолчанию
	//chk_servers.SetSelected("Estonia")

	//выключаем выбранный сервер
	btn_off := widget.NewButton("OFF", func() {

	})
	btn_off.Resize(fyne.NewSize(100, 30))
	btn_off.Move(fyne.NewPos(300, 220))
	btn_off.Disable()

	//включаем выбранный сервер
	btn_on := widget.NewButton("ON", func() {
		btn_off.Enable()
		switch string(chk_servers.Selected) {
		case "Estonia":
			{
				label1.SetText("команда в терминале 'wg-quick up estonia'")
				label2.SetText("")
				label3.SetText("")
			}
		case "Latvia":
			{
				label2.SetText("команда в терминале 'wg-quick up latvia'")
				label1.SetText("")
				label3.SetText("")
			}
		case "Luxemburg":
			{
				label3.SetText("команда в терминале 'wg-quick up luxem'")
				label2.SetText("")
				label1.SetText("")
			}
		default:
			label1.SetText("Сервер не выбран. Ошибка.")
		}
	})
	btn_on.Resize(fyne.NewSize(100, 30))
	btn_on.Move(fyne.NewPos(200, 220))

	//контейнер виджетов для отображения на форме
	box1 := container.NewWithoutLayout(
		chk_servers, btn_on,
		btn_off,
		textServer,
		label1,
		label2,
		label3,
		textAboutApp)
	myForm.SetContent(box1)

	//вызываем окно
	myForm.ShowAndRun()
}
