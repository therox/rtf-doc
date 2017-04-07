package rtfdoc

import "errors"

func getSize(format string, layout string) (Size, error) {
	switch format {
	case "A4", "a4":
		switch layout {
		case "landscape":
			return Size{
				width:  16848,
				height: 11952,
			}, nil
		case "portrait":
			return Size{
				width:  11952,
				height: 16848,
			}, nil
		default:
			return Size{}, errors.New("Incorrect document orientation")

		}
	case "A3", "a3":
		switch layout {
		case "landscape":
			return Size{
				width:  11952 * 2,
				height: 16848,
			}, nil
		case "portrait":
			return Size{
				width:  16848,
				height: 11952 * 2,
			}, nil
		default:
			return Size{}, errors.New("Incorrect document orientation")

		}
	case "A2", "a2":
		switch layout {
		case "landscape":
			return Size{
				width:  16848 * 2,
				height: 11952 * 2,
			}, nil
		case "portrait":
			return Size{
				width:  11952 * 2,
				height: 16848 * 2,
			}, nil
		default:
			return Size{}, errors.New("Incorrect document orientation")

		}
	default:
		return Size{}, errors.New("Incorrect document format")
	}

}
