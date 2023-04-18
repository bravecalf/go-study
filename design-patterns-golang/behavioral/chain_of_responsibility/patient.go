package chain_of_responsibility

type Patient struct {
	name             string
	RegistrationDone bool
	DoctorCheckDone  bool
	MedicineDone     bool
	PaymentDone      bool
}
