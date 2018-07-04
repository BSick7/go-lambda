package metric

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
)

type Point interface {
	SetResolution(resolution int64) Point
	ToAWS() cloudwatch.MetricDatum
}

type PointMarker interface {
	Point
	Mark()
}
