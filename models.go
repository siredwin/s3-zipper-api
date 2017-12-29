package main



type (
	// UserApiJwt : Custom struct for JWT userkey and secret.
	UserApiJwt struct {
		UserKey    string  `json:"userKey"      form:"userKey"      query:"userKey"`
		UserSecret string  `json:"userSecret"   form:"userSecret"   query:"userSecret"`
	}

	// GetToken : Struct for token
	GetToken struct {
		Token          string `json:"token"`
	}

	// TaskIDs : This is the task UUID Model
	TaskIDs struct {
		TaskUUID      []map[string]string  `json:"taskUUID,omitempty"`
		ChainTaskUUID []map[string]string  `json:"chainTaskUUID,omitempty"`
	}

)

const (
	HOST_IP = ":8000"
)

