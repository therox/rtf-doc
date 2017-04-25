package rtfdoc

import "errors"

func getSize(format string, layout string) (size, error) {
	switch format {
	case "A4", "a4":
		switch layout {
		case "landscape":
			return size{
				width:  16848,
				height: 11952,
			}, nil
		case "portrait":
			return size{
				width:  11952,
				height: 16848,
			}, nil
		default:
			return size{}, errors.New("Incorrect document orientation")

		}
	case "A3", "a3":
		switch layout {
		case "landscape":
			return size{
				width:  11952 * 2,
				height: 16848,
			}, nil
		case "portrait":
			return size{
				width:  16848,
				height: 11952 * 2,
			}, nil
		default:
			return size{}, errors.New("Incorrect document orientation")

		}
	case "A2", "a2":
		switch layout {
		case "landscape":
			return size{
				width:  16848 * 2,
				height: 11952 * 2,
			}, nil
		case "portrait":
			return size{
				width:  11952 * 2,
				height: 16848 * 2,
			}, nil
		default:
			return size{}, errors.New("Incorrect document orientation")

		}
	default:
		return size{}, errors.New("Incorrect document format")
	}

}
