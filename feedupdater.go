package main

import "net/http"

// Gets 10 feeds from DB ordered by last updated starting with nil entires
func (cfg *apiConfig) getFeedsToUpdate() ([]Feed, error) {
	var i []Feed
	return i, nil
}

// Updates feed data, uses mark as updated and get data from feed functions below
// then writes to the DB
func (cfg *apiConfig) updateFeed(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Takes in a feed and sets the last updated time to current time
func markFeedAsUpdated(feed Feed) (Feed, error) {
	return feed, nil
}

// Gets data from feed update fetch, needs to be updated to return data type and an error
func getDataFromFeedUpdate(feed Feed) error {
	return nil
}
