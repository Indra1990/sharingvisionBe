package usecaseimpl

import (
	"sharingvision_backendtest/app/post/usecase"

	validation "github.com/go-ozzo/ozzo-validation"
)

func articleCreatedValidate(cmd usecase.ArticlesRequest) (err error) {
	err = validation.ValidateStruct(
		&cmd,
		validation.Field(&cmd.Title, validation.Required, validation.Length(20, 0)),
		validation.Field(&cmd.Content, validation.Required, validation.Length(200, 0)),
		validation.Field(&cmd.Category, validation.Required, validation.Length(3, 0)),
		validation.Field(&cmd.Status, validation.Required, validation.In("publish", "draft", "thrash")),
	)

	if err != nil {
		return
	}

	return
}
