package tiled

import (
	"bytes"
	"fmt"
	"image"
	"io/fs"
	"path/filepath"
	"techdemo/interactions/yaml"

	"github.com/lafriks/go-tiled"
)

type loader struct {
	Files    fs.ReadFileFS
	Filename string
}

func newLoader(options ...LoaderOption) *loader {
	l := &loader{}
	for _, o := range options {
		o(l)
	}
	return l
}

type LoaderOption func(*loader)

func WithFiles(files fs.ReadFileFS) LoaderOption {
	return func(l *loader) {
		l.Files = files
	}
}

func withFilename(filename string) LoaderOption {
	return func(l *loader) {
		l.Filename = filename
	}
}

func Load(filename string, options ...LoaderOption) (*Tiled, error) {
	options = append(options, withFilename(filename))
	l := newLoader(options...)

	data, err := l.load()
	if err != nil {
		return nil, fmt.Errorf("error loading Tiled map: %s %w", filename, err)
	}
	return data, err
}

func (l *loader) load() (*Tiled, error) {
	data := &Tiled{}

	err := l.loadTilemap(data)
	if err != nil {
		return nil, err
	}

	err = l.loadTilesets(data)
	if err != nil {
		return nil, err
	}

	err = l.loadInteractions(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (l *loader) loadTilemap(data *Tiled) error {
	filename := l.Filename
	tilemap, err := tiled.LoadFile(filename, tiled.WithFileSystem(l.Files))
	if err != nil {
		return err
	}
	data.tilemap = tilemap
	return nil
}

func (l *loader) loadTilesets(data *Tiled) error {
	data.tilesets = map[string]image.Image{}
	dir := filepath.Dir(l.Filename)
	for _, ts := range data.tilemap.Tilesets {
		filename := dir + "/" + ts.Image.Source
		fmt.Println("loading tileset", ts.Name, "from", filename)

		source, err := l.Files.ReadFile(filename)
		if err != nil {
			return fmt.Errorf("error loading tileset image file: %w", err)
		}

		img, _, err := image.Decode(bytes.NewReader(source))
		if err != nil {
			return fmt.Errorf("error decoding tileset image file: %w", err)
		}
		data.tilesets[ts.Name] = img

	}
	return nil
}

func (l *loader) loadInteractions(data *Tiled) error {
	filename := data.tilemap.Properties.Get("interactionsFilename")[0]
	dir := filepath.Dir(l.Filename)
	filepath := dir + "/" + filename
	content, err := l.Files.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("failed to load interactions from %s: %w", filepath, err)
	}
	interactions, err := yaml.UnmarshallInteractions(content)
	if err != nil {
		return err
	}
	data.interactions = interactions
	return nil
}
