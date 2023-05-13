package encrypt

import "github.com/google/uuid"

func UUIDv1() (id uuid.UUID, err error) {
	id, err = uuid.NewUUID()
	return
}

func UUIDv2G() (id uuid.UUID, err error) {
	id, err = uuid.NewDCEGroup()
	return
}

func UUIDv2P() (id uuid.UUID, err error) {
	id, err = uuid.NewDCEPerson()
	return
}

func UUIDv3(data []byte) (id uuid.UUID, err error) {
	id, err = uuid.NewDCEPerson()
	if err != nil {
		return uuid.UUID{}, err
	}
	return uuid.NewMD5(id, data), err
}

func UUIDv4() uuid.UUID {
	return uuid.New()
}

func UUIDv5(data []byte) (id uuid.UUID, err error) {
	id, err = uuid.NewDCEPerson()
	if err != nil {
		return uuid.UUID{}, err
	}
	return uuid.NewSHA1(id, data), err
}
