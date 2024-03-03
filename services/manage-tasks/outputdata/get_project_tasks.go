package outputdata

import (
	entities "mvp-2-spms/domain-aggregate"
	"strconv"
	"time"
)

type GetProjectTasks struct {
	Tasks []getProjectTasksData `json:"tasks"`
}

func MapToGetProjectTasks(tasks []entities.Task) GetProjectTasks {
	outputProjects := []getProjectTasksData{}
	for _, task := range tasks {
		id, _ := strconv.Atoi(task.Id)
		outputProjects = append(outputProjects,
			getProjectTasksData{
				Id:          id,
				Name:        task.Name,
				Description: task.Description,
				Deadline:    task.Deadline,
				Status:      task.Status.String(),
			})
	}
	return GetProjectTasks{
		Tasks: outputProjects,
	}
}

type getProjectTasksData struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	Status      string    `json:"status"`
}
