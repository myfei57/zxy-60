package flight

import "bagsort/internal/model"

func (b *Book) FindByCode(code string) (model.Flight, bool, error) {
	flights, err := b.List()
	if err != nil {
		return model.Flight{}, false, err
	}
	for _, f := range flights {
		if f.Code == code {
			return f, true, nil
		}
	}
	return model.Flight{}, false, nil
}

func (b *Book) OpenFlights() ([]model.Flight, error) {
	flights, err := b.List()
	if err != nil {
		return nil, err
	}
	out := make([]model.Flight, 0)
	for _, f := range flights {
		if model.IsOpenFlightState(f.State) {
			out = append(out, f)
		}
	}
	return out, nil
}
