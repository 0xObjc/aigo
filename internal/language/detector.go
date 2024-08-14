package language

type Detector interface {
	DetectLanguage(dir string) string
}

var detectors []Detector

func RegisterDetector(detector Detector) {
	detectors = append(detectors, detector)
}

func Detectors() []Detector {
	return detectors
}
