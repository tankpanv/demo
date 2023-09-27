// Code generated by thriftgo (0.2.12). DO NOT EDIT.

package cell_base

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/tankpanv/demo/store/kitex_gen/wealth/stock/article_base"
)

type ShowType int64

const (
	ShowType_ArticleDetail ShowType = 1
)

func (p ShowType) String() string {
	switch p {
	case ShowType_ArticleDetail:
		return "ArticleDetail"
	}
	return "<UNSET>"
}

func ShowTypeFromString(s string) (ShowType, error) {
	switch s {
	case "ArticleDetail":
		return ShowType_ArticleDetail, nil
	}
	return ShowType(0), fmt.Errorf("not a valid ShowType string")
}

func ShowTypePtr(v ShowType) *ShowType { return &v }
func (p *ShowType) Scan(value interface{}) (err error) {
	var result sql.NullInt64
	err = result.Scan(value)
	*p = ShowType(result.Int64)
	return
}

func (p *ShowType) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

type HanderType int64

const (
	HanderType_ArticleDetail HanderType = 1
)

func (p HanderType) String() string {
	switch p {
	case HanderType_ArticleDetail:
		return "ArticleDetail"
	}
	return "<UNSET>"
}

func HanderTypeFromString(s string) (HanderType, error) {
	switch s {
	case "ArticleDetail":
		return HanderType_ArticleDetail, nil
	}
	return HanderType(0), fmt.Errorf("not a valid HanderType string")
}

func HanderTypePtr(v HanderType) *HanderType { return &v }
func (p *HanderType) Scan(value interface{}) (err error) {
	var result sql.NullInt64
	err = result.Scan(value)
	*p = HanderType(result.Int64)
	return
}

func (p *HanderType) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

type PackerType int64

const (
	PackerType_Recommend PackerType = 0
)

func (p PackerType) String() string {
	switch p {
	case PackerType_Recommend:
		return "Recommend"
	}
	return "<UNSET>"
}

func PackerTypeFromString(s string) (PackerType, error) {
	switch s {
	case "Recommend":
		return PackerType_Recommend, nil
	}
	return PackerType(0), fmt.Errorf("not a valid PackerType string")
}

func PackerTypePtr(v PackerType) *PackerType { return &v }
func (p *PackerType) Scan(value interface{}) (err error) {
	var result sql.NullInt64
	err = result.Scan(value)
	*p = PackerType(result.Int64)
	return
}

func (p *PackerType) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

type MainTabType int64

const (
	MainTabType_RecommendTab MainTabType = 0
)

func (p MainTabType) String() string {
	switch p {
	case MainTabType_RecommendTab:
		return "RecommendTab"
	}
	return "<UNSET>"
}

func MainTabTypeFromString(s string) (MainTabType, error) {
	switch s {
	case "RecommendTab":
		return MainTabType_RecommendTab, nil
	}
	return MainTabType(0), fmt.Errorf("not a valid MainTabType string")
}

func MainTabTypePtr(v MainTabType) *MainTabType { return &v }
func (p *MainTabType) Scan(value interface{}) (err error) {
	var result sql.NullInt64
	err = result.Scan(value)
	*p = MainTabType(result.Int64)
	return
}

func (p *MainTabType) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

type CellView struct {
	ShowType    *ShowType                 `thrift:"show_type,1,optional" frugal:"1,optional,ShowType" json:"show_type,omitempty"`
	HandlerType *HanderType               `thrift:"handler_type,2,optional" frugal:"2,optional,HanderType" json:"handler_type,omitempty"`
	ArticleData *article_base.ArticleData `thrift:"article_data,3,optional" frugal:"3,optional,article_base.ArticleData" json:"article_data,omitempty"`
	SubCellView []*CellView               `thrift:"sub_cell_view,4,optional" frugal:"4,optional,list<CellView>" json:"sub_cell_view,omitempty"`
}

func NewCellView() *CellView {
	return &CellView{}
}

func (p *CellView) InitDefault() {
	*p = CellView{}
}

var CellView_ShowType_DEFAULT ShowType

func (p *CellView) GetShowType() (v ShowType) {
	if !p.IsSetShowType() {
		return CellView_ShowType_DEFAULT
	}
	return *p.ShowType
}

var CellView_HandlerType_DEFAULT HanderType

func (p *CellView) GetHandlerType() (v HanderType) {
	if !p.IsSetHandlerType() {
		return CellView_HandlerType_DEFAULT
	}
	return *p.HandlerType
}

var CellView_ArticleData_DEFAULT *article_base.ArticleData

func (p *CellView) GetArticleData() (v *article_base.ArticleData) {
	if !p.IsSetArticleData() {
		return CellView_ArticleData_DEFAULT
	}
	return p.ArticleData
}

