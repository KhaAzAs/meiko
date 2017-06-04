package assignment

import (
	"database/sql"
	"fmt"

	"github.com/melodiez14/meiko/src/util/conn"
)

// GetByCourseID is the function for get all data from assignments table using courses_id
func GetByCourseID(courseID int64) ([]Assignment, error) {
	var assignments []Assignment
	query := fmt.Sprintf(queryGetByCourseID, courseID)
	err := conn.DB.Select(&assignments, query)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return assignments, nil
}

// GetIncompleteByUserID is the function for get all incomplete assignment from assignments table which is not exist in p_users_assignments using users_id
func GetIncompleteByUserID(userID int64) ([]Assignment, error) {
	var assignments []Assignment
	query := fmt.Sprintf(queryGetIncompleteByUserID, userID, userID)
	err := conn.DB.Select(&assignments, query)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return assignments, nil
}

// GetCompleteIDByUserID is the function for get assignments_id from p_users_assignments table using users_id
func GetCompleteIDByUserID(userID int64) ([]int64, error) {
	var assignmentsID []int64
	query := fmt.Sprintf(queryGetCompleteIDByUserID, userID)
	err := conn.DB.Select(&assignmentsID, query)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return assignmentsID, nil
}

// GetGradeParameterByCourseID is the function for get all grade parameters from grade_parameters table using courses_id
func GetGradeParameterByCourseID(courseID int64) ([]GradeParameter, error) {
	var gradeParameters []GradeParameter
	query := fmt.Sprintf(queryGetGradeParametersByCourseID, courseID)
	err := conn.DB.Select(&gradeParameters, query)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return gradeParameters, nil
}

// GetIDByGradeParameterID is the function for get assignments_id from assignments table using grade_paramenters_id
func GetIDByGradeParameterID(gradeParameterID int64) ([]int64, error) {
	var assignmentsID []int64
	query := fmt.Sprintf(queryGetIDByGradeParametersID, gradeParameterID)
	err := conn.DB.Select(&assignmentsID, query)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return assignmentsID, nil
}

// GetCompleteByUserIDGradeParameterID is the function for get all data from p_users_assignments table using users_id and assignments_id
func GetCompleteByUserIDGradeParameterID(userID, gradeParameterID int64) ([]CompleteAssignment, error) {
	var completes []CompleteAssignment
	query := fmt.Sprintf(queryGetCompleteByUserIDGradeParameterID, gradeParameterID, userID)
	err := conn.DB.Select(&completes, query)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return completes, nil
}
