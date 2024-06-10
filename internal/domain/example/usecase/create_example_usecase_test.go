package usecase

import (
	"github.com/charmingruby/kickstart/internal/domain/example/dto"
	"github.com/charmingruby/kickstart/internal/validation"
)

func (s *Suite) Test_CreateExample() {
	s.Run("it should be able to create an example", func() {
		dto := dto.CreateExampleDTO{Name: "Dummy Name"}

		err := s.exampleService.CreateExample(dto)

		items := s.exampleRepo.Items

		s.NoError(err)
		s.Equal(1, len(items))
		s.Equal(items[0].Name, dto.Name)
	})

	s.Run("it should be not able to create an example with validation error", func() {
		dto := dto.CreateExampleDTO{Name: ""}

		err := s.exampleService.CreateExample(dto)

		s.Error(err)
		s.Equal(validation.ErrMinLength("name", "3"), err.Error())
	})
}
