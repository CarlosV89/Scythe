package assignments

type DTO struct {
	ID              string `json:"id,omitempty"`
	AssignedProject Project
	AssignedClient  Client
	AssignedTasks   Tasks
}
