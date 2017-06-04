package assignment

const queryGetIncompleteByUserID = `
	SELECT
		id,
		name,
		status,
		upload_date,
		due_date
	FROM
		assigments
	WHERE
		EXISTS (
			SELECT 
				id
			FROM
				grade_parameters
			WHERE
				EXISTS (
					SELECT
						courses_id
					FROM
						p_users_courses
					WHERE
						users_id = (%d)
				)
		) AND id NOT IN (
			SELECT
				assignments_id
			FROM
				p_users_assignments
			WHERE
				users_id = (%d)
		);
`

const queryGetByCourseID = `
	SELECT
		id,
		name,
		status,
		upload_date,
		due_date
	FROM
		assignments
	WHERE
		EXISTS (
			SELECT
				id
			FROM
				grade_parameters
			WHERE
				courses_id = (%d)
		);
`

const queryGetCompleteIDByUserID = `
	SELECT
		assignments_id
	FROM
		p_users_assignments
	WHERE
		users_id = (%d);
`

const queryGetCompleteByUserIDGradeParameterID = `
	SELECT
		assignments_id,
		users_id,
		description,
		score
	FROM
		p_users_assignments
	WHERE
		assignments_id IN (
			SELECT
				id
			FROM
				assignments
			WHERE
				grade_parameters_id = (%d)
		) AND
		users_id = (%d) AND
		score IS NOT NULL;
`

const queryGetGradeParametersByCourseID = `
	SELECT
		id,
		type,
		percentage
	FROM
		grade_parameters
	WHERE
		courses_id = (%d);
`

const queryGetIDByGradeParametersID = `
	SELECT
		id
	FROM
		assignments
	WHERE
		grade_parameters_id = (%d)
`
