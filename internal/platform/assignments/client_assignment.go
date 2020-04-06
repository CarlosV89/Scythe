package assignments

type ClientAssignment interface {
	GetName() string
}

type clientAssignment struct {
	id   int32
	name string
}

func (ca *clientAssignment) GetID() int32 {
	return ca.id
}

func (ca *clientAssignment) GetName() string {
	return ca.name
}

func NewClientAssignment(id int32, name string) ClientAssignment {
	return &clientAssignment{id, name}
}

type ClientAssignmentDTO struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func (dto *ClientAssignmentDTO) ToModel() ClientAssignment {
	return NewClientAssignment(dto.ID, dto.Name)
}
