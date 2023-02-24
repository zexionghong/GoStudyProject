package class_path

type IncFactory interface {
	Create(version int) Phone
}

type AndroidFactory struct {
}

func (f *AndroidFactory) Create(version int) Phone {
	var phone Phone
	switch version {
	default:
		phone = HuaWei{

			Android: &Android{},
		}

	}
	return phone
}

type IphoneFactory struct {
}

func (f *IphoneFactory) Create(version int) Phone {
	var phone Phone
	switch version {
	default:
		phone = iphone14{
			size:   14,
			Iphone: &Iphone{},
		}
	case 14:
		phone = iphone14{
			size:   14,
			Iphone: &Iphone{},
		}
	case 12:
		phone = iphone12{
			size:   12,
			Iphone: &Iphone{}}

	}
	return phone
}
