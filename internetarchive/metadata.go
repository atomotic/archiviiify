package internetarchive

import (
	"context"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	iiif "github.com/atomotic/iiif-presentation.go/v3"
	"github.com/carlmjohnson/requests"

	getter "github.com/hashicorp/go-getter/v2"
)

const base = "https://archive.org/metadata/"

type Item struct {
	Created         int      `json:"created"`
	D1              string   `json:"d1"`
	D2              string   `json:"d2"`
	Dir             string   `json:"dir"`
	Files           []File   `json:"files"`
	FilesCount      int      `json:"files_count"`
	ItemLastUpdated int      `json:"item_last_updated"`
	ItemSize        int      `json:"item_size"`
	Metadata        Metadata `json:"metadata"`
	Server          string   `json:"server"`
	Uniq            int      `json:"uniq"`
	WorkableServers []string `json:"workable_servers"`
	JP2Zip          string
}

type Metadata struct {
	Identifier       string        `json:"identifier"`
	Mediatype        string        `json:"mediatype"`
	Collection       []string      `json:"collection"`
	Creator          string        `json:"creator"`
	Date             string        `json:"date"`
	Description      stringOrArray `json:"description"`
	Language         string        `json:"language"`
	Licenseurl       string        `json:"licenseurl"`
	Scanner          string        `json:"scanner"`
	Subject          stringOrArray `json:"subject"`
	Title            string        `json:"title"`
	Publicdate       string        `json:"publicdate"`
	Uploader         string        `json:"uploader"`
	Addeddate        string        `json:"addeddate"`
	IdentifierAccess string        `json:"identifier-access"`
	IdentifierArk    string        `json:"identifier-ark"`
	Imagecount       string        `json:"imagecount"`
	Ppi              string        `json:"ppi"`
	Ocr              string        `json:"ocr"`
	RepubState       string        `json:"repub_state"`
	BackupLocation   string        `json:"backup_location"`
}

type File struct {
	Name      string `json:"name"`
	Source    string `json:"source"`
	Mtime     string `json:"mtime,omitempty"`
	Size      string `json:"size,omitempty"`
	Md5       string `json:"md5"`
	Crc32     string `json:"crc32,omitempty"`
	Sha1      string `json:"sha1,omitempty"`
	Format    string `json:"format"`
	Rotation  string `json:"rotation,omitempty"`
	Original  string `json:"original,omitempty"`
	Btih      string `json:"btih,omitempty"`
	Summation string `json:"summation,omitempty"`
}

func (i *Item) Downloaded() bool {
	imagedir := filepath.Join("./data/images/", i.Metadata.Identifier)
	if _, err := os.Stat(imagedir); os.IsNotExist(err) {
		return false
	}
	return true
}

func (i *Item) Download() error {
	ctx := context.Background()

	req := &getter.Request{
		Src:              i.JP2Zip,
		Dst:              "data/images/",
		Pwd:              "./",
		GetMode:          getter.ModeAny,
		ProgressListener: defaultProgressBar,
	}

	client := &getter.Client{
		Getters: []getter.Getter{
			&getter.HttpGetter{
				ReadTimeout: 60 * time.Minute,
			},
		},
	}

	_, err := client.Get(ctx, req)
	if err != nil {
		return err
	}

	filename := path.Base(i.JP2Zip)
	filename = strings.TrimSuffix(filename, path.Ext(filename))
	err = os.Rename("./data/images/"+filename, "./data/images/"+i.Metadata.Identifier)
	if err != nil {
		return err
	}
	return nil
}

func (i *Item) Manifest() error {
	if len(i.Metadata.Identifier) == 0 {
		return fmt.Errorf("item error")
	}
	imagedir := filepath.Join("./data/images/", i.Metadata.Identifier)

	if _, err := os.Stat(imagedir); os.IsNotExist(err) {
		return fmt.Errorf("error finding image directory: %w", err)
	}

	manifest, _ := iiif.NewManifest(i.Metadata.Identifier, os.Getenv("BASE"))
	manifest.Label = iiif.Label{"it": {i.Metadata.Title}}

	values := reflect.ValueOf(i.Metadata)
	types := values.Type()
	for i := 0; i < values.NumField(); i++ {
		var newvalue string
		field := types.Field(i)
		value := values.Field(i)
		if field.Type.Kind() == reflect.Slice && field.Type.Elem().Kind() == reflect.String {
			strSlice := value.Interface().([]string)
			newvalue = strings.Join(strSlice, ", ")
		} else {
			newvalue = value.String()
		}

		meta := iiif.NewMetadata(field.Name, newvalue)
		manifest.Metadata = append(manifest.Metadata, meta)
	}

	jp2, err := filepath.Glob(filepath.Join(imagedir, string(os.PathSeparator), "*"))
	if err != nil {
		return fmt.Errorf("error globbing jp2 files: %w", err)
	}

	counter := 1
	for _, image := range jp2 {
		api := fmt.Sprintf("%s/iiif/%s/%s/info.json",
			os.Getenv("HOSTNAME"),
			i.Metadata.Identifier,
			path.Base(image))

		err := manifest.NewItem(
			fmt.Sprint(counter),
			fmt.Sprintf("Page %d", counter),
			strings.Replace(api, "/info.json", "", -1),
			nil,
			api)

		if err != nil {
			return fmt.Errorf("error creating manifest item: %w", err)
		}
		counter = counter + 1
	}

	m := manifest.Serialize()

	err = os.WriteFile("./data/manifests/"+i.Metadata.Identifier+".json", []byte(m), 0644)
	if err != nil {
		return fmt.Errorf("error writing manifest: %w", err)
	}

	return nil
}

func New(identifier string) (*Item, error) {
	ctx := context.Background()
	itemurl := fmt.Sprintf("%s%s", base, identifier)
	item := Item{}
	err := requests.URL(itemurl).ToJSON(&item).Fetch(ctx)
	if err != nil {
		return nil, err
	}
	if len(item.Metadata.Identifier) == 0 {
		return nil, fmt.Errorf("item error")
	}

	for _, file := range item.Files {
		if file.Format == "Single Page Processed JP2 ZIP" {
			item.JP2Zip = fmt.Sprintf("https://%s%s/%s", item.D1, item.Dir, file.Name)
		}
	}

	return &item, nil
}
