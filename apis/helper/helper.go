package helper

import (
	"errors"
	"fmt"
	"strings"

	"github.com/aniket-skroman/skroman_sales_service.git/utils"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var (
	ERR_INVALID_ID      error
	ERR_REQUIRED_PARAMS error
	Err_Lead_Exists     error
	Err_Data_Not_Found  error
	Err_Update_Failed   error
	Err_Delete_Failed   error
)

func init() {
	ERR_INVALID_ID = errors.New("invalid id found")
	ERR_REQUIRED_PARAMS = errors.New("please provide a required params")
	Err_Lead_Exists = errors.New("lead info already exists for current lead")
	Err_Data_Not_Found = errors.New("data not found")
	Err_Update_Failed = errors.New("failed to update resources")
	Err_Delete_Failed = errors.New("failed to delete resource")
}

func SetPaginationData(page int, total int64) {
	if page == 0 {
		utils.PREVIOUS_IDX = 0
	} else {
		utils.PREVIOUS_IDX = page - 1
	}

	utils.CURRENT_IDX = page
	utils.TOTALCOUNT = total
}

func ValidateUUID(input_id string) (uuid.UUID, error) {
	return uuid.Parse(input_id)
}

type ApiError struct {
	Field string
	Msg   string
}

func Error_handler(err error) []ApiError {
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ApiError, len(ve))
			for i, fe := range ve {
				out[i] = ApiError{fe.Field(), msgForTag(fe.Tag())}
			}
			return out
		}
		return nil
	}
	return nil
}

func Handle_required_param_error(err error) string {
	var ve validator.ValidationErrors
	var err_msg string
	if errors.As(err, &ve) {
		for _, fe := range ve {
			err_msg = fmt.Sprintf("%v - %v", fe.Field(), msgForTag(fe.Tag()))
			break
		}
	} else {
		if strings.Contains(err.Error(), "cannot unmarshal string into") {
			err_msg = "required a integer but found string, please check params"
		} else if strings.Contains(err.Error(), "cannot unmarshal number into") {
			err_msg = "required a string but found integer, please check params"
		} else {
			err_msg = "something went's wrong, invalid param detecte"

		}
	}

	return err_msg
}

func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "min":
		return "Invalid length for param"
	}
	return ""
}
