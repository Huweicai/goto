package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

type ModelType int

const (
	configName      = "config.yaml"
	LoopBackKey     = "@"
	TipSeparator    = ", "
	ScalarSeparator = " @@@ "

	TypeUnknown ModelType = iota
	TypeVector
	TypeScalar
)

// Nest the core config
type Nest struct {
	Data map[string]interface{}
}

type DeserializedScalar struct {
	Val       string
	Frequency int64
}

type Key struct {
	Key       string
	Val       string
	Frequency int64
}

func DeserializeScalar(s string) *DeserializedScalar {
	ss := strings.Split(s, ScalarSeparator)
	if len(ss) != 2 {
		return &DeserializedScalar{
			Val: s,
		}
	}

	frequency, err := strconv.ParseInt(ss[0], 10, 64)
	if err != nil {
		log.Printf("parse frequncy '%s' as int64 failed: %v\n", ss[0], err)
		return &DeserializedScalar{
			Val: s,
		}
	}

	return &DeserializedScalar{
		Val:       ss[1],
		Frequency: frequency,
	}
}

func (s *DeserializedScalar) Serialize() string {
	return fmt.Sprintf("%d%s%s", s.Frequency, ScalarSeparator, s.Val)
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
	conf, err := loadConfig(configPath)
	if err != nil {
		return nil, err
	}

	return &Nest{Data: conf}, nil
}

func (n *Nest) GetScalar(paths []string) (*DeserializedScalar, bool) {
	s, ok := getScalar(paths, n.Data)
	if ok {
		return DeserializeScalar(s), true
	}

	return nil, false
}

// IncScalar increase the scalar for usage statistic
func (n *Nest) IncScalar(paths []string) bool {
	return incScalar(n.Data, paths)
}

func incScalar(data interface{}, paths []string) bool {
	if len(paths) == 0 {
		return false
	}
	dataMap, ok := data.(map[string]interface{})
	if !ok {
		return false
	}

	currKey := paths[0]
	paths = paths[1:]

	currVal, ok := dataMap[currKey]
	if !ok {
		return false
	}

	switch typedCurrVal := currVal.(type) {
	case string:
		scalar := DeserializeScalar(typedCurrVal)
		scalar.Frequency++
		dataMap[currKey] = scalar.Serialize()

	case map[string]interface{}:
		incScalar(typedCurrVal, paths)
	}

	return false
}

func (n *Nest) Flush() error {
	return writeConfig(GetConfigPath(), n.Data)
}

func (n *Nest) AddScalar(paths []string, url string) {
	addScalar(n.Data, paths, url)
}

func (n *Nest) ListWithPre(paths []string) []*Key {
	switch len(paths) {
	case 0:
		return findMapPrefix(n.Data, "")
	case 1:
		return findMapPrefix(n.Data, strings.ToLower(paths[0]))
	default:
		out, ok := getByPath(paths[:len(paths)-1], n.Data)
		if !ok {
			return nil
		}

		m, ok := out.(map[string]interface{})
		if !ok {
			return nil
		}

		return findMapPrefix(m, strings.ToLower(paths[len(paths)-1]))
	}

	return nil
}

func findMapPrefix(m map[string]interface{}, pre string) []*Key {
	found := []*Key{}
	for k, v := range m {
		if strings.HasPrefix(k, pre) {
			switch typeOf(v) {
			case TypeVector:
				found = append(found, &Key{
					Key: k,
					Val: extractKeys(v.(map[string]interface{})),
				})
			case TypeScalar:
				scalar := DeserializeScalar(v.(string))
				found = append(found, &Key{
					Key:       k,
					Val:       scalar.Val,
					Frequency: scalar.Frequency,
				})
			}
		}
	}

	return found
}

func extractKeys(m map[string]interface{}) string {
	out := ""
	for k, _ := range m {
		out += k + TipSeparator
	}
	out = strings.TrimSuffix(out, TipSeparator)
	return out
}

func typeOf(i interface{}) ModelType {
	switch reflect.TypeOf(i).Kind() {
	case reflect.Map:
		return TypeVector
	case reflect.String:
		return TypeScalar
	default:
		log.Print("WARNNIG unknown type" + reflect.TypeOf(i).Kind().String())
		return TypeUnknown
	}
}

//load config from yaml
func loadConfig(path string) (conf map[string]interface{}, err error) {
	if !fileExists(path) {
		emt, err := os.Create(path)
		if err != nil {
			return nil, err
		}
		emt.Close()
	}
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

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

//put the config into yaml file
func writeConfig(path string, target interface{}) error {
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
	if len(paths) == 0 || len(m) == 0 {
		return
	}
	v, ok := getByPath(paths, m)
	if !ok {
		return
	}
	switch typeOf(v) {
	case TypeScalar:
		return v.(string), true
	case TypeVector:
		v, ok := v.(map[string]interface{})[LoopBackKey]
		if ok {
			return v.(string), true
		} else {
			return "", false
		}
	default:
		return
	}
}

func getByPath(paths []string, m map[string]interface{}) (out interface{}, ok bool) {
	if len(paths) == 0 || len(m) == 0 {
		return
	}
	out, ok = m[paths[0]]
	if !ok || len(paths) == 1 {
		return
	}
	return getByPath(paths[1:], out.(map[string]interface{}))
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
			//scalar do not need to be lower
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
