package usecase

import (
	"github.com/charmingruby/kickstart/internal/domain/example/entity"
	"github.com/charmingruby/kickstart/internal/validation"
)

func (s *Suite) Test_GetExample() {
	s.Run("it should be able get an example", func() {
		example, _ := entity.NewExample("Dummy Name")

		err := s.exampleRepo.Store(example)
		s.NoError(err)

		items := s.exampleRepo.Items
		s.Equal(1, len(items))

		result, err := s.exampleService.GetExample(example.ID)
		s.NoError(err)

		s.Equal(items[0].ID, result.ID)
	})

	s.Run("it should be not able to find nonexistent example", func() {
		example, _ := entity.NewExample("Dummy Name")

		err := s.exampleRepo.Store(example)
		s.NoError(err)

		items := s.exampleRepo.Items
		s.Equal(1, len(items))

		result, err := s.exampleService.GetExample("invalid id")
		s.Nil(result)
		s.Error(err)
		s.Equal(validation.NewNotFoundErr("example").Error(), err.Error())
	})
}
