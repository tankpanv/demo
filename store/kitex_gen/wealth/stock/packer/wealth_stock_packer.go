// Code generated by thriftgo (0.2.12). DO NOT EDIT.

package packer

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/tankpanv/demo/store/kitex_gen/wealth/stock/article_base"
	"github.com/tankpanv/demo/store/kitex_gen/wealth/stock/base"
	"github.com/tankpanv/demo/store/kitex_gen/wealth/stock/cell_base"
	"strings"
)

type ArticleDatasReq struct {
	AppId       *int16                `thrift:"app_id,1,optional" frugal:"1,optional,i16" json:"app_id,omitempty"`
	Source      *base.Source          `thrift:"source,2,optional" frugal:"2,optional,Source" json:"source,omitempty"`
	VersionCode *int16                `thrift:"version_code,3,optional" frugal:"3,optional,i16" json:"version_code,omitempty"`
	PackerType  *cell_base.PackerType `thrift:"packer_type,4,optional" frugal:"4,optional,PackerType" json:"packer_type,omitempty"`
	Keywords    []string              `thrift:"keywords,5,optional" frugal:"5,optional,list<string>" json:"keywords,omitempty"`
}

func NewArticleDatasReq() *ArticleDatasReq {
	return &ArticleDatasReq{}
}

func (p *ArticleDatasReq) InitDefault() {
	*p = ArticleDatasReq{}
}

var ArticleDatasReq_AppId_DEFAULT int16

func (p *ArticleDatasReq) GetAppId() (v int16) {
	if !p.IsSetAppId() {
		return ArticleDatasReq_AppId_DEFAULT
	}
	return *p.AppId
}

var ArticleDatasReq_Source_DEFAULT base.Source

func (p *ArticleDatasReq) GetSource() (v base.Source) {
	if !p.IsSetSource() {
		return ArticleDatasReq_Source_DEFAULT
	}
	return *p.Source
}

var ArticleDatasReq_VersionCode_DEFAULT int16

func (p *ArticleDatasReq) GetVersionCode() (v int16) {
	if !p.IsSetVersionCode() {
		return ArticleDatasReq_VersionCode_DEFAULT
	}
	return *p.VersionCode
}

var ArticleDatasReq_PackerType_DEFAULT cell_base.PackerType

func (p *ArticleDatasReq) GetPackerType() (v cell_base.PackerType) {
	if !p.IsSetPackerType() {
		return ArticleDatasReq_PackerType_DEFAULT
	}
	return *p.PackerType
}

var ArticleDatasReq_Keywords_DEFAULT []string

func (p *ArticleDatasReq) GetKeywords() (v []string) {
	if !p.IsSetKeywords() {
		return ArticleDatasReq_Keywords_DEFAULT
	}
	return p.Keywords
}
func (p *ArticleDatasReq) SetAppId(val *int16) {
	p.AppId = val
}
func (p *ArticleDatasReq) SetSource(val *base.Source) {
	p.Source = val
}
func (p *ArticleDatasReq) SetVersionCode(val *int16) {
	p.VersionCode = val
}
func (p *ArticleDatasReq) SetPackerType(val *cell_base.PackerType) {
	p.PackerType = val
}
func (p *ArticleDatasReq) SetKeywords(val []string) {
	p.Keywords = val
}

var fieldIDToName_ArticleDatasReq = map[int16]string{
	1: "app_id",
	2: "source",
	3: "version_code",
	4: "packer_type",
	5: "keywords",
}

func (p *ArticleDatasReq) IsSetAppId() bool {
	return p.AppId != nil
}

func (p *ArticleDatasReq) IsSetSource() bool {
	return p.Source != nil
}

func (p *ArticleDatasReq) IsSetVersionCode() bool {
	return p.VersionCode != nil
}

func (p *ArticleDatasReq) IsSetPackerType() bool {
	return p.PackerType != nil
}

func (p *ArticleDatasReq) IsSetKeywords() bool {
	return p.Keywords != nil
}

