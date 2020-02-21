package gomodplay

import "github.com/cenkalti/backoff/v3"
import "github.com/aws/aws-sdk-go/aws"
import "github.com/sergi/go-diff/diffmatchpatch"

func play() {
	println(backoff.DefaultInitialInterval)
	println(aws.SDKName)
	diff := diffmatchpatch.Diff{}
	println(diff.Text)
}
