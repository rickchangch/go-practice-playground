package structConvert

import (
	"fmt"
	"reflect"
	"time"

	"github.com/jinzhu/copier"
)

type structConvert struct{}

var Instance = new(structConvert)

func (s *structConvert) Run() {
	dataA := &StructA{
		ID:      "1aeqwqe",
		Name:    "eeee",
		Created: time.Now(),
	}
	dataB := &StructB{}
	structConvByCopier(dataB, dataA)
	fmt.Printf("- use copier lib: %v\n", dataB)
	structConvByReflect(dataB, dataA)
	fmt.Printf("- use reflect lib: %v\n", dataB)
}

// convert struct to another struct
type StructA struct {
	ID      string
	Name    string
	Created time.Time
}
type StructB struct {
	ID      string
	Name    string
	Created time.Time
}

// use `copier` to implement struct conversion
func structConvByCopier(targetEntity, referEntity interface{}) {
	copier.Copy(targetEntity, referEntity)
}

// use `reflect` to implement struct conversion
func structConvByReflect(targetEntity, referEntity interface{}) error {

	// reflect value of target entity
	refValTarget := reflect.ValueOf(targetEntity)
	if refValTarget.Kind() == reflect.Ptr {
		refValTarget = refValTarget.Elem()
	}
	refTypeTarget := refValTarget.Type()

	// reflect value of reference entity
	refValRefer := reflect.ValueOf(referEntity)
	if refValRefer.Kind() == reflect.Ptr {
		refValRefer = refValRefer.Elem()
	}
	refTypeRefer := refValRefer.Type()

	for i := 0; i < refValRefer.NumField(); i++ {

		// get field name of reference entity
		fieldName := refTypeRefer.Field(i).Name

		// get reflect of the field of target entity
		refFieldValTarget := refValTarget.FieldByName(fieldName)
		refFieldTypeTarget, exist := refTypeTarget.FieldByName(fieldName)

		// check if the field value of target entity is valid (indeed exist)
		if !exist || !refFieldValTarget.IsValid() {
			return fmt.Errorf("target field value do not exist")
		}

		// check if the field value of target entity can be changed
		if !refFieldValTarget.CanSet() {
			return fmt.Errorf("target value is unchangeable")
		}

		// check if the types of two fileds are the same
		if refFieldTypeTarget.Name != refTypeRefer.Field(i).Name {
			return fmt.Errorf("two fields not match")
		}

		// copy value from reference entity
		refFieldValTarget.Set(refValRefer.Field(i))
	}

	return nil
}
