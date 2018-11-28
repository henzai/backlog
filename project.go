package backlog

import "fmt"

type Projects []Project

type Project struct {
	ID                                int64
	ProjectKey                        string
	Name                              string
	ChartEnabled                      bool
	SubtaskingEnabled                 bool
	ProjectLeaderCanEditProjectLeader bool
	TextFormattingRule                string
	Archived                          bool
}

func (b *BackLog) GetProjects() (*Projects, error) {
	// TODO: クエリパラメーターに対応する
	URL := fmt.Sprintf(b.baseURL, "projects")
	var m Projects
	err := Get(URL, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (b *BackLog) GetProject(IDOrKey string) (*Project, error) {
	URL := fmt.Sprintf(b.baseURL, "projects/"+IDOrKey)
	var m Project
	err := Get(URL, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (b *BackLog) UpdateProject(IDOrKey string) (*Project, error) {
	return nil, nil
}

func (b *BackLog) DeleteProject(IDOrKey string) (bool, error) {
	URL := fmt.Sprintf(b.baseURL, "projects/"+IDOrKey)
	ok, err := delete(URL)
	if err != nil {
		return false, err
	}
	return ok, nil
}

func (b *BackLog) GetProjectImage(IDOrKey string) (interface{}, error) {
	URL := fmt.Sprintf(b.baseURL, "projects/"+IDOrKey+"/image")
	var m Project
	err := Get(URL, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (b *BackLog) AddProjectUser(IDOrKey, user string) (bool, error) {
	return false, nil
}

func (b *BackLog) DeleteProjectUser(IDOrKey, user string) (bool, error) {
	return false, nil
}

func (b *BackLog) GetProjectUser(IDOrKey, user string) (*Users, error) {
	return nil, nil
}

func (b *BackLog) GrantProjectAdmin(IDOrKey, user string) (bool, error) {
	return false, nil
}

func (b *BackLog) GetProjectAdmin(IDOrKey, user string) (bool, error) {
	return false, nil
}

func (b *BackLog) DeleteProjectAdmin(IDOrKey, user string) (bool, error) {
	return false, nil
}
