package usecase

import (
	"main/entity"

	"github.com/google/uuid"
)

type CreateCourse struct {
	Repository entity.CourseRepository
}

func (c CreateCourse) Execute(input CreateCourseInputDTO) (CreateCourseOutputDTO, error) {
	course := entity.Course{}
	course.ID = uuid.New().String()
	course.Name = input.Name
	course.Description = input.Description
	course.Status = input.Status

	err := c.Repository.Insert(course)

	if err != nil {
		return CreateCourseOutputDTO{}, err
	}

	output := CreateCourseOutputDTO(course)
	return output, nil
}
