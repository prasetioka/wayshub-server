package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	channeldto "wayshub-server/dto/channel"
	dto "wayshub-server/dto/result"
	"wayshub-server/models"
	"wayshub-server/pkg/bcrypt"
	"wayshub-server/repositories"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"

	"context"
	"fmt"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type handlerChannel struct {
	ChannelRepository repositories.ChannelRepository
}

func HandlerChannel(ChannelRepository repositories.ChannelRepository) *handlerChannel {
	return &handlerChannel{ChannelRepository}
}

func (h *handlerChannel) FindChannels(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	channels, err := h.ChannelRepository.FindChannels()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	for i, p := range channels {
		channels[i].Cover = os.Getenv("PATH_FILE") + p.Cover
	}

	for i, p := range channels {
		channels[i].Photo = os.Getenv("PATH_FILE") + p.Photo
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: channels}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerChannel) GetChannel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	channel, err := h.ChannelRepository.GetChannel(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	channel.Cover = os.Getenv("PATH_FILE") + channel.Cover
	channel.Photo = os.Getenv("PATH_FILE") + channel.Photo

	for i, p := range channel.Video {
		channel.Video[i].Thumbnail = os.Getenv("PATH_FILE") + p.Thumbnail
	}

	for i, p := range channel.Video {
		channel.Video[i].Video = os.Getenv("PATH_FILE") + p.Video
	}

	for i, p := range channel.Subscription {
		channel.Subscription[i].OtherPhoto = os.Getenv("PATH_FILE") + p.OtherPhoto
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: convertResponse(channel)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerChannel) UpdateChannel(w http.ResponseWriter, r *http.Request) {
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

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	dataCover := r.Context().Value("dataCover")
	fileCover := dataCover.(string)

	dataPhoto := r.Context().Value("dataPhoto")
	filePhoto := dataPhoto.(string)

	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	resp1, err := cld.Upload.Upload(ctx, fileCover, uploader.UploadParams{Folder: "wayshub"})
	if err != nil {
		fmt.Println(err.Error())
	}

	resp2, err := cld.Upload.Upload(ctx, filePhoto, uploader.UploadParams{Folder: "wayshub"})
	if err != nil {
		fmt.Println(err.Error())
	}

	request := channeldto.UpdateChannelRequest{
		ChannelName: r.FormValue("channelName"),
		Description: r.FormValue("description"),
	}

	password, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	channel, err := h.ChannelRepository.GetChannel(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Email != "" {
		channel.Email = request.Email
	}

	if request.Password != "" {
		channel.Password = password
	}

	if request.ChannelName != "" {
		channel.ChannelName = request.ChannelName
	}

	if fileCover != "false" {
		channel.Cover = resp1.SecureURL
	}

	if filePhoto != "false" {
		channel.Photo = resp2.SecureURL
	}

	data, err := h.ChannelRepository.UpdateChannel(channel)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	channel.Cover = os.Getenv("PATH_FILE") + channel.Cover
	channel.Photo = os.Getenv("PATH_FILE") + channel.Photo

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: updateResponse(data)}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerChannel) DeleteChannel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	channelInfo := r.Context().Value("channelInfo").(jwt.MapClaims)
	channelID := int(channelInfo["id"].(float64))

	channel, err := h.ChannelRepository.GetChannel(channelID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.ChannelRepository.DeleteChannel(channel)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: deleteResponse(data)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerChannel) PlusSubscriber(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	channel, err := h.ChannelRepository.GetChannel(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	channel.Subscriber = channel.Subscriber + 1

	data, err := h.ChannelRepository.PlusSubscriber(channel)
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

func (h *handlerChannel) MinusSubscriber(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	channel, err := h.ChannelRepository.GetChannel(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	channel.Subscriber = channel.Subscriber - 1

	data, err := h.ChannelRepository.MinusSubscriber(channel)
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

func convertResponse(u models.Channel) channeldto.ChannelResponse {
	return channeldto.ChannelResponse{
		ID:           u.ID,
		Email:        u.Email,
		ChannelName:  u.ChannelName,
		Description:  u.Description,
		Cover:        u.Cover,
		Photo:        u.Photo,
		Video:        u.Video,
		Subscription: u.Subscription,
		Subscriber:   u.Subscriber,
	}
}

func updateResponse(u models.Channel) channeldto.ChannelResponse {
	return channeldto.ChannelResponse{
		ID:          u.ID,
		Email:       u.Email,
		ChannelName: u.ChannelName,
		Description: u.Description,
		Cover:       u.Cover,
		Photo:       u.Photo,
	}
}

func deleteResponse(u models.Channel) channeldto.DeleteResponse {
	return channeldto.DeleteResponse{
		ID: u.ID,
	}
}
