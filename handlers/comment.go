package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	commentdto "wayshub-server/dto/comment"
	dto "wayshub-server/dto/result"
	"wayshub-server/models"
	"wayshub-server/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerComment struct {
	CommentRepository repositories.CommentRepository
}

func HandlerComment(CommentRepository repositories.CommentRepository) *handlerComment {
	return &handlerComment{CommentRepository}
}

func (h *handlerComment) FindComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	comments, err := h.CommentRepository.FindComments()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: comments}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerComment) GetComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	comment, err := h.CommentRepository.GetComment(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: comment}
	json.NewEncoder(w).Encode(response)
}

// func (h *handlerComment) CreateComment(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	channelInfo := r.Context().Value("channelInfo").(jwt.MapClaims)
// 	channelID := int(channelInfo["id"].(float64))

// 	request := commentdto.CreateCommentRequest{
// 		Comment: r.FormValue("comment"),
// 	}

// 	validation := validator.New()
// 	err := validation.Struct(request)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	comment := models.Comments{
// 		Comment:   request.Comment,
// 		ChannelID: channelID,
// 	}

// 	comment, err = h.CommentRepository.CreateComment(comment)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	comment, _ = h.CommentRepository.GetComment(comment.ID)

// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Status: "success", Data: comment}
// 	json.NewEncoder(w).Encode(response)
// }

func (h *handlerComment) CreateComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	channelInfo := r.Context().Value("channelInfo").(jwt.MapClaims)
	channelID := int(channelInfo["id"].(float64))

	request := commentdto.CreateCommentRequest{
		Comment: r.FormValue("comment"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	comment := models.Comments{
		Comment:   request.Comment,
		ChannelID: channelID,
		VideoID: id,
	}

	comment, err = h.CommentRepository.CreateComment(comment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	comment, _ = h.CommentRepository.GetComment(comment.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: comment}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerComment) UpdateComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	channelInfo := r.Context().Value("channelInfo").(jwt.MapClaims)
	channelID := int(channelInfo["id"].(float64))

	if channelID != id {
		w.WriteHeader(http.StatusUnauthorized)
		response := dto.ErrorResult{Code: http.StatusUnauthorized, Message: "Can't update channel!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	request := commentdto.CreateCommentRequest{
		Comment: r.FormValue("comment"),
	}

	comment, err := h.CommentRepository.GetComment(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Comment != "" {
		comment.Comment = request.Comment
	}

	data, err := h.CommentRepository.UpdateComment(comment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: data}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerComment) DeleteComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	channelInfo := r.Context().Value("channelInfo").(jwt.MapClaims)
	channelID := int(channelInfo["id"].(float64))

	comment, err := h.CommentRepository.GetComment(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if channelID != comment.ChannelID {
		w.WriteHeader(http.StatusUnauthorized)
		response := dto.ErrorResult{Code: http.StatusUnauthorized, Message: "PLease Login First!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.CommentRepository.DeleteComment(comment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: DeleteCommentResponse(data)}
	json.NewEncoder(w).Encode(response)
}

func DeleteCommentResponse(u models.Comments) commentdto.DeleteResponse {
	return commentdto.DeleteResponse{
		ID: u.ID,
	}
}