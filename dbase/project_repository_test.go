package dbase

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/braulio94/base_project/models"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

var (
	id, _ = uuid.NewV4()
	p     = models.Project{
		ID:             id,
		ProjectName:    "Projecto Base",
		Description:    "Este projecto é usado como projecto base para outros projectos",
		LastAccessed:   time.Now().String(),
		LastAccessedBy: "Braulio Cassule",
		Status:         "Available",
	}
)

type Suite struct {
	suite.Suite
	DataBase   *gorm.DB
	mock       sqlmock.Sqlmock
	repository ProjectRepository
	project    *models.Project
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)
	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)
	s.DataBase, err = gorm.Open("postgres", db)
	require.NoError(s.T(), err)
	s.DataBase.LogMode(true)
	s.repository = NewRepository(s.DataBase)
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestRepositoryCreate() {
	s.mock.ExpectQuery(
		fmt.Sprintf(
			"INSERT INTO %s ('id', 'project_name', 'description', 'last_accessed', 'last_accessed_by', 'status') "+
				"VALUES (%s, %s, %s, %s, %s, %s)",
			p.TableName(), p.ID.String(), p.ProjectName, p.Description, p.LastAccessed, p.LastAccessedBy, p.Status)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
			AddRow(id.String()))
	err := s.repository.Store(&p)
	require.NoError(s.T(), err)
}
