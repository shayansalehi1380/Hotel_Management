package main

import (
	mainmenu "hotel-management/main_menu"
	"hotel-management/room"
)

func main() {
	room.InitializeRoom(15)
	mainmenu.MainMenu()
}