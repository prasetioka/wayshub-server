package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	dto "wayshub-server/dto/result"
	subscribedto "wayshub-server/dto/subscribe"
	"wayshub-server/models"
	"wayshub-server/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerSubscribe struct {
	SubscribeRepository repositories.SubscribeRepository
}

func HandlerSubscribe(SubscribeRepository repositories.SubscribeRepository) *handlerSubscribe {
	return &handlerSubscribe{SubscribeRepository}
}

// func (h *handlerSubscribe) FindSubscribes(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	subscribes, err := h.SubscribeRepository.FindSubscribes()
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Status: "success", Data: subscribes}
// 	json.NewEncoder(w).Encode(response)
// }

func (h *handlerSubscribe) FindSubscribes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	channelInfo := r.Context().Value("channelInfo").(jwt.MapClaims)
	channelID := int(channelInfo["id"].(float64))

	subscribes, err := h.SubscribeRepository.FindSubscribes(channelID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	for i, p := range subscribes {
		subscribes[i].OtherPhoto = os.Getenv("PATH_FILE") + p.OtherPhoto
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: subscribes}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerSubscribe) GetSubscribe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var subscribe models.Subscribe
	subscribe, err := h.SubscribeRepository.GetSubscribe(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: subscribe}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerSubscribe) GetSubscribeByOther(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var subscribe models.Subscribe
	subscribe, err := h.SubscribeRepository.GetSubscribeByOther(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: subscribe}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerSubscribe) CreateSubscribe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	channelInfo := r.Context().Value("channelInfo").(jwt.MapClaims)
	channelID := int(channelInfo["id"].(float64))

	request := subscribedto.SubscribeRequest{
		ChannelID: channelID,
		OtherID:   id,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	other, err := h.SubscribeRepository.GetOther(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	subscribe := models.Subscribe{
		ChannelID:  request.ChannelID,
		OtherID:    request.OtherID,
		OtherName:  other.ChannelName,
		OtherPhoto: other.Photo,
		OtherVideo: other.Video,
	}

	subscribe, err = h.SubscribeRepository.CreateSubscribe(subscribe)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// subscribe, _ = h.SubscribeRepository.GetSubscribe(subscribe.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: subscribe}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerSubscribe) DeleteSubscribe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	channelInfo := r.Context().Value("channelInfo").(jwt.MapClaims)
	channelID := int(channelInfo["id"].(float64))

	subscribe, err := h.SubscribeRepository.GetSubscribe(channelID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.SubscribeRepository.DeleteSubscribe(subscribe)
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

func (h *handlerSubscribe) GetSubscription(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// id, _ := strconv.Atoi(mux.Vars(r)["id"])

	channelInfo := r.Context().Value("channelInfo").(jwt.MapClaims)
	channelID := int(channelInfo["id"].(float64))

	// var subscribe models.Channel
	subscribe, err := h.SubscribeRepository.GetSubscription(channelID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: subscribe}
	json.NewEncoder(w).Encode(response)
}
