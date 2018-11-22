package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/ghodss/yaml"
)

// TODO Support flag to read parameters.
const usage = `

template [text|html] <input-yaml-path> <output-go-path>

`

var knownProtoTypes = map[string]struct{}{
	"google.protobuf.Struct": {},
}

// metadata is a combination of read and derived metadata.
type metadata struct {
	Resources       []*entry   `json:"resources"`
	ProtoGoPackages []string   `json:"-"`
	ProtoDefs       []protoDef `json:"-"`
}

// entry in a metadata file
type entry struct {
	Kind           string   `json:"kind"`
	ListKind       string   `json:"listKind"`
	Singular       string   `json:"singular"`
	Plural         string   `json:"plural"`
	ShortNames     []string `json:"shortNames"`
	ShortNamesStr  string   `json:"-"`
	Group          string   `json:"group"`
	Version        string   `json:"version"`
	Proto          string   `json:"proto"`
	Converter      string   `json:"converter"`
	ProtoGoPackage string   `json:"protoPackage"`
}

// proto related metadata
type protoDef struct {
	MessageName string `json:"-"`
}

func main() {
	if len(os.Args) != 4 {
		fmt.Print(usage)
		fmt.Printf("%v\n", os.Args)
		os.Exit(-1)
	}

	// The tool can generate both K8s level, as well as the proto level metadata.
	isText := false
	switch os.Args[1] {
	case "text":
		isText = true
	case "html":
		isText = false
	default:
		fmt.Printf("Unknown target: %v", os.Args[2])
		fmt.Print(usage)
		os.Exit(-1)
	}

	input := os.Args[2]
	output := os.Args[3]

	m, err := readMetadata(input)
	if err != nil {
		fmt.Printf("Error reading metadata: %v", err)
		os.Exit(-2)
	}

	var contents []byte
	if isText {
		contents, err = applyTextTemplate(m)
	} else {
		contents, err = applyHTMLTemplate(m)
	}

	if err != nil {
		fmt.Printf("Error applying template: %v", err)
		os.Exit(-3)
	}

	if err = ioutil.WriteFile(output, contents, os.ModePerm); err != nil {
		fmt.Printf("Error writing output file: %v", err)
		os.Exit(-4)
	}
}

func readMetadata(path string) (*metadata, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to read input file: %v", err)
	}

	var m metadata

	if err = yaml.Unmarshal(b, &m); err != nil {
		return nil, fmt.Errorf("error marshalling input file: %v", err)
	}

	// Auto-complete listkind fields with defaults.
	for _, item := range m.Resources {
		if item.ListKind == "" {
			item.ListKind = item.Kind + "List"
		}

		item.ShortNamesStr = convertShortNames(item.ShortNames)
	}

	// Stable sort based on message name.
	sort.Slice(m.Resources, func(i, j int) bool {
		return strings.Compare(m.Resources[i].Proto, m.Resources[j].Proto) < 0
	})

	// Calculate the Go packages that needs to be imported for the proto types to be registered.
	names := make(map[string]struct{})
	for _, e := range m.Resources {
		if _, found := knownProtoTypes[e.Proto]; e.Proto == "" || found {
			continue
		}

		if e.ProtoGoPackage != "" {
			names[e.ProtoGoPackage] = struct{}{}
			continue
		}

		parts := strings.Split(e.Proto, ".")
		// Remove the first "istio", and the Proto itself
		parts = parts[1 : len(parts)-1]

		p := strings.Join(parts, "/")
		p = fmt.Sprintf("istio.io/api/%s", p)
		names[p] = struct{}{}
	}

	for k := range names {
		m.ProtoGoPackages = append(m.ProtoGoPackages, k)
	}
	sort.Strings(m.ProtoGoPackages)

	// Calculate the proto types that needs to be handled.
	// First, single instance the proto definitions.
	protoDefs := make(map[string]protoDef)
	for _, e := range m.Resources {
		if _, found := knownProtoTypes[e.Proto]; e.Proto == "" || found {
			continue
		}
		defn := protoDef{MessageName: e.Proto}

		if prevDefn, ok := protoDefs[e.Proto]; ok && defn != prevDefn {
			return nil, fmt.Errorf("proto definitions do not match: %+v != %+v", defn, prevDefn)
		}
		protoDefs[e.Proto] = defn
	}

	for _, v := range protoDefs {
		m.ProtoDefs = append(m.ProtoDefs, v)
	}

	// Then, stable sort based on message name.
	sort.Slice(m.ProtoDefs, func(i, j int) bool {
		return strings.Compare(m.ProtoDefs[i].MessageName, m.ProtoDefs[j].MessageName) < 0
	})

	return &m, nil
}

func convertShortNames(shortNames []string) string {
	if len(shortNames) == 0 {
		return ""
	}

	// Convert short names into string.
	sns := make([]string, len(shortNames))
	for i, sn := range shortNames {
		sns[i] = fmt.Sprintf("\"%s\"", sn)
	}

	return fmt.Sprintf("[]string{%s}", strings.Join(sns, ", "))
}
