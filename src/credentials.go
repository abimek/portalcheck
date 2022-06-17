package portalcheck

type Identifier int

type Credentials struct {
	Identifier Identifier
	Password   string
}
