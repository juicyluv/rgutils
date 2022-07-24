package convert

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func TimeToTimestamppb(v *time.Time) *timestamppb.Timestamp {
	if v == nil {
		return nil
	}

	return timestamppb.New(*v)
}

func TimestamppbToTime(v *timestamppb.Timestamp) *time.Time {
	if v == nil {
		return nil
	}

	t := v.AsTime()
	return &t
}
