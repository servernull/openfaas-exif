package function

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/dsoprea/go-exif"
)

type IfdEntry struct {
	IfdPath     string
	FqIfdPath   string
	IfdIndex    int
	TagId       uint16
	TagName     string
	TagTypeId   uint16
	TagTypeName string
	UnitCount   uint32
	Value       interface{}
	ValueString string
}

// Handle a serverless request
func Handle(req []byte) string {

	url := string(req)
	filePath := "temp." + filepath.Ext(url)
	if err := downloadFile(filePath, url); err != nil {
		return "{\"error\": \"error downloading image\"}"
	}

	if data, err := ioutil.ReadFile(filePath); err != nil {
		return "{\"error\": \"error reading image\"}"
	} else {

		exifEntries := []map[string]string{}
		if rawExif, err := exif.SearchAndExtractExif(data); err != nil {
			return "{\"status\": \"No EXIF found in image.\"}"
		} else {

			im := exif.NewIfdMappingWithStandard()
			ti := exif.NewTagIndex()

			entries := make([]IfdEntry, 0)
			visitor := func(fqIfdPath string, ifdIndex int, tagId uint16, tagType exif.TagType, valueContext exif.ValueContext) (err error) {
				defer func() {
					if state := recover(); state != nil {
					}
				}()

				ifdPath, err := im.StripPathPhraseIndices(fqIfdPath)
				if err != nil {
					return err
				}

				it, err := ti.Get(ifdPath, tagId)
				if err != nil {
					return err
				}

				valueString := ""
				var value interface{}
				if tagType.Type() == exif.TypeUndefined {
					var err2 error
					value, err2 = exif.UndefinedValue(ifdPath, tagId, valueContext, tagType.ByteOrder())
					if err2 != nil {
						return err2
					} else {
						valueString = fmt.Sprintf("%v", value)
					}
				} else {
					valueString, err = tagType.ResolveAsString(valueContext, true)
					if err != nil {
						return err
					}

					value = valueString
				}

				entry := IfdEntry{
					IfdPath:     ifdPath,
					FqIfdPath:   fqIfdPath,
					IfdIndex:    ifdIndex,
					TagId:       tagId,
					TagName:     it.Name,
					TagTypeId:   tagType.Type(),
					TagTypeName: tagType.Name(),
					UnitCount:   valueContext.UnitCount,
					Value:       value,
					ValueString: valueString,
				}

				entries = append(entries, entry)

				return nil
			}

			_, err = exif.Visit(exif.IfdStandard, im, ti, rawExif, visitor)
			if err == nil {
				for _, e := range entries {
					x := map[string]string{}
					x[e.TagName] = e.ValueString
					exifEntries = append(exifEntries, x)
				}
			}

		}

		os.Remove(filePath)
		output, _ := json.Marshal(exifEntries)
		return string(output)
	}
}

func downloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
