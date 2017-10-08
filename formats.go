package rtfdoc

import "errors"

func getSize(format string, layout string) (size, error) {
	switch format {
	case FormatA5:
		switch layout {
		case OrientationLandscape:
			return size{
				width:  11952,
				height: 16848 / 2,
			}, nil
		case OrientationPortrait:
			return size{
				width:  16848 / 2,
				height: 11952,
			}, nil
		default:
			return size{}, errors.New("Incorrect document orientation")

		}
	case FormatA4:
		switch layout {
		case OrientationLandscape:
			return size{
				width:  16848,
				height: 11952,
			}, nil
		case OrientationPortrait:
			return size{
				width:  11952,
				height: 16848,
			}, nil
		default:
			return size{}, errors.New("Incorrect document orientation")

		}
	case FormatA3:
		switch layout {
		case OrientationLandscape:
			return size{
				width:  11952 * 2,
				height: 16848,
			}, nil
		case OrientationPortrait:
			return size{
				width:  16848,
				height: 11952 * 2,
			}, nil
		default:
			return size{}, errors.New("Incorrect document orientation")

		}
	case FormatA2:
		switch layout {
		case OrientationLandscape:
			return size{
				width:  16848 * 2,
				height: 11952 * 2,
			}, nil
		case OrientationPortrait:
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
