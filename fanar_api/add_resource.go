package fanar

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/Modalessi/iau_resources/models"
)

type addResourceRequest struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

func addResource(fs *FanarServer, w http.ResponseWriter, r *http.Request) (*FanarResponse, error) {

	courseID := r.URL.Query().Get("id")
	if courseID == "" {
		return fanarMessageResponse(400, "invalid request, please pass the course id as query"), nil
	}

	resourceDetails := addResourceRequest{}
	err := json.Unmarshal([]byte(r.FormValue("data")), &resourceDetails)
	if err != nil {
		return fanarMessageResponse(400, "please send valid json"), nil
	}

	invalidTags := models.InvalidResourceTags(resourceDetails.Tags...)
	if len(invalidTags) > 0 {
		invalidTagsMSG := "invalid tags: " + strings.Join(invalidTags, ", ")
		return fanarMessageResponse(400, invalidTagsMSG), nil
	}

	if resourceDetails.Title == "" {
		return fanarMessageResponse(404, "title is missing, please check"), nil
	}

	course, err := fs.Storage.GetCourseByID(r.Context(), courseID)
	if err != nil {
		return fanarMessageResponse(404, "there is no course with such id"), err
	}

	file, metaData, err := r.FormFile("file")
	if err != nil {
		return fanarMessageResponse(401, "error getting file from request, please try again"), err
	}

	fileExt := filepath.Ext(metaData.Filename)

	userEmail := r.Context().Value(USER_EMAIL_KEY).(string)
	user, err := fs.Storage.GetUserByEmail(r.Context(), userEmail)

	resource := models.NewResource(*course.ID, resourceDetails.Title, resourceDetails.Description, fileExt, resourceDetails.Tags, *user.ID)

	contentType := metaData.Header.Get("Content-Type")
	fs.Storage.StoreResource(r.Context(), resource, file, contentType)

	return fanarMessageResponse(201, "file upladed succesffuly"), nil
}
