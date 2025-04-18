// used this for getting and saving the data from and to the json file
package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
	"github.com/hababisha/ejigo/models"
)


var path = filepath.Join("data", "courses.json")

func GetCourses() (map[string]*models.Course, error){
	courses := make(map[string]*models.Course)

	if _, err := os.Stat(path); os.IsNotExist(err){
		return courses, nil
	}
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bytes, &courses); err != nil {
		return nil, err
	}
	return courses, nil
}

func SaveCourses(courses map[string]*models.Course) error {
	bytes, err := json.MarshalIndent(courses, "", " ")
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	return os.WriteFile(path, bytes, 0644)
}
