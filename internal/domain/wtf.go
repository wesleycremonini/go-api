package domain

import "time"

type WTF struct {
	ID        int       `json:"id"`
	Wtf       string    `json:"wtf"`
	CreatedAt time.Time `json:"created_at"`
}

func (w *WTF) IsValid() error {
	if w.Wtf == "" {
		return Errorf(ErrINVALID, "Wtf required.")
	}

	return nil
}

type WtfService interface {
	WtfByID(id int) (*WTF, error)
	Wtfs(filter WtfFilter) ([]*WTF, int, error)
	CreateWtf(wtf *WTF) error
	UpdateWtf(id int, wtf *WtfUpdate) (*WTF, error)
	DeleteWtf(id int) error
}

type WtfFilter struct {
	IDs  []int    `json:"ids"`
	Wtfs []string `json:"wtfs"`

	Filter
}

type WtfUpdate struct {
	Wtf *string `json:"wtf"`
}
