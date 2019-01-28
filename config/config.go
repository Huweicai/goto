package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

type ModelType int

const (
	configName  = "config.yaml"
	LoopBackKey = "@"

	TypeUnknown ModelType = iota
	TypeVector
	TypeScalar
)

//the core config
type Nest struct {
	Data map[string]interface{}
}

func GetConfigPath() string {
	defaultPath := "./" + configName
	wd, err := os.Getwd()
	if err != nil {
		return defaultPath
	}
	if strings.HasSuffix(wd, "goto") {
		return "./" + configName
	} else if strings.HasSuffix(wd, "handler") || strings.HasSuffix(wd, "config") {
		return "../" + configName
	}
	return defaultPath
}

func NewNest(configPath string) (*Nest, error) {
	conf, err := loadYaml(configPath)
	if err != nil {
		return nil, err
	}
	return &Nest{Data: conf}, nil
}

func (n *Nest) GetScalar(paths []string) (URL string, ok bool) {
	return getScalar(paths, n.Data)
}

func (n *Nest) Flush() error {
	return writeYaml(GetConfigPath(), n.Data)
}

func (n *Nest) AddScalar(paths []string, url string) {
	addScalar(n.Data, paths, url)
}

func typeOf(i interface{}) ModelType {
	switch reflect.TypeOf(i).Kind() {
	case reflect.Map:
		return TypeVector
	case reflect.String:
		return TypeScalar
	default:
		return TypeUnknown
	}
}

//load config from yaml
func loadYaml(path string) (conf map[string]interface{}, err error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(f, &m)
	//to lower to compare
	if err != nil {
		return
	}
	conf = toMapStringInterface(m)
	toLower(conf)
	return
}

//put the config into yaml file
func writeYaml(path string, target interface{}) error {
	out, err := yaml.Marshal(target)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, out, 0644)
}

func addScalar(m map[string]interface{}, paths []string, url string) map[string]interface{} {
	if len(paths) == 0 {
		return nil
	}
	curK := paths[0]
	curV := m[curK]
	//find scalar key
	if len(paths) == 1 {
		if curV != nil && typeOf(curV) == TypeVector {
			k := m[curK]
			k.(map[string]interface{})[LoopBackKey] = url
		} else {
			m[curK] = url
		}
		return m
	}

	if curV == nil {
		m[curK] = addScalar(make(map[string]interface{}), paths[1:], url)
		return m
	}

	if typeOf(curV) == TypeVector {
		m[curK] = addScalar(curV.(map[string]interface{}), paths[1:], url)
	}

	if typeOf(curV) == TypeScalar {
		x := make(map[string]interface{})
		x[LoopBackKey] = curV
		m[curK] = addScalar(x, paths[1:], url)
	}
	return m
}

//get the specified URL
func getScalar(paths []string, m map[string]interface{}) (scalar string, ok bool) {
	if len(paths) == 0 {
		return
	}
	value, got := m[paths[0]]
	if !got {
		return
	}
	switch typeOf(value) {
	case TypeScalar:
		return value.(string), true
	case TypeVector:
		if len(paths) == 1 {
			v, ok := value.(map[string]interface{})[LoopBackKey]
			if ok {
				return v.(string), true
			} else {
				return "", false
			}
		}
		return getScalar(paths[1:], value.(map[string]interface{}))
	default:
		return
	}
}

//convert map[interface{}]intreface{} to map[string]interface{}
func toMapStringInterface(m map[interface{}]interface{}) map[string]interface{} {
	nm := make(map[string]interface{})
	for k, v := range m {
		if reflect.TypeOf(v).Kind() == reflect.Map {
			v = toMapStringInterface(v.(map[interface{}]interface{}))
		}
		nm[k.(string)] = v
	}
	return nm
}

//convert all the upper case to lower in both k and v
func toLower(m map[string]interface{}) map[string]interface{} {
	for k, v := range m {
		delete(m, k)
		switch typeOf(v) {
		case TypeScalar:
			v = strings.ToLower(v.(string))
		case TypeVector:
			v = toLower(v.(map[string]interface{}))
		default:
			continue
		}
		k = strings.ToLower(k)
		m[k] = v
	}
	return m
}
