package worker

import (
	"PharmaProject/task"
)

//	Machinery ..

var add_med_tasks = make(map[string]interface{})

func init() {
	add_med_tasks[task.TaskAddMedicine] = task.AddMedicine
}
