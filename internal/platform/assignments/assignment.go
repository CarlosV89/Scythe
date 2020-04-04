package assignments

type DTO struct {
	AssignedProject ProjectAssignmentDTO `json:"project"`
	AssignedClient  ClientAssignmentDTO  `json:"client"`
	AssignedTasks   []TaskAssignmentDTO  `json:"task_assignments"`
}

func (d *DTO) toModel() Assignment {
  
}


type ClientAssignmentDTO struct {
  Id   int    `json:"id"`
	Name string `json:"name"`
}

type TaskAssignmentDTO struct {
  Id   int     `json:"id"`
	Task TaskDTO `json:"task"`
}

type TaskDTO struct {
  Id   int    `json:"id"`
	Name string `json:"name"`
}

type Assignment interface {
	GetAssignedProject() ProjectAssignment
	GetAssignedClient() ClientAssignment
	GetAssignedTasks() []TaskAssignment
}

type assignment struct {
	assignedProject ProjectAssignment
	assignedClient  ClientAssignment
	assignedTasks   []TaskAssignment
}

type ProjectAssignmentDTO struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (d *ProjectAssignmentDTO) toModel() ProjectAssignment {
  
}

type ProjectAssignment interface {
	GetName() string
}

type projectAssignment struct {
	name string
}

func (pa *projectAssignment) GetName() string {
	return pa.name
}

type ClientAssignment interface {
	GetName() string
}

type clientAssignment struct {
	name string
}

func (ca *clientAssignment) GetName() string {
	return ca.name
}

type TaskAssignment interface {
	GetTaskName() string
}

type taskAssignment struct {
	task Task
}

func (ta *taskAssignment) GetTaskName() string {
	return ta.task.GetName()
}

type Task interface {
	GetName() string
}

type task struct {
	name string
}

func (t *task) GetName() string {
	return t.name
}

/*"project_assignments":[
  {
    "id":125066109,
    "is_project_manager":true,
    "is_active":true,
    "use_default_rates":true,
    "budget":null,
    "created_at":"2017-06-26T21:52:18Z",
    "updated_at":"2017-06-26T21:52:18Z",
    "hourly_rate":100.0,
    "project":{
      "id":14308069,
      "name":"Online Store - Phase 1",
      "code":"OS1"
    },
    "client":{
      "id":5735776,
      "name":"123 Industries"
    },
    "task_assignments":[
      {
        "id":155505013,
        "billable":true,
        "is_active":true,
        "created_at":"2017-06-26T21:52:18Z",
        "updated_at":"2017-06-26T21:52:18Z",
        "hourly_rate":100.0,
        "budget":null,
        "task":{
          "id":8083365,
          "name":"Graphic Design"
        }
      },
      {
        "id":155505014,
        "billable":true,
        "is_active":true,
        "created_at":"2017-06-26T21:52:18Z",
        "updated_at":"2017-06-26T21:52:18Z",
        "hourly_rate":100.0,
        "budget":null,
        "task":{
          "id":8083366,
          "name":"Programming"
        }
      },
      {
        "id":155505015,
        "billable":true,
        "is_active":true,
        "created_at":"2017-06-26T21:52:18Z",
        "updated_at":"2017-06-26T21:52:18Z",
        "hourly_rate":100.0,
        "budget":null,
        "task":{
          "id":8083368,
          "name":"Project Management"
        }
      },
      {
        "id":155505016,
        "billable":false,
        "is_active":true,
        "created_at":"2017-06-26T21:52:18Z",
        "updated_at":"2017-06-26T21:54:06Z",
        "hourly_rate":100.0,
        "budget":null,
        "task":{
          "id":8083369,
          "name":"Research"
        }
      }
    ]
  },*/
