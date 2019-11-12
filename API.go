package vdsinaAPI

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)



type AuthInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenInfo struct {
	Status    string `json:"status"`
	StatusMsg string `json:"status_msg"`
	Data      struct {
		Token string `json:"token"`
	} `json:"data"`
}

type Account struct {
	Status    string `json:"status"`
	StatusMsg string `json:"status_msg"`
	Data      struct {
		Account struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"account"`
		Created  string `json:"created"`
		Forecast string `json:"forecast"`
		Can      struct {
			AddUser       bool `json:"add_user"`
			AddService    bool `json:"add_service"`
			ConvertToCash bool `json:"convert_to_cash"`
		} `json:"can"`
	} `json:"data"`
}

type Balance struct {
	Status    string `json:"status"`
	StatusMsg string `json:"status_msg"`
	Data      struct {
		Real    string `json:"real"`
		Bonus   string `json:"bonus"`
		Partner string `json:"partner"`
	} `json:"data"`
}

type Limits struct {
	Status    string `json:"status"`
	StatusMsg string `json:"status_msg"`
	Data      struct {
		Server struct {
			Max int `json:"max"`
			Now int `json:"now"`
		} `json:"server"`
		ServerIP4 struct {
			Max      int `json:"max"`
			ChildMax int `json:"child_max"`
			Now      int `json:"now"`
		} `json:"server-ip4"`
		ServerIP6 struct {
			Max      int `json:"max"`
			ChildMax int `json:"child_max"`
			Now      int `json:"now"`
		} `json:"server-ip6"`
		Iso struct {
			Max int `json:"max"`
			Now int `json:"now"`
		} `json:"iso"`
		Backup struct {
			Max int `json:"max"`
			Now int `json:"now"`
		} `json:"backup"`
		Ssl struct {
			Max int `json:"max"`
			Now int `json:"now"`
		} `json:"ssl"`
		Domain struct {
			Max int `json:"max"`
			Now int `json:"now"`
		} `json:"domain"`
		DNS struct {
			Max int `json:"max"`
			Now int `json:"now"`
		} `json:"dns"`
	} `json:"data"`
}

type ServerGroup struct {
	Status    string `json:"status"`
	StatusMsg string `json:"status_msg"`
	Data      []struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Type        string `json:"type"`
		Description string `json:"description"`
	} `json:"data"`
}

type DataCenterList struct {
	Status    string `json:"status"`
	StatusMsg string `json:"status_msg"`
	Data      []struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Country string `json:"country"`
		Active  bool   `json:"active"`
	} `json:"data"`
}

type TemplateOfOS struct {
	Status    string `json:"status"`
	StatusMsg string `json:"status_msg"`
	Data      []struct {
		ID             int    `json:"id"`
		Name           string `json:"name"`
		Image          string `json:"image"`
		Active         bool   `json:"active"`
		HasInstruction bool   `json:"has_instruction"`
		SSHKey         bool   `json:"ssh-key"`
		ServerPlan     []int  `json:"server-plan"`
	} `json:"data"`
}

type ServerPlans struct {
	Status    string `json:"status"`
	StatusMsg string `json:"status_msg"`
	Data      []struct {
		ID          int     `json:"id"`
		Name        string  `json:"name"`
		Cost        float64 `json:"cost"`
		Period      string  `json:"period"`
		MinMoney    int     `json:"min_money"`
		CanBonus    bool    `json:"can_bonus"`
		Description string  `json:"description"`
		Data        struct {
			VCPU      string `json:"vCPU"`
			Workload  string `json:"Нагрузка"`
			RAM       string `json:"RAM"`
			SSD       string `json:"SSD"`
			Bandwidth string `json:"Трафик"`
		} `json:"data"`
		Selected     bool    `json:"selected"`
		Active       bool    `json:"active"`
		Enable       bool    `json:"enable"`
		BackupCost   float64 `json:"backup_cost"`
		BackupPeriod string  `json:"backup_period"`
		ServerGroup  int     `json:"server-group"`
	} `json:"data"`
}

type GETSSHKeys struct {
	Status    string `json:"status"`
	StatusMsg string `json:"status_msg"`
	Data      []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"data"`
}

type GETSSHKey struct {
	Status    string `json:"status"`
	StatusMsg string `json:"status_msg"`
	Data      struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Status string `json:"status"`
		Data   string `json:"data"`
	} `json:"data"`
}

