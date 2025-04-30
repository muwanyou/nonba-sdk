package enum

type Sex string

const (
	SexMale   Sex = "male"
	SexFemale Sex = "female"
)

const DefaultSex = SexMale

func (s Sex) String() string {
	return string(s)
}
