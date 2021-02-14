package service

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/vibhugarg123/book-my-show/appcontext"
	"github.com/vibhugarg123/book-my-show/config"
	"github.com/vibhugarg123/book-my-show/constants"
	"github.com/vibhugarg123/book-my-show/entities"
	"github.com/vibhugarg123/book-my-show/repository"
	"testing"
	"time"
)

type movieServiceTestSuite struct {
	suite.Suite
	mockMovieRepository *repository.MockMovieRepository
	service             MovieService
}

func (suite *movieServiceTestSuite) SetupTest() {
	config.LoadTestConfig()
	appcontext.Init()
	mockController := gomock.NewController(suite.T())
	suite.mockMovieRepository = repository.NewMockMovieRepository(mockController)
	suite.service = &movieService{
		movieRepository: suite.mockMovieRepository,
	}
}

func (suite *movieServiceTestSuite) TestAddMovieReturnsErrorWhenMovieObjectIsMissing() {
	movie := []entities.Movie{
		{
		},
	}
	movieReturned, err := suite.service.Add(movie[0])
	assert.Equal(suite.T(), entities.Movie{}, movieReturned)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), errors.Wrap(errors.New(constants.REQUEST_INVALID), constants.EMPTY_REQUEST_BODY).Error(), err.Error())
}

func (suite *movieServiceTestSuite) TestAddMovieReturnsErrorWhenMovieNameIsMissing() {
	movie := []entities.Movie{
		{
			DirectorName: "Salman Khan",
			ReleaseDate:  time.Date(2019, time.February, 1, 0, 0, 0, 0, time.UTC),
			IsActive:     true,
		},
	}
	movieReturned, err := suite.service.Add(movie[0])
	assert.Equal(suite.T(), entities.Movie{}, movieReturned)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), errors.Wrap(errors.New(constants.REQUEST_INVALID), constants.MOVIE_NAME_MANDATORY).Error(), err.Error())
}

func (suite *movieServiceTestSuite) TestAddMovieReturnsErrorWhenMovieDirectorNameIsMissing() {
	movie := []entities.Movie{
		{
			Name:        "ABCD",
			ReleaseDate: time.Date(2019, time.February, 1, 0, 0, 0, 0, time.UTC),
			IsActive:    true,
		},
	}
	movieReturned, err := suite.service.Add(movie[0])
	assert.Equal(suite.T(), entities.Movie{}, movieReturned)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), errors.Wrap(errors.New(constants.REQUEST_INVALID), constants.DIRECTOR_NAME_MANDATORY).Error(), err.Error())
}

func (suite *movieServiceTestSuite) TestAddMovieReturnsErrorWhenMovieReleaseDateIsMissing() {
	movie := []entities.Movie{
		{
			Name:         "ABCD",
			DirectorName: "SVG Production",
			IsActive:     true,
		},
	}
	movieReturned, err := suite.service.Add(movie[0])
	assert.Equal(suite.T(), entities.Movie{}, movieReturned)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), errors.Wrap(errors.New(constants.REQUEST_INVALID), constants.RELEASE_DATE_MANDATORY).Error(), err.Error())
}

func (suite *movieServiceTestSuite) TestAddMovieReturnsErrorWhenMovieAlreadyExists() {
	movie := []entities.Movie{
		{
			Name:         "ABCD",
			DirectorName: "SVG Production",
			ReleaseDate:  time.Date(2019, time.February, 1, 0, 0, 0, 0, time.UTC),
			IsActive:     true,
		},
	}
	suite.mockMovieRepository.EXPECT().FetchMovieByName(gomock.Any()).Return(movie, nil)
	movieReturned, err := suite.service.Add(movie[0])
	assert.Equal(suite.T(), entities.Movie{}, movieReturned)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), errors.Wrap(errors.New(constants.MOVIE_CREATION_FAILED), fmt.Sprintf(constants.MOVIE_ALREADY_EXISTS, movie[0].Name)).Error(), err.Error())
}

func (suite *movieServiceTestSuite) TestAddMovieWhenMovieIsCreatedSuccessfully() {
	movie := []entities.Movie{
		{
			Name:         "ABCD",
			DirectorName: "SVG Production",
			ReleaseDate:  time.Date(2019, time.February, 1, 0, 0, 0, 0, time.UTC),
			IsActive:     true,
		},
	}
	suite.mockMovieRepository.EXPECT().FetchMovieByName(gomock.Any()).Return(nil, nil)
	suite.mockMovieRepository.EXPECT().InsertMovie(gomock.Any()).Return(nil)
	movieReturned, err := suite.service.Add(movie[0])
	assert.Equal(suite.T(), movie[0], movieReturned)
	assert.Nil(suite.T(), err)
}

func TestMovieServiceTestSuite(t *testing.T) {
	suite.Run(t, new(movieServiceTestSuite))
}
