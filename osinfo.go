package osinfo

import (
	"bufio"
	"errors"
	"os"
	"strings"
	"strconv"
)

const OS_RELEASE_INFO_FILE_PATH string = "./os-release_valid"
var os_release_info_file_path = OS_RELEASE_INFO_FILE_PATH

type OSDetails struct {
	Name             string
	Version          string
	VersionId        string
	Id               string
	IdLide           string
	AnsiColor        string
	BuildId          string
	PrettyName       string
	CpeName          string
	Variant          string
	VariantId        string
	HomeUrl          string
	BugReportUrl     string
	SupportUrl       string
	PrivacyPolicyUrl string
	Unknown          map[string]string
}

func (osd *OSDetails) Read() error {
	var counter int = 1

	file, err := os.Open(os_release_info_file_path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, "=")
		if len(values) > 2 {
			return errors.New("Line "+strconv.Itoa(counter)+" of "+os_release_info_file_path+" has too many values!")
		} else {
			values[0] = strings.TrimSpace(values[0])
			values[1] = strings.TrimSpace(values[1])
			values[1] = strings.TrimPrefix(values[1], "\"")
			values[1] = strings.TrimSuffix(values[1], "\"")

			switch values[0] {
			case "NAME":
				osd.Name = values[1]
			case "VERSION":
				osd.Version = values[1]
			case "VERSION_ID":
				osd.VersionId = values[1]
			case "ID":
				osd.Id = values[1]
			case "ID_LIKE":
				osd.IdLide = values[1]
			case "ANSI_COLOR":
				osd.AnsiColor = values[1]
			case "PRETTY_NAME":
				osd.PrettyName = values[1]
			case "CPE_NAME":
				osd.CpeName = values[1]
			case "BUILD_ID":
				osd.BuildId = values[1]
			case "VARIANT":
				osd.Variant = values[1]
			case "VARIANT_ID":
				osd.VariantId = values[1]
			case "HOME_URL":
				osd.HomeUrl = values[1]
			case "SUPPORT_URL":
				osd.SupportUrl = values[1]
			case "BUG_REPORT_URL":
				osd.BugReportUrl = values[1]
			case "PRIVACY_POLICY_URL":
				osd.PrivacyPolicyUrl = values[1]
			default:
				if osd.Unknown == nil {
					osd.Unknown = make(map[string]string)
				}
				osd.Unknown[values[0]] = values[1]
			}

			counter++
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func (osd *OSDetails) SetPath(path string) {
	os_release_info_file_path = path
}
