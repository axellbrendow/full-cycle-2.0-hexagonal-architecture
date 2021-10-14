package application

type IProduct interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetId() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type Product struct {
	Id     string
	Name   string
	Status string
	Price  float64
}

func (p *Product) IsValid() (bool, error) {
	//
}

func (p *Product) Enable() error {
	//
}

func (p *Product) Disable() error {
	//
}

func (p *Product) GetId() string {
	//
}

func (p *Product) GetName() string {
	//
}

func (p *Product) GetStatus() string {
	//
}

func (p *Product) GetPrice() string {
	//
}
