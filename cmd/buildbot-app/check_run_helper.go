package main

import (
	"fmt"
	"time"
)

type CheckRunState string

const CheckRunStateQueued CheckRunState = "queued"
const CheckRunStateInProgress CheckRunState = "in_progress"
const CheckRunStateCompleted CheckRunState = "completed"

type CheckRunConclusion string

const CheckRunConclusionSuccess CheckRunConclusion = "success"
const CheckRunConclusionFailure CheckRunConclusion = "failure"
const CheckRunConclusionNeutral CheckRunConclusion = "neutral"
const CheckRunConclusionCancelled CheckRunConclusion = "cancelled"
const CheckRunConclusionSkipped CheckRunConclusion = "skipped"
const CheckRunConclusionTimedOut CheckRunConclusion = "timed_out"
const CheckRunConclusionActionRequired CheckRunConclusion = "action_required"

// See also https://docs.buildbot.net/latest/developer/results.html#build-result-codes
func CheckRunStateFromBuildbotResult(resultCode int) CheckRunConclusion {
	switch resultCode {
	case 0: // success
		return CheckRunConclusionSuccess
	case 1: // warning
		return CheckRunConclusionNeutral
	case 2: // failure
		return CheckRunConclusionFailure
	case 3: // skipped
		return CheckRunConclusionSkipped
	case 4: // exception
		return CheckRunConclusionFailure
	case 5: // retry
		return CheckRunConclusionFailure
	case 6: // cancelled
		return CheckRunConclusionCancelled
	}
	return CheckRunConclusionFailure
}

// GetTimePrefix returns a string like "[Mon, 02 Jan 2006 15:04:05 -0700]: "
// which can be used to format log messages from buildbot in the github check
// run's page. If "t" is nil, the current time is used.
func GetTimePrefix(t time.Time) string {
	return fmt.Sprintf("[%s]: ", t.Format(time.RFC1123Z))
}

func WrapMsgWithTimePrefix(message string, t time.Time) string {
	return fmt.Sprintf("%s%s", GetTimePrefix(t), message)
}
