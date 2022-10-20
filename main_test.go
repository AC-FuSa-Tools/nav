package main

import (
	"testing"
	"path/filepath"
	"os"
	"runtime"
	)

func Compare_configs(c1 configuration, c2 configuration) (bool){

	res:=true
	res = res && c1.DBURL == c2.DBURL
	res = res && c1.DBPort == c2.DBPort
	res = res && c1.DBUser == c2.DBUser
	res = res && c1.DBPassword == c2.DBPassword
	res = res && c1.DBTargetDB == c2.DBTargetDB
	res = res && c1.Symbol == c2.Symbol
	res = res && c1.Instance == c2.Instance
	res = res && c1.Mode == c2.Mode
	res = res && c1.MaxDepth == c2.MaxDepth
	res = res && c1.Jout == c2.Jout
	res = res && len(c1.Excluded) == len(c2.Excluded)
	for i, item := range c1.Excluded {
		res = res && item == c2.Excluded[i]
		}
	return res
}

func TestConfig(t *testing.T){

var     Test_config  configuration = configuration{
        DBURL:          "dummy",
        DBPort:         1234,
        DBUser:         "dummy",
        DBPassword:     "dummy",
        DBTargetDB:     "dummy",
        Symbol:         "dummy",
        Instance:       1234,
        Mode:           1234,
        Excluded:       []string{"dummy1", "dummy2", "dummy3"},
        MaxDepth:       1234,              //0: no limit
        Jout:           "JsonOutputPlain",
        cmdlineNeeds:   map[string] bool{},
        }


	os.Args=[]string{"nav"}
	conf, err := args_parse(cmd_line_item_init())
	if err==nil {
		t.Error("Error validating empty command line input against mandatory args")
		}

	if !Compare_configs(conf, Default_config) {
		t.Error("unexpected change in default config")
		}

	os.Args=[]string{"nav", "-i", "1", "-s"}
        conf, err = args_parse(cmd_line_item_init())
	if err==nil {
		t.Error("error Missing switch argument not detected", conf)
		}

	os.Args=[]string{"nav", "-i", "a", "-s", "symb"}
        conf, err = args_parse(cmd_line_item_init())
	if err==nil {
		t.Error("error switch arg type mismatch not detected", conf)
		}

	os.Args=[]string{"nav", "-i", "a", "-s", "symb", "-f", }
        conf, err = args_parse(cmd_line_item_init())
	if err==nil {
		t.Error("error Missing optional switch argument not detected", conf)
		}

	_, filename, _, _ := runtime.Caller(0)
	current := filepath.Dir(filename)

	os.Args=[]string{"nav", "-i", "1", "-s", "symb", "-f", current+"/t_files/dummy.json"}
        conf, err = args_parse(cmd_line_item_init())
	if err==nil {
		t.Error("undetected missing file", conf)
		}
	if !Compare_configs(conf, Default_config) {
		t.Error("unexpected change in default config")
		}

	os.Args=[]string{"nav", "-i", "1", "-s", "symb", "-f", current+"/t_files/test1.json"}
        conf, err = args_parse(cmd_line_item_init())
	if err!=nil {
		t.Error("Unexpected conf error while reading from existing file", err, current+"/t_files/test1.json")
		}
	if !Compare_configs(conf, Test_config) {
		t.Error("unexpected difference between actual and loaded config", conf, Test_config)
		}

	tmp:=Test_config
	tmp.DBUser="new"
	os.Args=[]string{"nav", "-i", "1", "-s", "symb", "-f", current+"/t_files/test1.json", "-u", "new"}
        conf, err = args_parse(cmd_line_item_init())
	if err!=nil {
		t.Error("Unexpected conf error while reading from existing file", err, current+"/t_files/test1.json")
		}
	if !Compare_configs(conf, tmp) {
		t.Error("unexpected difference between actual and loaded modified config", conf, Test_config)
		}

}
