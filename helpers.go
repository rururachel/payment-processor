package paymentprocessor

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

// generateUUID returns a random UUID
func generateUUID() string {
	return uuid.New().String()
}

// hashString returns the SHA-256 hash of a given string
func hashString(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// parseJSON parses a JSON string into a map
func parseJSON(s string) (map[string]interface{}, error) {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(s), &data); err != nil {
		return nil, err
	}
	return data, nil
}

// isValidUUID checks if a given string is a valid UUID
func isValidUUID(uuid string) bool {
	_, err := uuid.Parse(uuid)
	return err == nil
}

// getHTTPHeader returns the value of a given HTTP header
func getHTTPHeader(r *http.Request, key string) string {
	return r.Header.Get(key)
}

// getIntHTTPHeader returns the integer value of a given HTTP header
func getIntHTTPHeader(r *http.Request, key string) (int, error) {
	value := getHTTPHeader(r, key)
	if value == "" {
		return 0, errors.New("header not found")
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// getFloatHTTPHeader returns the float value of a given HTTP header
func getFloatHTTPHeader(r *http.Request, key string) (float64, error) {
	value := getHTTPHeader(r, key)
	if value == "" {
		return 0, errors.New("header not found")
	}
	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, err
	}
	return f, nil
}

// generateRandomString returns a random string of a given length
func generateRandomString(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(b))
}

// currentTime returns the current time in UTC
func currentTime() time.Time {
	return time.Now().UTC()
}