type SSHKey struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type SimpleResponse struct {
	Status    string      `json:"status"`
	StatusMsg string      `json:"status_msg"`
	Data      interface{} `json:"data"`
}

type ISOList struct {
	Status    string `json:"status"`
	StatusMsg string `json:"status_msg"`
	Data      []struct {
		ID         int    `json:"id"`
		Name       string `json:"name"`
		FullName   string `json:"full_name"`
		Created    string `json:"created"`
		Updated    string `json:"updated"`
		End        string `json:"end"`
		Status     string `json:"status"`
		StatusText string `json:"status_text"`
		File       struct {
			Size string `json:"size"`
			MD5  string `json:"md5"`
		} `json:"file"`
		Attached bool        `json:"attached"`
		Server   interface{} `json:"server"`
		Can      struct {
			Delete bool `json:"delete"`
		} `json:"can"`
	} `json:"data"`
}

type ISOSingle struct {
	Status    string `json:"status"`
	StatusMsg string `json:"status_msg"`
	Data      struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Status  string `json:"status"`
		Created string `json:"created"`
		Updated string `json:"updated"`
		End     string `json:"end"`
		File    struct {
			Size string `json:"size"`
			MD5  string `json:"md5"`
			URL  string `json:"url"`
		} `json:"file"`
		Attached bool        `json:"attached"`
		Server   interface{} `json:"server"`
		Can      struct {
			Delete bool `json:"delete"`
		} `json:"can"`
	} `json:"data"`
}

type ISO struct {
	url string `json:"url"`
}

type ISOCreated struct {
	Status    string `json:"status"`
	StatusMsg string `json:"status_msg"`
	Data      struct {
		ID int `json:"id"`
	} `json:"data"`
}

