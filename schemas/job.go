package schemas

type JobLocation string

const (
	JobLocationRemote JobLocation = "REMOTE"
	JobLocationOnsite JobLocation = "ONSITE"
	JobLocationHybrid JobLocation = "HYBRID"
)

type JobType string

const (
	JobTypeFullTime   JobType = "FULLTIME"
	JobTypePartTime   JobType = "PARTTIME"
	JobTypeContract   JobType = "CONTRACT"
	JobTypeFreelance  JobType = "FREELANCE"
	JobTypeInternship JobType = "INTERNSHIP"
)

type JobCategory string

const (
	JobCategoryProgramming     JobCategory = "PROGRAMMING"
	JobCategoryBlockchain      JobCategory = "BLOCKCHAIN"
	JobCategoryDesign          JobCategory = "DESIGN"
	JobCategoryMarketing       JobCategory = "MARKETING"
	JobCategoryCustomerSupport JobCategory = "CUSTOMERSUPPORT"
	JobCategoryWriting         JobCategory = "WRITING"
	JobCategoryProduct         JobCategory = "PRODUCT"
	JobCategoryService         JobCategory = "SERVICE"
	JobCategoryHumanResource   JobCategory = "HUMANRESOURCE"
	JobCategoryElse            JobCategory = "ELSE"
)

type JobExperienceLevel string

const (
	JobExperienceLevelAll                  JobExperienceLevel = "ALL"
	JobExperienceLevelJunior               JobExperienceLevel = "JUNIOR"
	JobExperienceLevelMid                  JobExperienceLevel = "MID"
	JobExperienceLevelSenior               JobExperienceLevel = "SENIOR"
	JobExperienceLevelNoExperienceRequired JobExperienceLevel = "NOEXPERIENCEREQUIRED"
)

type JobStatus string

const (
	JobStatusDraft     JobStatus = "DRAFT"
	JobStatusPaid      JobStatus = "PAID"
	JobStatusConfirmed JobStatus = "CONFIRMED"
	JobStatusHold      JobStatus = "HOLD"
	JobStatusReview    JobStatus = "REVIEW"
	JobStatusClosed    JobStatus = "CLOSED"
)

type UserJobCreatedByUserResponse struct {
	Slug string `json:"slug"`
}