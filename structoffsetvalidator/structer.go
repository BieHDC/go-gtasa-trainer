package structoffsetvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

func getFieldTagValue[Type any](structoid Type, fieldname, tagname string) (offset string, ret error) {
	// The whole thing is kinda wonky because some reflect calls can panic without
	// a way to cleanly get an error returned
	defer func() {
		convertRecoverToError := func(r interface{}) error {
			switch x := r.(type) {
			case string:
				return errors.New(x)
			case error:
				return x
			default:
				return errors.New(fmt.Sprint(x))
			}
		}
		if r := recover(); r != nil {
			// Name returns are only used here
			ret = convertRecoverToError(r)
			offset = ""
		}
	}()
	if reflect.TypeOf(structoid).Kind() != reflect.Pointer {
		panic("you must pass a pointer to an allocated struct")
	}

	structinfo := reflect.ValueOf(structoid).Type().Elem()
	if structinfo == nil {
		return "", errors.New("failed to get type of")
	}
	//log.Println("type is: ", typeoffed.Name())

	fieldvalue, ok := structinfo.FieldByName(fieldname)
	if !ok {
		return "", errors.New("field not found: " + fieldname)
	}
	if fieldvalue.Name == "" {
		return "", errors.New("field value bad")
	}
	//log.Println("good field value: " + fieldvalue.Name)

	tagval := fieldvalue.Tag.Get(tagname) //this is specific to my struct tags
	if tagval == "" {
		return "", errors.New("bad tag: " + tagval)
	}
	//log.Println("has offset:" + tagval)
	return tagval, nil
}

// Parses the "offset" field and checks if it matches with the compilers
// Generated offset. put this into an init() func for each struct to
// validate it.
func ValidateOffsets[Type any](structoid *Type) error {
	if reflect.TypeOf(structoid).Kind() != reflect.Pointer {
		panic("you must pass a pointer to an allocated struct")
	}

	var faildump strings.Builder
	hasfailed := false

	structinfo := reflect.ValueOf(structoid).Elem()

	// base address of our struct
	address := uintptr(unsafe.Pointer(structoid))
	faildump.WriteString(fmt.Sprintf("Base address of struct (%s) 0x%x\n", structinfo.Type().Name(), address))

	fields := reflect.VisibleFields(structinfo.Type())
	for _, field := range fields {
		faildump.WriteString(fmt.Sprintf("Name: %s\n", field.Name))
		faildump.WriteString(fmt.Sprintf("\tType: %s\n", field.Type.String()))
		//faildump.WriteString(fmt.Sprintf("\tTag: %s\n", field.Tag))

		tagval, err := getFieldTagValue(structoid, field.Name, "offset")
		if err == nil {
			offset, err := strconv.ParseUint(tagval, 0, 32)

			if err != nil {
				faildump.WriteString("\t!!! failed to get a custom offset for " + field.Name + " !!!\n")
				hasfailed = true
				continue
			}

			faildump.WriteString(fmt.Sprintf("\toffset is: 0x%x\n", offset))
			calculatedaddr := address + uintptr(offset)

			ptroffield := structinfo.FieldByName(field.Name).UnsafeAddr()
			diff := int(calculatedaddr) - int(ptroffield)

			if calculatedaddr != ptroffield {
				faildump.WriteString(fmt.Sprintf("\texpecting result to be 0x%x\n", calculatedaddr))
				faildump.WriteString(fmt.Sprintf("\tactual address is 0x%x\n", ptroffield))
				faildump.WriteString(fmt.Sprintf("\tdifference: %d || offset should be 0x%x if correct\n", diff, int(offset)+diff))
				faildump.WriteString(fmt.Sprintf("\t!!! ERROR for field %s !!!\n", field.Name))
				hasfailed = true
				continue
			}

			continue
		} else {
			// if there is no offset, we assume padding
			faildump.WriteString(fmt.Sprintf("\tIgnoring...\n"))
		}

	}

	if hasfailed {
		return errors.New(faildump.String())
	}
	return nil
}
