package dockerhub

import "github.com/devlikeapro/patrons-perks/internal/patron"

type DockerHubPerkInfo struct {
	Name       string
	PasswordID string
}

func (p *DockerHubPerkInfo) Info() interface{} {
	return p
}

type DockerHubPerk struct{}

func (p *DockerHubPerk) GiveTo(pat *patron.Patron) error {
	passwordID, err := createDockerHubPassword(pat)
	if err != nil {
		return err
	}
	pat.Perks["docker-hub"] = &DockerHubPerkInfo{Name: "Docker Hub", PasswordID: passwordID}
	return nil
}

func (p *DockerHubPerk) TakeFrom(pat *patron.Patron) error {
	delete(pat.Perks, "docker-hub")
	return nil
}

func (p *DockerHubPerk) IsEnabledFor(pat *patron.Patron) bool {
	_, ok := pat.Perks["docker-hub"]
	return ok
}

func createDockerHubPassword(pat *patron.Patron) (string, error) {
	// TODO: Implement the logic to create a Docker Hub password for the patron
	return "123456", nil
}
