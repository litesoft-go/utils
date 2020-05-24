package times

import "time"

type Passage struct {
	expectedDuration time.Duration
}

//noinspection GoUnusedExportedFunction
func NewPassage(expected time.Duration) *Passage {
	return &Passage{expectedDuration: expected}
}

func (in *Passage) IsSufficientTimeErrIgnore(tPrevious time.Time, err error) bool {
	return in.IsSufficientTimePerrIgnore(&tPrevious, err)
}

func (in *Passage) IsSufficientTimePerrIgnore(tPrevious *time.Time, inErr error) bool {
	sufficient, _ := in.IsSufficientTimePerr(tPrevious, inErr)
	return sufficient
}

func (in *Passage) IsSufficientTimeErr(tPrevious time.Time, err error) (bool, error) {
	return in.IsSufficientTimePerr(&tPrevious, err)
}

func (in *Passage) IsSufficientTimePerr(tPrevious *time.Time, inErr error) (sufficient bool, err error) {
	err = inErr
	if err == nil {
		sufficient = in.IsSufficientTimeP(tPrevious)
	}
	return
}

func (in *Passage) IsSufficientTimeP(tPrevious *time.Time) bool {
	if tPrevious == nil {
		return true
	}
	return in.IsSufficientTime(*tPrevious)
}

func (in *Passage) IsSufficientTime(tPrevious time.Time) bool {
	if in == nil {
		return true
	}
	tNow := time.Now().UTC()
	actualDuration := tNow.Sub(tPrevious)
	return in.expectedDuration <= actualDuration
}
