package guest

type Guest struct {
	ID   int
	Name string
}

var GuestIdCounter int
var guestList []Guest

func RegistrationGuest(name string) Guest {
	GuestIdCounter++
	guest := Guest{ID: GuestIdCounter, Name: name}
	guestList = append(guestList, guest)
	return guest

}

func FindGuestByID(ID int) *Guest {

	for i:=range guestList{
		if guestList[i].ID == ID{
            return &guestList[i]
        }
	}
	return nil

}
