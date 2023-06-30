package models

import "gorm.io/gorm"

// `projects` table models

// Projects table row
type Project struct {
	gorm.Model

	Name          string `gorm:"column:name"`
	Description   string `gorm:"size:2550;column:description"`
	Tags          string `gorm:"column:tags"`
	RepoLink      string `gorm:"column:repo_link"`
	CommChannel   string `gorm:"column:comm_channel"`
	ReadmeLink    string `gorm:"column:readme_link"`
	ProjectStatus bool   `gorm:"default:false;column:project_status"`

	// for stats
	LastPullTime int64 `gorm:"column:last_pull_time"`

	// stats table
	CommitCount  uint `gorm:"column:commit_count"`
	PullCount    uint `gorm:"column:pull_count"`
	AddedLines   uint `gorm:"column:added_lines"`
	RemovedLines uint `gorm:"column:removed_lines"`

	// list of students who contributed to the project (a string of usernames separated by comma(,))
	Contributors string `gorm:"column:contributors"`

	// list of URLs to PRs contributed to the project (a string of links separated by comma(,))
	Pulls string `gorm:"column:pulls"`

	// foreign keys
	Mentor_id          int32  `gorm:"column:mentor_id"`
	Mentor             Mentor `gorm:"ForeignKey:Mentor_id"`
	SecondaryMentor_id int32  `gorm:"column:secondary_mentor_id"`
	SecondaryMentor    Mentor `gorm:"ForeignKey:SecondaryMentor_id"`
}
