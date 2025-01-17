package util

import (
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"os"
)

func GetObjectsListIterator(c *cos.Client, prefix, marker string, include, exclude string) (objects []cos.Object, isTruncated bool, nextMarker string) {
	opt := &cos.BucketGetOptions{
		Prefix:       prefix,
		Delimiter:    "",
		EncodingType: "",
		Marker:       marker,
		MaxKeys:      0,
	}


	res, _, err := c.Bucket.Get(context.Background(), opt)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	objects = append(objects, res.Contents...)
	isTruncated = res.IsTruncated
	nextMarker = res.NextMarker

	if len(include) > 0 {
		objects = MatchCosPattern(objects, include, true)
	}
	if len(exclude) > 0 {
		objects = MatchCosPattern(objects, exclude, false)
	}

	return objects, isTruncated, nextMarker
}
