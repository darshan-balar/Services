package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/darshan/services/database"
	"github.com/spf13/viper"
)


type VersionController struct {
	DB *sql.DB
}

func GetVersionController(DB *sql.DB) VersionController {
	return VersionController{DB}
}


type Versions struct {
    Name           string `json:"name"`
    Description    string `json:"description"`
}

func (v *VersionController) GetVersions(w http.ResponseWriter, r *http.Request){

    defer func(){
        if r := recover(); r != nil {
            http.Error(w, "Internal server error", http.StatusInternalServerError)
        }
    }()

    inputParams := make(map[string]interface{})

    // Parse the request and get the input parameters
    err := parseRequest(r,inputParams)
    if err != nil {
        http.Error(w, err.Error() ,  http.StatusBadRequest)
        return
    }
    params := make([]interface{}, 0)

    // Build the query based on the input parameters
    query := queryBuilder(&params, inputParams, r)

    rows, err := database.ExecuteQuery(v.DB, query, params)
    if err != nil {
        http.Error(w, "Error fetching versions"+err.Error(), http.StatusInternalServerError)
        return  
    }
    defer rows.Close()
    var versions []Versions
    // Iterate over the rows and scan the values into a struct
    for rows.Next() {
        var version Versions
        if err := rows.Scan(&version.Name, &version.Description); err != nil {
            http.Error(w, "Error scanning rows", http.StatusInternalServerError)
            return
        }
        versions = append(versions, version)
    }
    if len(versions) < 1{
        http.Error(w, "No versions found", http.StatusNotFound)
        return
    }
    // Encode the response as JSON and write it to the response writer
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(versions)
}

func parseRequest(r *http.Request, inputParams map[string]interface{}) error{

    var serviceId int

    serviceIdStr := r.URL.Query().Get("service_id")

    if serviceIdStr == "" {
        return errors.New("Service id is required")
    }

    serviceId, err := strconv.Atoi(serviceIdStr)
    if err != nil {
        return errors.New("Invalid service id")
    }
    inputParams["service_id"] = serviceId
    return nil
}

func queryBuilder(params *[]interface{}, inputParams map[string]interface{}, r *http.Request) string {
    
    serviceID := inputParams["service_id"].(int)
    query := viper.GetString("queries.versionQuery")
    *params = append(*params, serviceID)
    return query
}