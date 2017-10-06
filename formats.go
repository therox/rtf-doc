package rtfdoc

import "errors"

func getSize(format string, layout string) (size, error) {
	switch format {
	case FORMAT_A5:
		switch layout {
		case ORIENTATION_LANDSCAPE:
			return size{
				width:  11952,
				height: 16848 / 2,
			}, nil
		case ORIENTATION_PORTRAIT:
			return size{
				width:  16848 / 2,
				height: 11952,
			}, nil
		default:
			return size{}, errors.New("Incorrect document orientation")

		}
	case FORMAT_A4:
		switch layout {
		case ORIENTATION_LANDSCAPE:
			return size{
				width:  16848,
				height: 11952,
			}, nil
		case ORIENTATION_PORTRAIT:
			return size{
				width:  11952,
				height: 16848,
			}, nil
		default:
			return size{}, errors.New("Incorrect document orientation")

		}
	case FORMAT_A3:
		switch layout {
		case ORIENTATION_LANDSCAPE:
			return size{
				width:  11952 * 2,
				height: 16848,
			}, nil
		case ORIENTATION_PORTRAIT:
			return size{
				width:  16848,
				height: 11952 * 2,
			}, nil
		default:
			return size{}, errors.New("Incorrect document orientation")

		}
	case FORMAT_A2:
		switch layout {
		case ORIENTATION_LANDSCAPE:
			return size{
				width:  16848 * 2,
				height: 11952 * 2,
			}, nil
		case ORIENTATION_PORTRAIT:
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
