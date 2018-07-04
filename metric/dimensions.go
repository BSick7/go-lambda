package metric

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
)

type Dimension struct {
	Name  string
	Value string
}

type Dimensions []Dimension

func (d Dimensions) ToAWS() []cloudwatch.Dimension {
	if len(d) <= 0 {
		return nil
	}
	dims := make([]cloudwatch.Dimension, 0)
	for _, dim := range d {
		dims = append(dims, cloudwatch.Dimension{
			Name:  aws.String(dim.Name),
			Value: aws.String(dim.Value),
		})
	}
	return dims
}
