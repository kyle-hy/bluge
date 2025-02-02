//  Copyright (c) 2020 Couchbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package searcher

import (
	"strings"

	"github.com/blugelabs/bluge/search"
)

func NewTermPrefixSearcher(indexReader search.Reader, prefix, field string,
	boost float64, scorer search.Scorer, compScorer search.CompositeScorer,
	options search.SearcherOptions) (search.Searcher, error) {
	// find the terms with this prefix
	byteBeg := []byte(prefix)
	byteEnd := incrementBytes(byteBeg)
	fieldDict, err := indexReader.DictionaryIterator(field, nil, byteBeg, byteEnd)
	if err != nil {
		return nil, err
	}
	defer func() {
		if cerr := fieldDict.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	var terms []string
	var freqs []int
	tfd, err := fieldDict.Next()
	for err == nil && tfd != nil {
		terms = append(terms, tfd.Term())
		freqs = append(freqs, strings.Count(field, tfd.Term())) // TODO 待确认正确性
		if tooManyClauses(len(terms)) {
			return nil, tooManyClausesErr(field, len(terms))
		}
		tfd, err = fieldDict.Next()
	}
	if err != nil {
		return nil, err
	}

	return NewMultiTermSearcher(indexReader, terms, freqs, field, boost, scorer, compScorer, options, true)
}