var CellView_SubCellView_DEFAULT []*CellView

func (p *CellView) GetSubCellView() (v []*CellView) {
	if !p.IsSetSubCellView() {
		return CellView_SubCellView_DEFAULT
	}
	return p.SubCellView
}
func (p *CellView) SetShowType(val *ShowType) {
	p.ShowType = val
}
func (p *CellView) SetHandlerType(val *HanderType) {
	p.HandlerType = val
}
func (p *CellView) SetArticleData(val *article_base.ArticleData) {
	p.ArticleData = val
}
func (p *CellView) SetSubCellView(val []*CellView) {
	p.SubCellView = val
}

var fieldIDToName_CellView = map[int16]string{
	1: "show_type",
	2: "handler_type",
	3: "article_data",
	4: "sub_cell_view",
}

func (p *CellView) IsSetShowType() bool {
	return p.ShowType != nil
}

func (p *CellView) IsSetHandlerType() bool {
	return p.HandlerType != nil
}

func (p *CellView) IsSetArticleData() bool {
	return p.ArticleData != nil
}

func (p *CellView) IsSetSubCellView() bool {
	return p.SubCellView != nil
}

func (p *CellView) Read(iprot thrift.TProtocol) (err error) {

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
			if fieldTypeId == thrift.I32 {
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
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 4:
			if fieldTypeId == thrift.LIST {
				if err = p.ReadField4(iprot); err != nil {
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
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_CellView[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *CellView) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return err
	} else {
		tmp := ShowType(v)
		p.ShowType = &tmp
	}
	return nil
}

func (p *CellView) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return err
	} else {
		tmp := HanderType(v)
		p.HandlerType = &tmp
	}
	return nil
}

func (p *CellView) ReadField3(iprot thrift.TProtocol) error {
	p.ArticleData = article_base.NewArticleData()
	if err := p.ArticleData.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *CellView) ReadField4(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return err
	}
	p.SubCellView = make([]*CellView, 0, size)
	for i := 0; i < size; i++ {
		_elem := NewCellView()
		if err := _elem.Read(iprot); err != nil {
			return err
		}

		p.SubCellView = append(p.SubCellView, _elem)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return err
	}
	return nil
}

func (p *CellView) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("CellView"); err != nil {
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

func (p *CellView) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetShowType() {
		if err = oprot.WriteFieldBegin("show_type", thrift.I32, 1); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteI32(int32(*p.ShowType)); err != nil {
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

func (p *CellView) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetHandlerType() {
		if err = oprot.WriteFieldBegin("handler_type", thrift.I32, 2); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteI32(int32(*p.HandlerType)); err != nil {
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

func (p *CellView) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetArticleData() {
		if err = oprot.WriteFieldBegin("article_data", thrift.STRUCT, 3); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.ArticleData.Write(oprot); err != nil {
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

func (p *CellView) writeField4(oprot thrift.TProtocol) (err error) {
	if p.IsSetSubCellView() {
		if err = oprot.WriteFieldBegin("sub_cell_view", thrift.LIST, 4); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteListBegin(thrift.STRUCT, len(p.SubCellView)); err != nil {
			return err
		}
		for _, v := range p.SubCellView {
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
	return thrift.PrependError(fmt.Sprintf("%T write field 4 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 end error: ", p), err)
}

func (p *CellView) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CellView(%+v)", *p)
}

func (p *CellView) DeepEqual(ano *CellView) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.ShowType) {
		return false
	}
	if !p.Field2DeepEqual(ano.HandlerType) {
		return false
	}
	if !p.Field3DeepEqual(ano.ArticleData) {
		return false
	}
	if !p.Field4DeepEqual(ano.SubCellView) {
		return false
	}
	return true
}

func (p *CellView) Field1DeepEqual(src *ShowType) bool {

	if p.ShowType == src {
		return true
	} else if p.ShowType == nil || src == nil {
		return false
	}
	if *p.ShowType != *src {
		return false
	}
	return true
}
func (p *CellView) Field2DeepEqual(src *HanderType) bool {

	if p.HandlerType == src {
		return true
	} else if p.HandlerType == nil || src == nil {
		return false
	}
	if *p.HandlerType != *src {
		return false
	}
	return true
}
func (p *CellView) Field3DeepEqual(src *article_base.ArticleData) bool {

	if !p.ArticleData.DeepEqual(src) {
		return false
	}
	return true
}
func (p *CellView) Field4DeepEqual(src []*CellView) bool {

	if len(p.SubCellView) != len(src) {
		return false
	}
	for i, v := range p.SubCellView {
		_src := src[i]
		if !v.DeepEqual(_src) {
			return false
		}
	}
	return true
}
