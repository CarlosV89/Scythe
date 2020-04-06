package assignments

type ProjectAssignment interface {
	GetID() int32
	GetName() string
}

type projectAssignment struct {
	id   int32
	name string
}

func NewProjectAssignment(id int32, name string) ProjectAssignment {
	return &projectAssignment{id, name}
}

func (pa *projectAssignment) GetName() string {
	return pa.name
}

func (pa *projectAssignment) GetID() int32 {
	return pa.id
}

type ProjectAssignmentDTO struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func (d *ProjectAssignmentDTO) ToModel() ProjectAssignment {
	return NewProjectAssignment(d.ID, d.Name)
}
