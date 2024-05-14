package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/darshan/services/database"
	"github.com/spf13/viper"
)

type ServiceController struct {
	DB *sql.DB
}

func NewServiceController(DB *sql.DB) ServiceController {
	return ServiceController{DB}
}

type Service struct {
	ServiceID    int    `json:"service_id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	VersionCount int    `json:"version_count"`
}

func (s *ServiceController) Services(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	}()

	inputParams := make(map[string]interface{})

	err := parseServiceRequest(r, inputParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	params := make([]interface{}, 0)

	query := serviceQueryBuilder(&params, inputParams, r)

	// Execute the query with an offset parameter

	rows, err := database.ExecuteQuery(s.DB, query, params)
	if err != nil {
		http.Error(w, "Error fetching services"+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var services []Service
	for rows.Next() {
		var service Service
		if err := rows.Scan(&service.ServiceID, &service.Name, &service.Description, &service.VersionCount); err != nil {
			http.Error(w, "Error scanning rows", http.StatusInternalServerError)
			return
		}
		services = append(services, service)
	}

	// Encode the response as JSON and write it to the response writer
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services)
}

func parseServiceRequest(r *http.Request, inputParams map[string]interface{}) error {

	var err error
	var name string
	var offset int
	var limit int

	pageStr := r.URL.Query().Get("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return fmt.Errorf("Invalid Page Number")
	}
	if page < 1 {
		page = 1
	}

	limitStr := r.URL.Query().Get("limit")
	if limitStr == "" {
		limit = 2
	} else {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			return fmt.Errorf("Invalid Limit")
		}
		if limit < 1 {
			limit = 1
		}
	}

	offset = (page - 1) * limit // 2 records per page
	name = r.URL.Query().Get("name")
	sortBy := r.URL.Query().Get("sort_by")
	if sortBy != "name" {
		sortBy = "created_on"
	}
	sortType := r.URL.Query().Get("sort_type")
	if sortType == "D" {
		sortType = "DESC"
	} else {
		sortType = ""
	}

	inputParams["name"] = name
	inputParams["offset"] = offset
	inputParams["limit"] = limit
	inputParams["sortBy"] = sortBy
	inputParams["sortType"] = sortType

	return nil
}

func serviceQueryBuilder(params *[]interface{}, inputParams map[string]interface{}, r *http.Request) string {

	var query string
	name := inputParams["name"].(string)
	offset := inputParams["offset"].(int)
	limit := inputParams["limit"].(int)
	sortBy := inputParams["sortBy"].(string)
	sortType := inputParams["sortType"].(string)

	if name != "" {
		*params = append(*params, name)
		query = viper.GetString("queries.serviceQuery")
	} else {
		query = viper.GetString("queries.listingQuery")
		condition := sortBy + " " + sortType
		query = strings.Replace(query, "{condition}", condition, -1)
		*params = append(*params, limit, offset)
	}
	return query
}
