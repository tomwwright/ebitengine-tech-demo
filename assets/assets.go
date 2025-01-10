package assets

import "io/fs"

type Asset string

const (
	AssetAudioText  Asset = "assets/audio/text.wav"
	AssetAudioMusic Asset = "assets/audio/music.ogg"
)

type FileSystemAssets struct {
	files fs.ReadFileFS
	cache map[Asset][]byte
}

func NewFileSystemAssets(files fs.ReadFileFS) *FileSystemAssets {
	return &FileSystemAssets{
		files: files,
		cache: make(map[Asset][]byte),
	}
}

func (a *FileSystemAssets) GetAsset(asset Asset) ([]byte, error) {
	if hit := a.cache[asset]; hit != nil {
		return hit, nil
	}

	b, err := a.files.ReadFile(string(asset))
	if err != nil {
		return nil, err
	}

	a.cache[asset] = b
	return b, nil
}
