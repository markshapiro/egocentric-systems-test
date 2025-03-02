package calculator

import (
	"encoding/json"
	"io"
	"math/big"
	"net/http"
)

type OperationHandler struct {
	service OperationService
}

func NewOperationHandler(service OperationService) OperationHandler {
	return OperationHandler{service}
}

func (t OperationHandler) MountEndpoints(mux *http.ServeMux) {
	mux.HandleFunc("/add", t.addHandler)
	mux.HandleFunc("/subtract", t.subHandler)
	mux.HandleFunc("/multiply", t.subMultiply)
	mux.HandleFunc("/divide", t.subDivide)
	mux.HandleFunc("/getRecentN", t.getRecentN)
}

func (t OperationHandler) addHandler(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodPost:

		defer r.Body.Close()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}

		var operands OperandsDto

		err = json.Unmarshal(body, &operands)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		result, err := t.service.Add(operands.OperandA, operands.OperandB)
		if err != nil {
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}

		writeResponse(w, &result, http.StatusOK)

	default:
		notFoundHandler(w)
	}
}

func (t OperationHandler) subHandler(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodPost:

		defer r.Body.Close()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}

		var operands OperandsDto

		err = json.Unmarshal(body, &operands)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		result, err := t.service.Subtract(operands.OperandA, operands.OperandB)
		if err != nil {
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}

		writeResponse(w, &result, http.StatusOK)

	default:
		notFoundHandler(w)
	}
}

func (t OperationHandler) subMultiply(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodPost:

		defer r.Body.Close()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}

		var operands OperandsDto

		err = json.Unmarshal(body, &operands)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		result, err := t.service.Multiply(operands.OperandA, operands.OperandB)
		if err != nil {
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}

		writeResponse(w, &result, http.StatusOK)

	default:
		notFoundHandler(w)
	}
}

func (t OperationHandler) subDivide(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodPost:

		defer r.Body.Close()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}

		var operands OperandsDto

		err = json.Unmarshal(body, &operands)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		if operands.OperandB.Cmp(big.NewFloat(0.0)) == 0 {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		result, err := t.service.Divide(operands.OperandA, operands.OperandB)
		if err != nil {
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}

		writeResponse(w, &result, http.StatusOK)

	default:
		notFoundHandler(w)
	}
}

func (t OperationHandler) getRecentN(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet:

		result, err := t.service.GetRecentN(5)
		if err != nil {
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}

		var resultDto []ResultDto

		for _, r := range result {
			resultDto = append(resultDto, ResultDto{r.OperandA, r.OperandB, r.Operator, r.Result})
		}

		writeResponse(w, resultDto, http.StatusOK)

	default:
		notFoundHandler(w)
	}
}

func writeResponse(w http.ResponseWriter, resp any, status int) {
	b, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	w.WriteHeader(status)
	w.Write(b)
}

func notFoundHandler(w http.ResponseWriter) {
	http.Error(w, "Not Found", http.StatusNotFound)
}
