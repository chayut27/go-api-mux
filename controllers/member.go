package controllers

import (
	"encoding/json"
	"fmt"
	"go-api/models"
	"go-api/repository"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllMember(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("content-type", "application/json")
	members, err := repository.ListAllMember()
	if err != nil {
		json.NewEncoder(w).Encode(models.Get_data_error())
	} else {
		if members == nil {
			members = make([]models.Member, 0)
		}
		json.NewEncoder(w).Encode(members)
	}
}

func GetOneMember(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	param := mux.Vars(r)
	objID, _ := primitive.ObjectIDFromHex(param["id"])
	members, err := repository.FindOneMember(objID)
	if err != nil {
		json.NewEncoder(w).Encode(models.Get_data_error())
	} else {
		json.NewEncoder(w).Encode(members)
	}
}

func PostMember(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("content-type", "application/json")
	var member models.Member

	err := json.NewDecoder(req.Body).Decode(&member)
	if err != nil {
		// fmt.Println("err", err)
		json.NewEncoder(w).Encode(models.InvalidSyntax())
	} else {
		// fmt.Println("fname", member.Fname)
		member.CreateAt = time.Now()
		err := repository.InsertMember(member)
		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode(models.InsertError())
		} else {
			json.NewEncoder(w).Encode(models.InsertSuccess())
		}
	}
}

func PutMember(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	param := mux.Vars(r)
	objID, _ := primitive.ObjectIDFromHex(param["id"])

	_, err := repository.FindOneMember(objID)
	if err != nil {
		json.NewEncoder(w).Encode(models.Get_data_error())
	} else {

		var member models.Member
		err = json.NewDecoder(r.Body).Decode(&member)
		if err != nil {
			json.NewEncoder(w).Encode(models.InvalidSyntax())
		} else {
			err = repository.UpdateMember(objID, member)
			if err != nil {
				json.NewEncoder(w).Encode(models.UpdateError())
			} else {
				json.NewEncoder(w).Encode(models.UpdateSuccess())
			}
		}
	}
}

func DelMember(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	param := mux.Vars(r)
	objID, _ := primitive.ObjectIDFromHex(param["id"])

	_, err := repository.FindOneMember(objID)
	if err != nil {
		json.NewEncoder(w).Encode(models.Get_data_error())
	} else {
		err = repository.DeleteMember(objID)
		if err != nil {
			json.NewEncoder(w).Encode(models.DeleteError())
		} else {
			json.NewEncoder(w).Encode(models.DeleteSuccess())
		}
	}
}
