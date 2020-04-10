package users

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/vanor89/scythe/internal/platform/harvest"
)

const endpoint = "users"

type Service struct {
	h harvest.Harvester
}

func (s *Service) FindCurrentUser() User {
	var user User
	var dto UserDTO

	body := s.h.Request(
		"GET",
		nil,
		nil,
		fmt.Sprintf("%s/me", endpoint),
	)

	err := json.Unmarshal(body, &dto)
	user, err = dto.ToModel()

	if err != nil {
		log.Fatal(err)
	}

	return user
}
