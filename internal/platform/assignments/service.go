package assignments

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/vanor89/scythe/internal/platform/harvest"
)

// Service for handling user assignments
type Service interface {
	FindCurrentUserAssignments() []Assignment
}

type service struct {
	h harvest.Harvester
}

// NewService is our Assignments Service constructor.
func NewService() Service {
	return &service{}
}

func (s *service) getEndpoint(userID string) string {
	return fmt.Sprintf("users/%s/project_assignments", userID)
}

func (s *service) FindCurrentUserAssignments() []Assignment {
	var userAssignments []Assignment
	var dtos []AssignmentDTO

	body := s.h.Request(
		"GET",
		nil,
		nil,
		s.getEndpoint("me"),
	)

	if err := json.Unmarshal(body, &dtos); err != nil {
		log.Fatal(err)
	}

	for _, assignmentDto := range dtos {
		userAssignments = append(userAssignments, assignmentDto.ToModel())
	}

	return userAssignments
}
