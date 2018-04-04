package project

import (
	"testing"
	"fmt"
	"encoding/json"
	"os"
)

func TestNew(t *testing.T) {
	pn := "hk-prod"
	if p := New(pn); p.ProjectName != pn {
		t.Error(fmt.Sprintf("expect projectname = %s, actual %s", pn, p.ProjectName))
	}
}
func TestToJson(t *testing.T){
	p := New("hk-prod")
	pjson := p.ToJson()

	pp := &Project{}
	json.Unmarshal([]byte(pjson), &pp)

	if pp.ProjectName != p.ProjectName{
		t.Error(fmt.Sprintf("expect %s, actual %s", p.ProjectName, pp.ProjectName))
	}
}
func TestLoad(t *testing.T) {
	p := New("hk-prod")
	p.Save()

	pp := Load("hk-prod")
	if pp.ProjectName != p.ProjectName{
		t.Error(fmt.Sprintf("expect %s, actual %s", p.ProjectName, pp.ProjectName))
	}

	os.Remove("./hk-prod.config.json")
}
func TestSave(t *testing.T){
	p := New("hk-prod")
	p.Save()
	if _, err := os.Stat("./" + p.Filename()); os.IsNotExist(err){
		t.Error("expect have file hk-prod.config.json, actual file doesn't exist")
	}else{
		os.Remove("./" + p.Filename())
	}
}


func Test_filename(t *testing.T){
	expect := "hk-name.config.json"
	pn := "hk-name"
	if fn := filename(pn); fn != expect{
		t.Error(fmt.Sprintf("expect %s, actual %s", expect, fn))
	}
}