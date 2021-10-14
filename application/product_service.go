package application

type ProductService struct {
	Persistence IProductPersistence
}

func (s *ProductService) Get(id string) (IProduct, error) {
	product, err := s.Persistence.Get(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) Create(name string, price float64) (IProduct, error) {
	product := NewProduct()
	product.Name = name
	product.Price = price
	_, err := product.IsValid()

	if err != nil {
		return &Product{}, err
	}

	result, err := s.Persistence.Save(product)

	if err != nil {
		return &Product{}, err
	}

	return result, err
}

func (s *ProductService) Enable(product IProduct) (IProduct, error) {
	err := product.Enable()

	if err != nil {
		return &Product{}, err
	}

	result, err := s.Persistence.Save(product)

	if err != nil {
		return &Product{}, err
	}

	return result, nil
}

func (s *ProductService) Disable(product IProduct) (IProduct, error) {
	err := product.Disable()

	if err != nil {
		return &Product{}, err
	}

	result, err := s.Persistence.Save(product)

	if err != nil {
		return &Product{}, err
	}

	return result, nil
}
