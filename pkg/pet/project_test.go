package pet_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/viqueen/go-template/pkg/pet/mocks"
	"testing"
)

func TestProject(t *testing.T) {
	project := mocks.NewProject(t)
	project.EXPECT().Get().Return("project")
	assert.Equal(t, "project", project.Get())
}
