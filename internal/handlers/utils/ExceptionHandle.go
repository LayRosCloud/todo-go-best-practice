package handlersUtils

import (
	"leafall/todo-service/utils/exceptions"
	"net/http"
)

const TotalCountHeader = "X-Total-Count"

func HandleError(err error, w http.ResponseWriter) {
	if err == nil {
		return
	}

	switch e := err.(type) {
		case *exceptions.NotFoundError:
			exceptions.WriteError(w, exceptions.GetNotFoundError(e.Error(), e.Details))
			break;
		case *exceptions.BadRequestError:
			exceptions.WriteError(w, exceptions.GetBadRequestError(e.Error(), e.Details))
			break;
		default:
			exceptions.WriteError(w, exceptions.ErrorResponse{
				Status: 500, 
				Message: "Internal server error",
				Details: e.Error(),
			})
			break
	}
}
