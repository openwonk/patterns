package patterns

//	Design: Memento
//	Intent: Restore object back to a previous state by externalizing the object’s internal state
//	Scenario: Take apart and recontruct phone (say, for repair)

type Parts string
type Picture Phone
type Phone struct {
	parts map[int]Parts
}

func (this *Phone) RemovePart(part int) {
	this.parts[part] = "", false
}

func (this *Phone) TakePicture() *Picture {
	return NewPicture(this.parts)
}

func (this *Phone) RestoreFromPicture(pic *Picture) {
	this.parts = pic.getSavedParts()
}

func (this *Picture) getSavedParts() map[int]Parts {
	return this.parts
}

func NewPicture(partsToSave map[int]Part) *Picture {
	this := new(Picture)
	this.parts = make(map[int]Part)
	for index, part := range partsToSave {
		this.parts[index] = part
	}
	return this
}

func NewPhone() *Phone {
	this := new(Phone)
	this.parts = make(map[int]Parts)
	this.parts[0] = Parts{"body"}
	this.parts[1] = Parts{"display"}
	this.parts[2] = Parts{"keyboard"}
	this.parts[3] = Parts{"mainboard"}
	return this

}

func Caretaker() {
	phone := NewPhone()
	pictures := make([]*Picture, 3)
	pictures[0] = phone.TakePicture()

	// Remove body and keyboard
	phone.RemovePart(0)
	phone.RemovePart(2)

	pictures[1] = phone.TakePicture()
	phone.RemovePart(1)

	pictures[2] = phone.TakePicture()
	phone.RemovePart(3)

	phone.RestoreFromPicture(pictures[2])
	phone.RestoreFromPicture(pictures[1])
	phone.RestoreFromPicture(pictures[0])

}
