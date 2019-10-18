package dbase

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/braulio94/base_project/models"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"regexp"
	"testing"
	"time"
)

type Suite struct {
	suite.Suite
	DataBase   *gorm.DB
	mock       sqlmock.Sqlmock
	repository ProjectRepository
	project    *models.Project
}

func (s *Suite) SetupSuite(){
	var err      error
	_, s.mock, err = sqlmock.New()
	s.DataBase = GetDB()
	require.NoError(s.T(), err)
	s.DataBase.LogMode(true)
	s.repository = NewRepository(s.DataBase)
}

var (
	id, _ = uuid.NewV4()
	p = models.Project{
		ID: id,
		ProjectName: "Projecto Base",
		Description: "Este projecto é usado como projecto base para outros projectos",
		LastAccessed: time.Now(),
		LastAccessedBy: "Braulio Cassule",
		Status: "Available",
	}
)

func (s *Suite) AfterTest(_,_ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T){
	InitDB()
	suite.Run(t, new(Suite))
}

func (s *Suite) TestRepositoryCreate(){
	s.mock.ExpectQuery(
		regexp.QuoteMeta(`
			INSERT INTO $1 ("id", "project_name", "description", "last_accessed", "last_accessed_by", "status")
			VALUES ($2,$3,$4,$5,$6,$7) RETURNING "project"."id"
		`)).
		WithArgs(p.TableName(), p.ID, p.ProjectName, p.Description, p.LastAccessed, p.LastAccessedBy, p.Status).
		WithArgs(sqlmock.NewRows([]string{"id"}).AddRow(id.String()))
	storeErr := s.repository.Store(&p)
	require.NoError(s.T(), storeErr)
}

func (s *Suite) TestRepositoryGet(){

}