// Code generated by goctl. DO NOT EDIT.
package types

type BoolResp struct {
	Result bool `json:"result"`
}

type IdReq struct {
	Id uint64 `json:"id"`
}

type PasswordLoginReq struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type TokenResp struct {
	Token  string `json:"token"`
	Result bool   `json:"result"`
}
