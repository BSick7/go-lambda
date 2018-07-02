package metric

import (
	"fmt"

	"github.com/BSick7/go-lambda/services"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
)

var (
	maxMetricsPerPut = 20
)

func PutCloudwatch(namespace string, values []Point) error {
	svc, err := services.NewCloudwatch()
	if err != nil {
		return fmt.Errorf("could not create cloudwatch service: %s", err)
	}

	batch := make([]cloudwatch.MetricDatum, 0)
	for _, item := range values {
		if len(batch) >= maxMetricsPerPut {
			req := svc.PutMetricDataRequest(&cloudwatch.PutMetricDataInput{
				Namespace:  aws.String(namespace),
				MetricData: batch,
			})
			if _, err := req.Send(); err != nil {
				return err
			}
			batch = []cloudwatch.MetricDatum{item.ToAWS()}
		} else {
			batch = append(batch, item.ToAWS())
		}
	}
	if len(batch) > 0 {
		req := svc.PutMetricDataRequest(&cloudwatch.PutMetricDataInput{
			Namespace:  aws.String(namespace),
			MetricData: batch,
		})
		_, err := req.Send()
		return err
	}
	return nil
}
