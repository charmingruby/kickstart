package example_usecase

import (
	"github.com/charmingruby/kickstart/internal/core"
	"github.com/charmingruby/kickstart/test/factory"
)

func (s *Suite) Test_GetExample() {
	s.Run("it should be able get an example", func() {
		example, err := factory.MakeExample(s.exampleRepo, "exmaple")
		s.NoError(err)

		items := s.exampleRepo.Items
		s.Equal(1, len(items))

		result, err := s.useCase.GetExampleUseCase(example.ID)
		s.NoError(err)

		s.Equal(items[0].ID, result.ID)
	})

	s.Run("it should be not able to find nonexistent example", func() {
		_, err := factory.MakeExample(s.exampleRepo, "exmaple")
		s.NoError(err)

		items := s.exampleRepo.Items
		s.Equal(1, len(items))

		result, err := s.useCase.GetExampleUseCase("invalid id")
		s.Nil(result)
		s.Error(err)
		s.Equal(core.NewNotFoundErr("example").Error(), err.Error())
	})
}
