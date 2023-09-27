// Code generated by Kitex v0.6.1. DO NOT EDIT.

package packer

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/apache/thrift/lib/go/thrift"

	"github.com/cloudwego/kitex/pkg/protocol/bthrift"
	"github.com/tankpanv/demo/store/kitex_gen/wealth/stock/article_base"
	"github.com/tankpanv/demo/store/kitex_gen/wealth/stock/base"
	"github.com/tankpanv/demo/store/kitex_gen/wealth/stock/cell_base"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = thrift.TProtocol(nil)
	_ = bthrift.BinaryWriter(nil)
	_ = article_base.KitexUnusedProtection
	_ = base.KitexUnusedProtection
	_ = cell_base.KitexUnusedProtection
)

func (p *ArticleDatasReq) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I16 {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.I32 {
				l, err = p.FastReadField2(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.I16 {
				l, err = p.FastReadField3(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 4:
			if fieldTypeId == thrift.I32 {
				l, err = p.FastReadField4(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 5:
			if fieldTypeId == thrift.LIST {
				l, err = p.FastReadField5(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ArticleDatasReq[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ArticleDatasReq) FastReadField1(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI16(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		p.AppId = &v

	}
	return offset, nil
}

func (p *ArticleDatasReq) FastReadField2(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI32(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		tmp := base.Source(v)
		p.Source = &tmp

	}
	return offset, nil
}

func (p *ArticleDatasReq) FastReadField3(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI16(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		p.VersionCode = &v

	}
	return offset, nil
}

func (p *ArticleDatasReq) FastReadField4(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI32(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		tmp := cell_base.PackerType(v)
		p.PackerType = &tmp

	}
	return offset, nil
}

func (p *ArticleDatasReq) FastReadField5(buf []byte) (int, error) {
	offset := 0

	_, size, l, err := bthrift.Binary.ReadListBegin(buf[offset:])
	offset += l
	if err != nil {
		return offset, err
	}
	p.Keywords = make([]string, 0, size)
	for i := 0; i < size; i++ {
		var _elem string
		if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
			return offset, err
		} else {
			offset += l

			_elem = v

		}

		p.Keywords = append(p.Keywords, _elem)
	}
	if l, err := bthrift.Binary.ReadListEnd(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	return offset, nil
}

// for compatibility
func (p *ArticleDatasReq) FastWrite(buf []byte) int {
	return 0
}

func (p *ArticleDatasReq) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "ArticleDatasReq")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
		offset += p.fastWriteField3(buf[offset:], binaryWriter)
		offset += p.fastWriteField2(buf[offset:], binaryWriter)
		offset += p.fastWriteField4(buf[offset:], binaryWriter)
		offset += p.fastWriteField5(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *ArticleDatasReq) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("ArticleDatasReq")
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
		l += p.field3Length()
		l += p.field4Length()
		l += p.field5Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *ArticleDatasReq) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetAppId() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "app_id", thrift.I16, 1)
		offset += bthrift.Binary.WriteI16(buf[offset:], *p.AppId)

		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *ArticleDatasReq) fastWriteField2(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetSource() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "source", thrift.I32, 2)
		offset += bthrift.Binary.WriteI32(buf[offset:], int32(*p.Source))

		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *ArticleDatasReq) fastWriteField3(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetVersionCode() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "version_code", thrift.I16, 3)
		offset += bthrift.Binary.WriteI16(buf[offset:], *p.VersionCode)

		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *ArticleDatasReq) fastWriteField4(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetPackerType() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "packer_type", thrift.I32, 4)
		offset += bthrift.Binary.WriteI32(buf[offset:], int32(*p.PackerType))

		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *ArticleDatasReq) fastWriteField5(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetKeywords() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "keywords", thrift.LIST, 5)
		listBeginOffset := offset
		offset += bthrift.Binary.ListBeginLength(thrift.STRING, 0)
		var length int
		for _, v := range p.Keywords {
			length++
			offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, v)

		}
		bthrift.Binary.WriteListBegin(buf[listBeginOffset:], thrift.STRING, length)
		offset += bthrift.Binary.WriteListEnd(buf[offset:])
		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *ArticleDatasReq) field1Length() int {
	l := 0
	if p.IsSetAppId() {
		l += bthrift.Binary.FieldBeginLength("app_id", thrift.I16, 1)
		l += bthrift.Binary.I16Length(*p.AppId)

		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *ArticleDatasReq) field2Length() int {
	l := 0
	if p.IsSetSource() {
		l += bthrift.Binary.FieldBeginLength("source", thrift.I32, 2)
		l += bthrift.Binary.I32Length(int32(*p.Source))

		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *ArticleDatasReq) field3Length() int {
	l := 0
	if p.IsSetVersionCode() {
		l += bthrift.Binary.FieldBeginLength("version_code", thrift.I16, 3)
		l += bthrift.Binary.I16Length(*p.VersionCode)

		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *ArticleDatasReq) field4Length() int {
	l := 0
	if p.IsSetPackerType() {
		l += bthrift.Binary.FieldBeginLength("packer_type", thrift.I32, 4)
		l += bthrift.Binary.I32Length(int32(*p.PackerType))

		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *ArticleDatasReq) field5Length() int {
	l := 0
	if p.IsSetKeywords() {
		l += bthrift.Binary.FieldBeginLength("keywords", thrift.LIST, 5)
		l += bthrift.Binary.ListBeginLength(thrift.STRING, len(p.Keywords))
		for _, v := range p.Keywords {
			l += bthrift.Binary.StringLengthNocopy(v)

		}
		l += bthrift.Binary.ListEndLength()
		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *ArticleDatasResp) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.LIST {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ArticleDatasResp[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ArticleDatasResp) FastReadField1(buf []byte) (int, error) {
	offset := 0

	_, size, l, err := bthrift.Binary.ReadListBegin(buf[offset:])
	offset += l
	if err != nil {
		return offset, err
	}
	p.ArticleDatas = make([]*article_base.ArticleData, 0, size)
	for i := 0; i < size; i++ {
		_elem := article_base.NewArticleData()
		if l, err := _elem.FastRead(buf[offset:]); err != nil {
			return offset, err
		} else {
			offset += l
		}

		p.ArticleDatas = append(p.ArticleDatas, _elem)
	}
	if l, err := bthrift.Binary.ReadListEnd(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	return offset, nil
}

// for compatibility
func (p *ArticleDatasResp) FastWrite(buf []byte) int {
	return 0
}

func (p *ArticleDatasResp) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "ArticleDatasResp")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *ArticleDatasResp) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("ArticleDatasResp")
	if p != nil {
		l += p.field1Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *ArticleDatasResp) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetArticleDatas() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "ArticleDatas", thrift.LIST, 1)
		listBeginOffset := offset
		offset += bthrift.Binary.ListBeginLength(thrift.STRUCT, 0)
		var length int
		for _, v := range p.ArticleDatas {
			length++
			offset += v.FastWriteNocopy(buf[offset:], binaryWriter)
		}
		bthrift.Binary.WriteListBegin(buf[listBeginOffset:], thrift.STRUCT, length)
		offset += bthrift.Binary.WriteListEnd(buf[offset:])
		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *ArticleDatasResp) field1Length() int {
	l := 0
	if p.IsSetArticleDatas() {
		l += bthrift.Binary.FieldBeginLength("ArticleDatas", thrift.LIST, 1)
		l += bthrift.Binary.ListBeginLength(thrift.STRUCT, len(p.ArticleDatas))
		for _, v := range p.ArticleDatas {
			l += v.BLength()
		}
		l += bthrift.Binary.ListEndLength()
		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *StorePackerServiceArticleDatasArgs) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_StorePackerServiceArticleDatasArgs[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *StorePackerServiceArticleDatasArgs) FastReadField1(buf []byte) (int, error) {
	offset := 0

	tmp := NewArticleDatasReq()
	if l, err := tmp.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.Request = tmp
	return offset, nil
}

// for compatibility
func (p *StorePackerServiceArticleDatasArgs) FastWrite(buf []byte) int {
	return 0
}

func (p *StorePackerServiceArticleDatasArgs) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "ArticleDatas_args")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *StorePackerServiceArticleDatasArgs) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("ArticleDatas_args")
	if p != nil {
		l += p.field1Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *StorePackerServiceArticleDatasArgs) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "request", thrift.STRUCT, 1)
	offset += p.Request.FastWriteNocopy(buf[offset:], binaryWriter)
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *StorePackerServiceArticleDatasArgs) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("request", thrift.STRUCT, 1)
	l += p.Request.BLength()
	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *StorePackerServiceArticleDatasResult) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				l, err = p.FastReadField0(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_StorePackerServiceArticleDatasResult[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *StorePackerServiceArticleDatasResult) FastReadField0(buf []byte) (int, error) {
	offset := 0

	tmp := NewArticleDatasResp()
	if l, err := tmp.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.Success = tmp
	return offset, nil
}

// for compatibility
func (p *StorePackerServiceArticleDatasResult) FastWrite(buf []byte) int {
	return 0
}

func (p *StorePackerServiceArticleDatasResult) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "ArticleDatas_result")
	if p != nil {
		offset += p.fastWriteField0(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *StorePackerServiceArticleDatasResult) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("ArticleDatas_result")
	if p != nil {
		l += p.field0Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *StorePackerServiceArticleDatasResult) fastWriteField0(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetSuccess() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "success", thrift.STRUCT, 0)
		offset += p.Success.FastWriteNocopy(buf[offset:], binaryWriter)
		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *StorePackerServiceArticleDatasResult) field0Length() int {
	l := 0
	if p.IsSetSuccess() {
		l += bthrift.Binary.FieldBeginLength("success", thrift.STRUCT, 0)
		l += p.Success.BLength()
		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *StorePackerServiceArticleDatasArgs) GetFirstArgument() interface{} {
	return p.Request
}

func (p *StorePackerServiceArticleDatasResult) GetResult() interface{} {
	return p.Success
}