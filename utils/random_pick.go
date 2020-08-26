package utils

import (
    "fmt"
    "math/rand"
    "time"
    "github.com/luis-novoa/models"
)

func randomPick(technicians []models.Technician) models.Technician {
    rand.Seed(time.Now().UTC().UnixNano())
    randomIndex := rand.Intn(len(technicians))
    return technicians[randomIndex]
    
}