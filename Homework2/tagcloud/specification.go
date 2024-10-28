package tagcloud

import "sort"

// TagCloud aggregates statistics about used tags
type TagCloud struct {
	tags map[string]int
}

// TagStat represents statistics regarding single tag
type TagStat struct {
	Tag             string
	OccurrenceCount int
}

// New should create a valid TagCloud instance
func New() *TagCloud {
	return &TagCloud{
		tags: make(map[string]int),
	}
}

// AddTag should add a tag to the cloud if it wasn't present and increase tag occurrence count
// thread-safety is not needed
func (tc *TagCloud) AddTag(tag string) {
	tc.tags[tag]++
}

// TopN should return top N most frequent tags ordered in descending order by occurrence count
// if there are multiple tags with the same occurrence count then the order is defined by implementation
// if n is greater that TagCloud size then all elements should be returned
// thread-safety is not needed
// there are no restrictions on time complexity
func (tc *TagCloud) TopN(n int) []TagStat {

	// create a slice which holds tag statistics
	var stats []TagStat
	for tag, count := range tc.tags {
		stats = append(stats, TagStat{tag, count})
	}

	// sort the slice by occurrence count in descending order
	sort.Slice(stats, func(i, j int) bool {
		return stats[i].OccurrenceCount > stats[j].OccurrenceCount
	})

	// return the Top N element or if n is greater than TagCloud size then return all elements
	if n > len(stats) {
		return stats
	}

	return stats[:n]
}
