package domain

type BucketRequest struct {
	Name   string
	Region string
}

type BucketPlan struct {
	Name      string
	Region    string
	Operation string
	DryRun    bool
}
