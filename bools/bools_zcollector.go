package bools

import "fmt"

type OrCollector struct {
	value bool
}

//noinspection GoUnusedExportedFunction
func New(values ...bool) *OrCollector {
	for _, b := range values {
		if b {
			return &OrCollector{value: true}
		}
	}
	return &OrCollector{}
}

func (in *OrCollector) Get() bool {
	return in.value
}

func (in *OrCollector) OrWithErrPrinted(value bool, err error) {
	in.value = OrWithErrPrinted(in.value, value, err)
}

func (in *OrCollector) OrWithErrIgnored(value bool, err error) {
	in.value = OrWithErrIgnored(in.value, value, err)
}

func (in *OrCollector) OrWithErr(value bool, err error) error {
	in.Or(value)
	return err
}

func (in *OrCollector) Or(value bool) {
	in.value = Or(in.value, value)
}

func (in *OrCollector) True() {
	in.value = true
}

func OrWithErrPrinted(b1, b2 bool, err error) bool {
	fmt.Println("********************* ", err)
	return OrWithErrIgnored(b1, b2, err)
}

func OrWithErrIgnored(b1, b2 bool, err error) (ored bool) {
	ored, _ = OrWithErr(b1, b2, err)
	return
}

func OrWithErr(b1, b2 bool, err error) (bool, error) {
	return Or(b1, b2), err
}

func Or(b1, b2 bool) bool {
	return b1 || b2
}
