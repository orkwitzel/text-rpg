package savefiles

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"rpg/internal/game"
	"strings"
	"time"
)

const CURRENT_VERSION = 1

type SaveFile struct {
	ID        string    `json:"id"`
	Game      game.Game `json:"game"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Version   int       `json:"version"`
}

func newSaveFile(g game.Game) SaveFile {
	return SaveFile{
		ID:        g.ID,
		Game:      g,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Version:   CURRENT_VERSION,
	}
}

var SaveFilesLocation = filepath.Join(os.Getenv("HOME"), "rpg", "saveFiles")

// SaveGame saves the game, updating an existing save file with the same game ID or creating a new one.
func SaveGame(g game.Game) {
	initSaveFilesDirectory()

	existingSaveFile := loadSaveFile(g.ID)
	if existingSaveFile != nil {
		existingSaveFile.Game = g
		existingSaveFile.UpdatedAt = time.Now()
		writeSaveFile(*existingSaveFile)
		return
	}

	writeSaveFile(newSaveFile(g))
}

func initSaveFilesDirectory() {
	if _, err := os.Stat(SaveFilesLocation); os.IsNotExist(err) {
		err = os.MkdirAll(SaveFilesLocation, 0755)
		if err != nil {
			log.Fatalf("Failed to create save files directory: %v", err)
		}
	}
}

func getSaveFilePath(id string) string {
	return filepath.Join(SaveFilesLocation, id+".json")
}

func LoadSaveFile(gameID string) *SaveFile {
	return loadSaveFile(gameID)
}

func ListSaveFiles() []SaveFile {
	initSaveFilesDirectory()

	entries, err := os.ReadDir(SaveFilesLocation)
	if err != nil {
		return nil
	}

	saveFiles := make([]SaveFile, 0)
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".json" {
			continue
		}

		gameID := strings.TrimSuffix(entry.Name(), ".json")
		saveFile := loadSaveFile(gameID)
		if saveFile != nil {
			saveFiles = append(saveFiles, *saveFile)
		}
	}

	return saveFiles
}

func loadSaveFile(gameID string) *SaveFile {
	data, err := os.ReadFile(getSaveFilePath(gameID))
	if err != nil {
		return nil
	}

	var saveFile SaveFile
	if err := json.Unmarshal(data, &saveFile); err != nil {
		return nil
	}

	return &saveFile
}

func writeSaveFile(saveFile SaveFile) {
	file, err := os.Create(getSaveFilePath(saveFile.ID))
	if err != nil {
		log.Fatalf("Failed to create save file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(saveFile); err != nil {
		log.Fatalf("Failed to encode save file: %v", err)
	}
}
