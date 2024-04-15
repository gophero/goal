package ord

import (
	"encoding/json"
	"fmt"
	"github.com/gophero/goal/cmd"

	"github.com/pkg/errors"
)

var InscribeError = fmt.Errorf("inscribe error")

const (
	minPostage = 330
)

type Client interface {
	InscribeRaw(req *InscribeRequest) (*InscribeResult, error)
	Inscribe(wallet string, destination string, feeRate uint, file string, opts ...InscribeOption) (*InscribeResult, error)
	InscribeBatch(wallet string, feeRate uint, batchFile string, opts ...InscribeOption) (*InscribeResult, error)
	QueryInscriptions(req *InscriptionQueryRequest) ([]*Inscription, error)
}

type InscribeOption func(req *InscribeRequest)

type CBOR map[string]any

type BaseRequest struct {
	Wallet string
}

type InscribeRequest struct {
	BaseRequest
	Destination           string
	FeeRate               uint
	ParentInscriptionId   string
	DelegateInscriptionId string
	Postage               uint64
	// Reinscribe            bool
	// MetaData              CBOR
	File      string
	BatchFile string
	DryRun    bool
}

type InscriptionQueryRequest struct {
	BaseRequest
}

func NewInscribeRequest(wallet string, feeRate uint) *InscribeRequest {
	return &InscribeRequest{BaseRequest: BaseRequest{Wallet: wallet}, FeeRate: feeRate}
}

func (req *InscribeRequest) checkBaseParam() error {
	if req.Wallet == "" {
		return errors.Wrap(InscribeError, "need a wallet")
	}
	if req.FeeRate <= 0 {
		return errors.Wrap(InscribeError, "need to set feeRate")
	}
	if req.File == "" && req.BatchFile == "" {
		return errors.Wrap(InscribeError, "need a file or batchFile")
	}
	return nil
}

func (req *InscribeRequest) checkInscribeParam() error {
	if err := req.checkBaseParam(); err != nil {
		return err
	}
	if req.Destination == "" {
		return errors.Wrap(InscribeError, "need a destination")
	}
	if req.Postage > 0 && req.Postage < minPostage {
		return errors.Wrap(InscribeError, "postage not enough")
	}
	return nil
}

func (req *InscribeRequest) checkBatchInscribeParam() error {
	if err := req.checkBaseParam(); err != nil {
		return err
	}
	if req.BatchFile == "" {
		return errors.Wrap(InscribeError, "need a batch file, see https://docs.ordinals.com/guides/batch-inscribing.html")
	}
	return nil
}

type ordClient struct {
	setting OrdSetting
}

func NewClient(setting OrdSetting) Client {
	return &ordClient{setting: setting}
}

func (o *ordClient) InscribeRaw(req *InscribeRequest) (*InscribeResult, error) {
	if err := req.checkBaseParam(); err != nil {
		return nil, err
	}
	command, err := Cmd.inscribeCmd(o.setting, req)
	fmt.Println("command:", command)
	if err != nil {
		return nil, err
	}
	r := cmd.Exec(command)
	fmt.Println("result:", r)
	return nil, nil
}

func (o *ordClient) Inscribe(wallet string, destination string, feeRate uint, file string, opts ...InscribeOption) (*InscribeResult, error) {
	req := &InscribeRequest{}
	req.Wallet = wallet
	req.Destination = destination
	req.FeeRate = feeRate
	req.File = file
	for _, v := range opts {
		v(req)
	}
	if err := req.checkInscribeParam(); err != nil {
		return nil, err
	}
	command, err := Cmd.inscribeCmd(o.setting, req)
	if err != nil {
		return nil, err
	}
	r := cmd.Exec("bash", "-c", command)
	var ir InscribeResult
	if err = json.Unmarshal([]byte(r), &ir); err != nil {
		return nil, fmt.Errorf("inscribe failed, result: %s, err: %w", r, err)
	}
	return &ir, nil
}

func (o *ordClient) InscribeBatch(wallet string, feeRate uint, batchFile string, opts ...InscribeOption) (*InscribeResult, error) {
	req := &InscribeRequest{}
	req.Wallet = wallet
	req.FeeRate = feeRate
	req.BatchFile = batchFile
	for _, v := range opts {
		v(req)
	}
	if err := req.checkBatchInscribeParam(); err != nil {
		return nil, err
	}
	command, err := Cmd.inscribeCmd(o.setting, req)
	if err != nil {
		return nil, err
	}
	r := cmd.Exec("bash", "-c", command)
	var ir InscribeResult
	if err = json.Unmarshal([]byte(r), &ir); err != nil {
		return nil, fmt.Errorf("inscribe failed, result: %s, err: %w", r, err)
	}
	return &ir, nil
}

func (o *ordClient) QueryInscriptions(req *InscriptionQueryRequest) ([]*Inscription, error) {
	if req.Wallet == "" {
		return []*Inscription{}, fmt.Errorf("need a wallet name")
	}
	cmdstr, err := Cmd.queryInscriptionCmd(o.setting, req)
	if err != nil {
		return []*Inscription{}, err
	}
	r := cmd.Exec("bash", "-c", cmdstr)
	var is []*Inscription
	err = json.Unmarshal([]byte(r), &is)
	if err != nil {
		return []*Inscription{}, fmt.Errorf("query failed, data: %s, error: %w", r, err)
	}
	return is, nil
}

func Postage(postage uint64) InscribeOption {
	return func(req *InscribeRequest) {
		req.Postage = postage
	}
}

func ParentInscriptionId(pid string) InscribeOption {
	return func(req *InscribeRequest) {
		req.ParentInscriptionId = pid
	}
}

func DelegateInscriptionId(did string) InscribeOption {
	return func(req *InscribeRequest) {
		req.DelegateInscriptionId = did
	}
}

func DryRun() InscribeOption {
	return func(req *InscribeRequest) {
		req.DryRun = true
	}
}
