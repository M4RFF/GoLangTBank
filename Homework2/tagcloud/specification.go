package tagcloud

import (
	"slices"
	"strings"
)

// TagCloud aggregates statistics about used tags
type TagCloud struct {
	tags map[string]int // cloud of tags
}

// TagStat represents statistics regarding single tag
type TagStat struct {
	Tag             string
	OccurrenceCount int
}

// New should create a valid TagCloud instance
// Pointer returns
func New() *TagCloud {
	return &TagCloud{make(map[string]int)} // creat a new cloud of tags
}

// AddTag should add a tag to the cloud if it wasn't present and increase tag occurrence count
// thread-safety is not needed
func (tgs *TagCloud) AddTag(tag string) {
	tag = strings.ToLower(tag) // add tag to a cloud
	tgs.tags[tag]++            // increase tag occurrence count
}

// TopN should return top N most frequent tags ordered in descending order by occurrence count
// if there are multiple TAGS with the same occurrence COUNT then the order is defined by implementation
// if n is greater that TagCloud size then all elements should be returned
// thread-safety is not needed
// there are no restrictions on time complexity
func (tgs *TagCloud) TopN(n int) []TagStat {
	tagStats := make([]TagStat, 0, len(tgs.tags)) // create tagStats which collects tags

	for tag, count := range tgs.tags {
		tagStats = append(tagStats, TagStat{
			Tag:             tag,
			OccurrenceCount: count,
		}) // add each Tag and his Count
	}

	// sort by descending frequency of tag use
	slices.SortFunc(tagStats, func(a, b TagStat) int {
		if a.OccurrenceCount > b.OccurrenceCount {
			return -1
		}
		if a.OccurrenceCount < b.OccurrenceCount {
			return 1
		}
		return 0
	})

	if n > len(tagStats) { // if n greater than number of tags
		return tagStats // i return all elements from tagStats
	}

	return tagStats[:n] // return first n elements from tagStats
}