func (p *ArticleDatasReq) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I16 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.I32 {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.I16 {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 4:
			if fieldTypeId == thrift.I32 {
				if err = p.ReadField4(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 5:
			if fieldTypeId == thrift.LIST {
				if err = p.ReadField5(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ArticleDatasReq[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ArticleDatasReq) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI16(); err != nil {
		return err
	} else {
		p.AppId = &v
	}
	return nil
}

func (p *ArticleDatasReq) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return err
	} else {
		tmp := base.Source(v)
		p.Source = &tmp
	}
	return nil
}

func (p *ArticleDatasReq) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI16(); err != nil {
		return err
	} else {
		p.VersionCode = &v
	}
	return nil
}

func (p *ArticleDatasReq) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return err
	} else {
		tmp := cell_base.PackerType(v)
		p.PackerType = &tmp
	}
	return nil
}

func (p *ArticleDatasReq) ReadField5(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return err
	}
	p.Keywords = make([]string, 0, size)
	for i := 0; i < size; i++ {
		var _elem string
		if v, err := iprot.ReadString(); err != nil {
			return err
		} else {
			_elem = v
		}

		p.Keywords = append(p.Keywords, _elem)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return err
	}
	return nil
}

func (p *ArticleDatasReq) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("ArticleDatasReq"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}
		if err = p.writeField4(oprot); err != nil {
			fieldId = 4
			goto WriteFieldError
		}
		if err = p.writeField5(oprot); err != nil {
			fieldId = 5
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *ArticleDatasReq) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetAppId() {
		if err = oprot.WriteFieldBegin("app_id", thrift.I16, 1); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteI16(*p.AppId); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *ArticleDatasReq) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetSource() {
		if err = oprot.WriteFieldBegin("source", thrift.I32, 2); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteI32(int32(*p.Source)); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *ArticleDatasReq) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetVersionCode() {
		if err = oprot.WriteFieldBegin("version_code", thrift.I16, 3); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteI16(*p.VersionCode); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *ArticleDatasReq) writeField4(oprot thrift.TProtocol) (err error) {
	if p.IsSetPackerType() {
		if err = oprot.WriteFieldBegin("packer_type", thrift.I32, 4); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteI32(int32(*p.PackerType)); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 end error: ", p), err)
}

func (p *ArticleDatasReq) writeField5(oprot thrift.TProtocol) (err error) {
	if p.IsSetKeywords() {
		if err = oprot.WriteFieldBegin("keywords", thrift.LIST, 5); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteListBegin(thrift.STRING, len(p.Keywords)); err != nil {
			return err
		}
		for _, v := range p.Keywords {
			if err := oprot.WriteString(v); err != nil {
				return err
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 5 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 5 end error: ", p), err)
}

func (p *ArticleDatasReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ArticleDatasReq(%+v)", *p)
}

func (p *ArticleDatasReq) DeepEqual(ano *ArticleDatasReq) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.AppId) {
		return false
	}
	if !p.Field2DeepEqual(ano.Source) {
		return false
	}
	if !p.Field3DeepEqual(ano.VersionCode) {
		return false
	}
	if !p.Field4DeepEqual(ano.PackerType) {
		return false
	}
	if !p.Field5DeepEqual(ano.Keywords) {
		return false
	}
	return true
}

func (p *ArticleDatasReq) Field1DeepEqual(src *int16) bool {

	if p.AppId == src {
		return true
	} else if p.AppId == nil || src == nil {
		return false
	}
	if *p.AppId != *src {
		return false
	}
	return true
}
func (p *ArticleDatasReq) Field2DeepEqual(src *base.Source) bool {

	if p.Source == src {
		return true
	} else if p.Source == nil || src == nil {
		return false
	}
	if *p.Source != *src {
		return false
	}
	return true
}
func (p *ArticleDatasReq) Field3DeepEqual(src *int16) bool {

	if p.VersionCode == src {
		return true
	} else if p.VersionCode == nil || src == nil {
		return false
	}
	if *p.VersionCode != *src {
		return false
	}
	return true
}
func (p *ArticleDatasReq) Field4DeepEqual(src *cell_base.PackerType) bool {

	if p.PackerType == src {
		return true
	} else if p.PackerType == nil || src == nil {
		return false
	}
	if *p.PackerType != *src {
		return false
	}
	return true
}
func (p *ArticleDatasReq) Field5DeepEqual(src []string) bool {

	if len(p.Keywords) != len(src) {
		return false
	}
	for i, v := range p.Keywords {
		_src := src[i]
		if strings.Compare(v, _src) != 0 {
			return false
		}
	}
	return true
}

type ArticleDatasResp struct {
	ArticleDatas []*article_base.ArticleData `thrift:"ArticleDatas,1,optional" frugal:"1,optional,list<article_base.ArticleData>" json:"ArticleDatas,omitempty"`
}

func NewArticleDatasResp() *ArticleDatasResp {
	return &ArticleDatasResp{}
}

func (p *ArticleDatasResp) InitDefault() {
	*p = ArticleDatasResp{}
}

var ArticleDatasResp_ArticleDatas_DEFAULT []*article_base.ArticleData

func (p *ArticleDatasResp) GetArticleDatas() (v []*article_base.ArticleData) {
	if !p.IsSetArticleDatas() {
		return ArticleDatasResp_ArticleDatas_DEFAULT
	}
	return p.ArticleDatas
}
func (p *ArticleDatasResp) SetArticleDatas(val []*article_base.ArticleData) {
	p.ArticleDatas = val
}

var fieldIDToName_ArticleDatasResp = map[int16]string{
	1: "ArticleDatas",
}

func (p *ArticleDatasResp) IsSetArticleDatas() bool {
	return p.ArticleDatas != nil
}

func (p *ArticleDatasResp) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.LIST {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ArticleDatasResp[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ArticleDatasResp) ReadField1(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return err
	}
	p.ArticleDatas = make([]*article_base.ArticleData, 0, size)
	for i := 0; i < size; i++ {
		_elem := article_base.NewArticleData()
		if err := _elem.Read(iprot); err != nil {
			return err
		}

		p.ArticleDatas = append(p.ArticleDatas, _elem)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return err
	}
	return nil
}

func (p *ArticleDatasResp) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("ArticleDatasResp"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *ArticleDatasResp) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetArticleDatas() {
		if err = oprot.WriteFieldBegin("ArticleDatas", thrift.LIST, 1); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteListBegin(thrift.STRUCT, len(p.ArticleDatas)); err != nil {
			return err
		}
		for _, v := range p.ArticleDatas {
			if err := v.Write(oprot); err != nil {
				return err
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *ArticleDatasResp) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ArticleDatasResp(%+v)", *p)
}

func (p *ArticleDatasResp) DeepEqual(ano *ArticleDatasResp) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.ArticleDatas) {
		return false
	}
	return true
}

func (p *ArticleDatasResp) Field1DeepEqual(src []*article_base.ArticleData) bool {

	if len(p.ArticleDatas) != len(src) {
		return false
	}
	for i, v := range p.ArticleDatas {
		_src := src[i]
		if !v.DeepEqual(_src) {
			return false
		}
	}
	return true
}

type StorePackerService interface {
	ArticleDatas(ctx context.Context, request *ArticleDatasReq) (r *ArticleDatasResp, err error)
}

type StorePackerServiceClient struct {
	c thrift.TClient
}

func NewStorePackerServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *StorePackerServiceClient {
	return &StorePackerServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewStorePackerServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *StorePackerServiceClient {
	return &StorePackerServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewStorePackerServiceClient(c thrift.TClient) *StorePackerServiceClient {
	return &StorePackerServiceClient{
		c: c,
	}
}

func (p *StorePackerServiceClient) Client_() thrift.TClient {
	return p.c
}

func (p *StorePackerServiceClient) ArticleDatas(ctx context.Context, request *ArticleDatasReq) (r *ArticleDatasResp, err error) {
	var _args StorePackerServiceArticleDatasArgs
	_args.Request = request
	var _result StorePackerServiceArticleDatasResult
	if err = p.Client_().Call(ctx, "ArticleDatas", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

type StorePackerServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      StorePackerService
}

func (p *StorePackerServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *StorePackerServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *StorePackerServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewStorePackerServiceProcessor(handler StorePackerService) *StorePackerServiceProcessor {
	self := &StorePackerServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("ArticleDatas", &storePackerServiceProcessorArticleDatas{handler: handler})
	return self
}
func (p *StorePackerServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x
}

type storePackerServiceProcessorArticleDatas struct {
	handler StorePackerService
}

func (p *storePackerServiceProcessorArticleDatas) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := StorePackerServiceArticleDatasArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("ArticleDatas", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := StorePackerServiceArticleDatasResult{}
	var retval *ArticleDatasResp
	if retval, err2 = p.handler.ArticleDatas(ctx, args.Request); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing ArticleDatas: "+err2.Error())
		oprot.WriteMessageBegin("ArticleDatas", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("ArticleDatas", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type StorePackerServiceArticleDatasArgs struct {
	Request *ArticleDatasReq `thrift:"request,1" frugal:"1,default,ArticleDatasReq" json:"request"`
}

func NewStorePackerServiceArticleDatasArgs() *StorePackerServiceArticleDatasArgs {
	return &StorePackerServiceArticleDatasArgs{}
}

func (p *StorePackerServiceArticleDatasArgs) InitDefault() {
	*p = StorePackerServiceArticleDatasArgs{}
}

var StorePackerServiceArticleDatasArgs_Request_DEFAULT *ArticleDatasReq

func (p *StorePackerServiceArticleDatasArgs) GetRequest() (v *ArticleDatasReq) {
	if !p.IsSetRequest() {
		return StorePackerServiceArticleDatasArgs_Request_DEFAULT
	}
	return p.Request
}
func (p *StorePackerServiceArticleDatasArgs) SetRequest(val *ArticleDatasReq) {
	p.Request = val
}

var fieldIDToName_StorePackerServiceArticleDatasArgs = map[int16]string{
	1: "request",
}

func (p *StorePackerServiceArticleDatasArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *StorePackerServiceArticleDatasArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_StorePackerServiceArticleDatasArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *StorePackerServiceArticleDatasArgs) ReadField1(iprot thrift.TProtocol) error {
	p.Request = NewArticleDatasReq()
	if err := p.Request.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *StorePackerServiceArticleDatasArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("ArticleDatas_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *StorePackerServiceArticleDatasArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("request", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Request.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *StorePackerServiceArticleDatasArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("StorePackerServiceArticleDatasArgs(%+v)", *p)
}

func (p *StorePackerServiceArticleDatasArgs) DeepEqual(ano *StorePackerServiceArticleDatasArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Request) {
		return false
	}
	return true
}

func (p *StorePackerServiceArticleDatasArgs) Field1DeepEqual(src *ArticleDatasReq) bool {

	if !p.Request.DeepEqual(src) {
		return false
	}
	return true
}

type StorePackerServiceArticleDatasResult struct {
	Success *ArticleDatasResp `thrift:"success,0,optional" frugal:"0,optional,ArticleDatasResp" json:"success,omitempty"`
}

func NewStorePackerServiceArticleDatasResult() *StorePackerServiceArticleDatasResult {
	return &StorePackerServiceArticleDatasResult{}
}

func (p *StorePackerServiceArticleDatasResult) InitDefault() {
	*p = StorePackerServiceArticleDatasResult{}
}

var StorePackerServiceArticleDatasResult_Success_DEFAULT *ArticleDatasResp

func (p *StorePackerServiceArticleDatasResult) GetSuccess() (v *ArticleDatasResp) {
	if !p.IsSetSuccess() {
		return StorePackerServiceArticleDatasResult_Success_DEFAULT
	}
	return p.Success
}
func (p *StorePackerServiceArticleDatasResult) SetSuccess(x interface{}) {
	p.Success = x.(*ArticleDatasResp)
}

var fieldIDToName_StorePackerServiceArticleDatasResult = map[int16]string{
	0: "success",
}

func (p *StorePackerServiceArticleDatasResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *StorePackerServiceArticleDatasResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField0(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_StorePackerServiceArticleDatasResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *StorePackerServiceArticleDatasResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = NewArticleDatasResp()
	if err := p.Success.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *StorePackerServiceArticleDatasResult) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("ArticleDatas_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *StorePackerServiceArticleDatasResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Success.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *StorePackerServiceArticleDatasResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("StorePackerServiceArticleDatasResult(%+v)", *p)
}

func (p *StorePackerServiceArticleDatasResult) DeepEqual(ano *StorePackerServiceArticleDatasResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field0DeepEqual(ano.Success) {
		return false
	}
	return true
}

func (p *StorePackerServiceArticleDatasResult) Field0DeepEqual(src *ArticleDatasResp) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}