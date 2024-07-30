package usecase

import (
	"github.com/charmingruby/kickstart/internal/common/core"
	"github.com/charmingruby/kickstart/internal/example/domain/dto"
)

func (s *Suite) Test_CreateExample() {
	s.Run("it should be able to create an example", func() {
		dto := dto.CreateExampleUseCaseDTO{Name: "Dummy Name"}

		err := s.useCase.CreateExampleUseCase(dto)

		items := s.exampleRepo.Items

		s.NoError(err)
		s.Equal(1, len(items))
		s.Equal(items[0].Name, dto.Name)
	})

	s.Run("it should be not able to create an example with core error", func() {
		dto := dto.CreateExampleUseCaseDTO{Name: ""}

		err := s.useCase.CreateExampleUseCase(dto)

		s.Error(err)
		s.Equal(core.ErrMinLength("name", "3"), err.Error())
	})
}
