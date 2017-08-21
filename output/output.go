package output

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/koolay/sqlto/config"
)

// Export to dest
func Export(ctx *config.Context, items []map[string]interface{}) error {
	var err error
	dest := ctx.Output
	reAPIFlag := regexp.MustCompile(`^https?:\/\/[^\s]+`)
	var itemsBytes []byte
	if reAPIFlag.Match([]byte(dest)) {
	} else if strings.HasSuffix(dest, ".json") {
		itemsBytes, err = toBytes(ctx, items)
		if err != nil {
			return err
		}
		toJSON(dest, itemsBytes)
	} else if strings.HasSuffix(dest, ".xls") {
		return toXls(ctx.Output, items)
	} else if strings.HasSuffix(dest, ".csv") {
		fmt.Println("not supported yet")
		return nil
	}
	itemsBytes, err = toBytes(ctx, items)
	if err == nil {
		fmt.Println(string(itemsBytes))
	}
	return err
}

func toBytes(ctx *config.Context, items []map[string]interface{}) ([]byte, error) {
	if ctx.Pretty {
		return json.MarshalIndent(items, "", " ")
	}
	return json.Marshal(items)
}

func getSavePath(input string) (fullname string, err error) {
	dir, filename := filepath.Split(input)
	if dir == "" {
		tmpDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			return "", err
		}
		dir = tmpDir
		fullname = fmt.Sprintf("%s/%s", dir, filename)
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = fmt.Errorf("%s not exists", dir)
	}
	return
}
