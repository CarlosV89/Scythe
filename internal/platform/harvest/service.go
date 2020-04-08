package assignments

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/vanor89/scythe/internal/platform/harvest"
)

type Service struct {
	h harvest.Harvester
}

func (s *Service) getEndpoint(userId string) string {
	return fmt.Sprintf("users/%s/project_assignments", userId);
}

func (s *Service) FindCurrentUsersAssignments() Assignment[] {
	var userAssignments []Assignment
	var dto AssignmentDTO

	body := s.h.Request(
		"GET",
		nil,
		nil,
		s.getEndpoint("me"),
	)

	err := json.Unmarshal(body, &dto)
	assignment, err = &dto.toModel(&dto)

	if err != nil {
		log.Fatal(err)
	}

	return user
}
