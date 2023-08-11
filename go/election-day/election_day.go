package electionday

import "strconv"

// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {
	var newVote *int = &initialVotes
	return newVote
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
	voteCount := counter
	if voteCount != nil {
		return *voteCount
	} else {
		return 0
	}
}

// IncrementVoteCount increments the value in a vote counter.
func IncrementVoteCount(counter *int, increment int) {
	*counter += increment
}

// NewElectionResult creates a new election result.
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	result := &ElectionResult{
		Name:  candidateName,
		Votes: votes,
	}
	return result
}

// DisplayResult creates a message with the result to be displayed.
func DisplayResult(result *ElectionResult) string {
	resultDisplay := result.Name + " " + "(" + strconv.Itoa(result.Votes) + ")"
	return resultDisplay
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map.
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	results[candidate] -= 1
}
