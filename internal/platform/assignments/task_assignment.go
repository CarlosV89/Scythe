package assignments

type TaskAssignment interface {
	GetTask() Task
}

type taskAssignment struct {
	id   int32
	task Task
}

func (ta *taskAssignment) GetID() int32 {
	return ta.id
}

func (ta *taskAssignment) GetTask() Task {
	return ta.task
}

func NewTaskAssignment(id int32, t Task) TaskAssignment {
	return &taskAssignment{id, t}
}

type TaskAssignmentDTO struct {
	ID   int32   `json:"id"`
	Task TaskDTO `json:"task"`
}

func (d *TaskAssignmentDTO) ToModel() TaskAssignment {
	return NewTaskAssignment(d.ID, d.Task.ToModel())
}

type Task interface {
	GetName() string
}

type task struct {
	id   int32
	name string
}

func (t *task) GetID() int32 {
	return t.id
}

func (t *task) GetName() string {
	return t.name
}

func NewTask(id int32, name string) Task {
	return &task{id, name}
}

type TaskDTO struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func (d *TaskDTO) ToModel() Task {
	return NewTask(d.ID, d.Name)
}