type Server struct {
	Status    string `json:"status"`
	StatusMsg string `json:"status_msg"`
	Data      struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Status      string `json:"status"`
		Created     string `json:"created"`
		Updated     string `json:"updated"`
		End         string `json:"end"`
		Autoprolong bool   `json:"autoprolong"`
		IP          []struct {
			ID      int    `json:"id"`
			IP      string `json:"ip"`
			Type    string `json:"type"`
			Host    string `json:"host"`
			Gateway string `json:"gateway"`
			Netmask string `json:"netmask"`
			Mac     string `json:"mac"`
		} `json:"ip"`
		IPLocal    interface{} `json:"ip_local"`
		Host       string      `json:"host"`
		ServerPlan struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"server-plan"`
		ServerGroup struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"server-group"`
		Template struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"template"`
		Datacenter struct {
			ID      int    `json:"id"`
			Name    string `json:"name"`
			Country string `json:"country"`
		} `json:"datacenter"`
		SSHKey interface{} `json:"ssh-key"`
		Can    struct {
			Reboot  bool `json:"reboot"`
			Update  bool `json:"update"`
			Delete  bool `json:"delete"`
			Prolong bool `json:"prolong"`
			Backup  bool `json:"backup"`
			IPLocal bool `json:"ip_local"`
		} `json:"can"`
		Bandwidth struct {
			CurrentMonth string `json:"current_month"`
			PastMonth    string `json:"past_month"`
		} `json:"bandwidth"`
	} `json:"data"`
}

type ServerList struct {
	Status    string `json:"status"`
	StatusMsg string `json:"status_msg"`
	Data      []struct {
		ID         int         `json:"id"`
		Name       string      `json:"name"`
		FullName   string      `json:"full_name"`
		Created    string      `json:"created"`
		Updated    string      `json:"updated"`
		End        string      `json:"end"`
		Status     string      `json:"status"`
		StatusText interface{} `json:"status_text"`
		IP         struct {
			ID      int    `json:"id"`
			IP      string `json:"ip"`
			Type    string `json:"type"`
			Host    string `json:"host"`
			Gateway string `json:"gateway"`
			Netmask string `json:"netmask"`
			Mac     string `json:"mac"`
		} `json:"ip"`
		ServerPlan struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"server-plan"`
		Template struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"template"`
		Datacenter struct {
			ID      int    `json:"id"`
			Name    string `json:"name"`
			Country string `json:"country"`
		} `json:"datacenter"`
		Can struct {
			Reboot  bool `json:"reboot"`
			Update  bool `json:"update"`
			Delete  bool `json:"delete"`
			Prolong bool `json:"prolong"`
		} `json:"can"`
	} `json:"data"`
}

func DoRequestAndDecode(method, url, token string, body io.Reader, decodeTo interface{}) error {
	req, err := http.NewRequest(method, url, body)

	if err != nil {
		return err
	}

	if token != "" {
		req.Header.Set("Authorization", token)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(&decodeTo)
}

func GetToken(info AuthInfo) (token *TokenInfo, err error) {
	payloadBytes, err := json.Marshal(info)
	if err != nil {
		return token, err
	}
	body := bytes.NewReader(payloadBytes)

	err = DoRequestAndDecode("POST", "https://userapi.vdsina.ru/v1/auth", "", body, &token)
	if err != nil {
		return token, err
	}
	return token, nil
}

func UseToken(tkn string) *TokenInfo {
	return &TokenInfo{
		Data: struct {
			Token string `json:"token"`
		}{tkn},
	}
}

func (t TokenInfo) Info() (acc *Account, err error) {
	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

	err = DoRequestAndDecode("GET", "https://userapi.vdsina.ru/v1/account", t.Data.Token, nil, &acc)
	if err != nil {
		return acc, err
	}

	return acc, nil
}

func (t TokenInfo) Balance() (bal *Balance, err error) {
	err = DoRequestAndDecode("GET", "https://userapi.vdsina.ru/v1/account.balance", t.Data.Token, nil, &bal)
	if err != nil {
		return bal, err
	}
	return bal, nil
}

func (t TokenInfo) Limits() (lim *Limits, err error) {
	err = DoRequestAndDecode("GET", "https://userapi.vdsina.ru/v1/account.limit", t.Data.Token, nil, &lim)
	if err != nil {
		return lim, err
	}

	return lim, nil
}

func (t TokenInfo) ServerGroup() (srv *ServerGroup, err error) {
	err = DoRequestAndDecode("GET", "https://userapi.vdsina.ru/v1/server-group", t.Data.Token, nil, &srv)
	if err != nil {
		return srv, err
	}
	return srv, nil
}

func (t TokenInfo) DataCenterList(id int) (dc *DataCenterList, err error) {
	idString := strconv.FormatInt(int64(id), 10)
	err = DoRequestAndDecode("GET", "https://userapi.vdsina.ru/v1/datacenter/"+idString, t.Data.Token, nil, &dc)
	if err != nil {
		return dc, err
	}
	return dc, nil
}

func (t TokenInfo) Template(id int) (tmpl *TemplateOfOS, err error) {
	idString := strconv.FormatInt(int64(id), 10)
	err = DoRequestAndDecode("GET", "https://userapi.vdsina.ru/v1/template/"+idString, t.Data.Token, nil, &tmpl)
	if err != nil {
		return tmpl, err
	}
	return tmpl, nil
}

func (t TokenInfo) ServerPlans(id int) (srvplan *ServerPlans, err error) {
	idString := strconv.FormatInt(int64(id), 10)
	err = DoRequestAndDecode("GET", "https://userapi.vdsina.ru/v1/server-plan/"+idString, t.Data.Token, nil, &srvplan)
	if err != nil {
		return srvplan, err
	}
	return srvplan, nil
}

func (t TokenInfo) GetSSHKeys() (sshKeys *GETSSHKeys, err error) {
	err = DoRequestAndDecode("GET", "https://userapi.vdsina.ru/v1/ssh-key/", t.Data.Token, nil, &sshKeys)
	if err != nil {
		return sshKeys, err
	}
	return sshKeys, nil
}

func (t TokenInfo) GetSSHKey(id int) (sshKey *GETSSHKey, err error) {
	idString := strconv.FormatInt(int64(id), 10)
	err = DoRequestAndDecode("GET", "https://userapi.vdsina.ru/v1/ssh-key/"+idString, t.Data.Token, nil, &sshKey)
	if err != nil {
		return sshKey, err
	}
	return sshKey, nil
}

func (t TokenInfo) EditSSHKey(id int, info SSHKey) (sshKey *SimpleResponse, err error) {
	idString := strconv.FormatInt(int64(id), 10)

	payloadBytes, err := json.Marshal(info)
	if err != nil {
		return sshKey, err
	}
	body := bytes.NewReader(payloadBytes)

	err = DoRequestAndDecode("PUT", "https://userapi.vdsina.ru/v1/ssh-key/"+idString, t.Data.Token, body, &sshKey)
	if err != nil {
		return sshKey, err
	}

	return sshKey, nil
}

func (t TokenInfo) NewSSHKey(info SSHKey) (sshKey *SimpleResponse, err error) {
	payloadBytes, err := json.Marshal(info)
	if err != nil {
		return sshKey, err
	}
	body := bytes.NewReader(payloadBytes)

	err = DoRequestAndDecode("POST", "https://userapi.vdsina.ru/v1/ssh-key/", t.Data.Token, body, &sshKey)
	if err != nil {
		return sshKey, err
	}

	return sshKey, nil
}

func (t TokenInfo) DeleteSSHKey(id int) (sshKey *SimpleResponse, err error) {
	idString := strconv.FormatInt(int64(id), 10)
	err = DoRequestAndDecode("DELETE", "https://userapi.vdsina.ru/v1/ssh-key/"+idString, t.Data.Token, nil, &sshKey)
	if err != nil {
		return sshKey, err
	}
	return sshKey, nil
}

func (t TokenInfo) GetISOs() (ISOs *ISOList, err error) {
	err = DoRequestAndDecode("GET", "https://userapi.vdsina.ru/v1/iso/", t.Data.Token, nil, &ISOs)
	if err != nil {
		return ISOs, err
	}
	return ISOs, nil
}

func (t TokenInfo) GetISO(id int) (ISO *ISOSingle, err error) {
	idString := strconv.FormatInt(int64(id), 10)
	err = DoRequestAndDecode("GET", "https://userapi.vdsina.ru/v1/iso/"+idString, t.Data.Token, nil, &ISO)
	if err != nil {
		return ISO, err
	}
	return ISO, nil
}

func (t TokenInfo) NewISO(info ISO) (ISO *ISOCreated, err error) {
	payloadBytes, err := json.Marshal(info)
	if err != nil {
		return ISO, err
	}
	body := bytes.NewReader(payloadBytes)

	err = DoRequestAndDecode("POST", "https://userapi.vdsina.ru/v1/iso/", t.Data.Token, body, &ISO)
	if err != nil {
		return ISO, err
	}

	return ISO, nil
}

func (t TokenInfo) DeleteISO(id int) (ISO *SimpleResponse, err error) {
	idString := strconv.FormatInt(int64(id), 10)
	err = DoRequestAndDecode("DELETE", "https://userapi.vdsina.ru/v1/iso/"+idString, t.Data.Token, nil, &ISO)
	if err != nil {
		return ISO, err
	}
	return ISO, nil
}