package metric

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
)

type Point interface {
	ToAWS() cloudwatch.MetricDatum
}
