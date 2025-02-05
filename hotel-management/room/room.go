package room

import (
	"fmt"
	"hotel-management/guest"
)

type Room struct {
	Number   int
	Occupied bool
	Guest    *guest.Guest
}

var room []Room

func InitializeRoom(numrooms int) {

	for i := 1; i <= numrooms; i++ {
		room = append(room, Room{Number: i, Occupied: false, Guest: nil})
	}
}

func ReserveRoom(number int, guest *guest.Guest) {
	for i := range room {

		if room[i].Number == number {
			if room[i].Occupied {

				fmt.Println("this room is already reserved")
			} else {
				room[i].Occupied = true
				room[i].Guest = guest
				fmt.Printf("Room %d is already reserved for guest %s . \n", number, guest.Name)
			}
			return
		}

	}

	fmt.Println("Room Not Found")
}

func CheckOut(RoomNumber int) {

	for i := range room {

		if room[i].Number == RoomNumber {

			if room[i].Occupied == true {
				fmt.Printf("Room %d is Checked out", RoomNumber)
				room[i].Occupied = false
				room[i].Guest = nil
			} else {
				fmt.Println("This room is not occupied")
			}

		} else {
			fmt.Println("Room Not Found")
		}

		return
	}

}

func ViewRoomStatus() {
	for _, rm := range room {

		status := "Available"

		if rm.Occupied == true {
			status = fmt.Sprintf("Occupied %s", rm.Guest.Name)
		}

		fmt.Printf("Room %d is %s\n", rm.Number, status)
	}
}
