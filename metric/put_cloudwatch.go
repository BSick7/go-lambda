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

func PutCloudwatch(namespace string, values []cloudwatch.MetricDatum) error {
	svc, err := services.NewCloudwatch()
	if err != nil {
		return fmt.Errorf("could not create cloudwatch service: %s", err)
	}

	flush := func(batch []cloudwatch.MetricDatum) error {
		if len(batch) <= 0 {
			return nil
		}
		req := svc.PutMetricDataRequest(&cloudwatch.PutMetricDataInput{
			Namespace:  aws.String(namespace),
			MetricData: batch,
		})
		_, err := req.Send()
		return err
	}

	batch := make([]cloudwatch.MetricDatum, 0)
	for _, item := range values {
		batch = append(batch, item)
		if len(batch) >= maxMetricsPerPut {
			if err := flush(batch); err != nil {
				return err
			}
			batch = make([]cloudwatch.MetricDatum, 0)
		}
	}
	return flush(batch)
}
