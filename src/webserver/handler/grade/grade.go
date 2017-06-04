package grade

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/melodiez14/meiko/src/module/assignment"
	"github.com/melodiez14/meiko/src/module/course"
	"github.com/melodiez14/meiko/src/util/auth"
	"github.com/melodiez14/meiko/src/webserver/template"
)

func GetSummaryHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var res []profileSummaryResponse
	u := r.Context().Value("User").(*auth.User)

	// get all enrolled course using using userID
	courses, err := course.GetByUserID(u.ID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError).
			AddError(err.Error()))
		return
	}

	// if there is no enrolled course
	if len(courses) < 1 {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusOK).
			SetData(res))
	}

	// iterate all courses to get the summary
	for _, v := range courses {

		pSummary := profileSummaryResponse{
			CourseName: v.Name,
			Parameter:  []profileParameter{},
			UCU:        v.UCU,
		}

		// get all grade parameters
		gps, err := assignment.GetGradeParameterByCourseID(v.ID)
		if err != nil {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusInternalServerError).
				AddError(err.Error()))
			return
		}

		for _, gp := range gps {

			var total float32

			cas, err := assignment.GetCompleteByUserIDGradeParameterID(u.ID, gp.ID)
			if err != nil {
				template.RenderJSONResponse(w, new(template.Response).
					SetCode(http.StatusInternalServerError).
					AddError(err.Error()))
				return
			}

			len := len(cas)
			if len < 1 {
				pSummary.Parameter = append(pSummary.Parameter, profileParameter{
					Name:       gp.Type,
					Percentage: 0,
				})
				continue
			}

			for _, a := range cas {
				total += a.Score
			}

			percentage := (total / float32(len)) * (float32(gp.Percentage) / 100)
			pSummary.Parameter = append(pSummary.Parameter, profileParameter{
				Name:       gp.Type,
				Percentage: percentage,
			})
		}

		// append summary per courses
		res = append(res, pSummary)
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(res))
	return
}
