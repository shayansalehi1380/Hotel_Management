package mainmenu

import (
	"fmt"
	"hotel-management/guest"
	"hotel-management/room"

 )

func MainMenu() {

	for {

		fmt.Println("\n----Hotel Management Syetem ----")
		fmt.Println("1.Register Guest")
		fmt.Println("2.Reserve Room")
		fmt.Println("3.Check out for guest")
		fmt.Println("4.View Room Status")
		fmt.Println("5.Exit")
		fmt.Println("Choose an options...")

		var vooroddi int
		fmt.Scanln(&vooroddi)

		switch vooroddi {
		case 1:
			var name string
			fmt.Println("Enter Guest Name")
			fmt.Scanln(&name)
			guest.RegistrationGuest(name)

		case 2:
			var guestID, roomnumber int
			fmt.Println("Plase Enter the guest id ")
			fmt.Scanln(&guestID)
			guestFinded := guest.FindGuestByID(guestID)
			if guestFinded == nil {
				fmt.Println("Guest Not Found")
				return
			}
			fmt.Println("Plase Enter the  room number ")
			fmt.Scanln(&roomnumber)
			if &roomnumber == nil {

				fmt.Println("room  Not Found")
				return

			}
			room.ReserveRoom(roomnumber, guestFinded)

		case 3:
			var roomnumber int
			fmt.Println("Plase Enter the  room number ")
			fmt.Scanln(&roomnumber)
			room.CheckOut(roomnumber)
		case 4:
			room.ViewRoomStatus()


		case 5:
			fmt.Println("Exiting...")
            return

		}

	}
}
