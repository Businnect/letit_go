package resources

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/Businnect/letit_go/interfaces"
	"github.com/Businnect/letit_go/schemas"
)

type JobResource struct {
	client interfaces.ClientInterface
}

type CreateUserJobRequest struct {
	CompanyName        string
	CompanyDescription string
	CompanyWebsite     string
	JobTitle           string
	JobDescription     string
	JobHowToApply      string

	CompanyLogo            *FilePayload
	CompanyLocation        *string
	JobLocation            schemas.JobLocation
	JobType                schemas.JobType
	JobCategory            schemas.JobCategory
	JobExperienceLevel     schemas.JobExperienceLevel
	JobMinimumSalary       *int
	JobMaximumSalary       *int
	JobPayInCryptocurrency bool
	JobSkills              *string
}

func NewJobResource(c interfaces.ClientInterface) *JobResource {
	return &JobResource{client: c}
}

func (r *JobResource) CreateWithCompany(ctx context.Context, req CreateUserJobRequest) (*schemas.UserJobCreatedByUserResponse, error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	_ = w.WriteField("company_name", req.CompanyName)
	_ = w.WriteField("company_description", req.CompanyDescription)
	_ = w.WriteField("company_website", req.CompanyWebsite)
	_ = w.WriteField("job_title", req.JobTitle)
	_ = w.WriteField("job_description", req.JobDescription)
	_ = w.WriteField("job_how_to_apply", req.JobHowToApply)
	_ = w.WriteField("job_pay_in_cryptocurrency", strconv.FormatBool(req.JobPayInCryptocurrency))

	if req.JobLocation == "" {
		req.JobLocation = schemas.JobLocationRemote
	}
	_ = w.WriteField("job_location", string(req.JobLocation))

	if req.JobType == "" {
		req.JobType = schemas.JobTypeFullTime
	}
	_ = w.WriteField("job_type", string(req.JobType))

	if req.JobCategory == "" {
		req.JobCategory = schemas.JobCategoryProgramming
	}
	_ = w.WriteField("job_category", string(req.JobCategory))

	if req.JobExperienceLevel == "" {
		req.JobExperienceLevel = schemas.JobExperienceLevelAll
	}
	_ = w.WriteField("job_experience_level", string(req.JobExperienceLevel))

	if req.CompanyLocation != nil {
		_ = w.WriteField("company_location", *req.CompanyLocation)
	}
	if req.JobMinimumSalary != nil {
		_ = w.WriteField("job_minimum_salary", strconv.Itoa(*req.JobMinimumSalary))
	}
	if req.JobMaximumSalary != nil {
		_ = w.WriteField("job_maximum_salary", strconv.Itoa(*req.JobMaximumSalary))
	}
	if req.JobSkills != nil {
		_ = w.WriteField("job_skills", *req.JobSkills)
	}

	if req.CompanyLogo != nil {
		part, err := w.CreateFormFile("company_logo", req.CompanyLogo.Filename)
		if err != nil {
			return nil, err
		}
		if _, err = io.Copy(part, req.CompanyLogo.Reader); err != nil {
			return nil, err
		}
	}

	if err := w.Close(); err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", "/api/v1/client/job", &b)
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", w.FormDataContentType())

	respBody, err := r.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer respBody.Close()

	var result schemas.UserJobCreatedByUserResponse
	if err := json.NewDecoder(respBody).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *JobResource) Delete(ctx context.Context, slug string) error {
	data := map[string]string{
		"slug": slug,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodDelete, "/api/v1/client/job", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	httpRequest.Header.Set("Content-Type", "application/json")

	respBody, err := r.client.Do(httpRequest)
	if err != nil {
		return err
	}
	defer respBody.Close()

	return nil
}
