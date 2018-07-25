package metric

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
)

// NewCloudwatchEmitter creates an emitter that publishes to AWS Cloudwatch Metrics
func NewCloudwatchEmitter(namespace string) Emitter {
	return &cloudwatchEmitter{namespace: namespace}
}

type cloudwatchEmitter struct {
	namespace string
	points    []*Point
}

func (e *cloudwatchEmitter) Contextualize(ctx context.Context) context.Context {
	return WithEmitter(ctx, e)
}

func (e *cloudwatchEmitter) Add(point *Point) {
	e.points = append(e.points, point)
}

func (e *cloudwatchEmitter) Flush() error {
	points := make([]cloudwatch.MetricDatum, 0)
	for _, point := range e.points {
		points = append(points, toAWSDatum(point))
	}
	return PutCloudwatch(e.namespace, points)
}

func toAWSDatum(point *Point) cloudwatch.MetricDatum {
	dims := make([]cloudwatch.Dimension, 0)
	for name, value := range point.Tags {
		dims = append(dims, cloudwatch.Dimension{
			Name:  aws.String(name),
			Value: aws.String(value),
		})
	}

	return cloudwatch.MetricDatum{
		MetricName:        aws.String(point.Metric),
		Unit:              cloudwatch.StandardUnit(point.Unit),
		Timestamp:         aws.Time(point.Timestamp),
		StorageResolution: point.Resolution,
		Value:             aws.Float64(point.Value),
		Dimensions:        dims,
	}
}
