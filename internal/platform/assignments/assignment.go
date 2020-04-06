package assignments

type AssignmentDTO struct {
	AssignedProject ProjectAssignmentDTO `json:"project"`
	AssignedClient  ClientAssignmentDTO  `json:"client"`
	AssignedTasks   []TaskAssignmentDTO  `json:"task_assignments"`
}

func (d *AssignmentDTO) ToModel() Assignment {
	var tasks []TaskAssignment
	for _, task := range d.AssignedTasks {
		tasks = append(tasks, task.ToModel())
	}
	return NewAssignment(d.AssignedProject.ToModel(), d.AssignedClient.ToModel(), tasks)
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

func NewAssignment(ap ProjectAssignment, ac ClientAssignment, at []TaskAssignment) Assignment {
	return &assignment{ap, ac, at}
}

func (a *assignment) GetAssignedProject() ProjectAssignment {
	return a.assignedProject
}

func (a *assignment) GetAssignedClient() ClientAssignment {
	return a.assignedClient
}

func (a *assignment) GetAssignedTasks() []TaskAssignment {
	return a.assignedTasks
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
